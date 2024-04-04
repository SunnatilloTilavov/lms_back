package postgres

import (
	"database/sql"
	"fmt"
	"clone/lms_back/api/models"
	"clone/lms_back/pkg"
	"strconv"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)

type GroupRepo struct {
	db *pgxpool.Pool
}

func NewGroup(db *pgxpool.Pool) GroupRepo {
	return GroupRepo{
		db: db,
	}
}

func (g *GroupRepo) Create(ctx context.Context,group models.CreateGroup) (models.CreateGroup, error) {
	id := uuid.New()
	var group_unique_id string

	maxQuery := `
	SELECT MAX(group_id) 
	FROM "group"
	`
	err := g.db.QueryRow(ctx,maxQuery).Scan(&group_unique_id)
	if err != nil {
		if err.Error() != "can't scan into dest[0]: cannot scan null into *string" && err.Error() != "no rows in result set" {
			return group, err
		} else {
			group_unique_id = "GR-0000000"
		}
	}

	digit := 0
	if len(group_unique_id) > 2 {
		digit, err = strconv.Atoi(group_unique_id[3:])
		if err != nil {
			return group, err
		}
	}

	query := `INSERT INTO "group" (
		id,
		group_id,
		branch_id,
		teacher_id,
		type,
		created_at) 
		VALUES($1,$2,$3,$4,$5,CURRENT_TIMESTAMP)
		`

	_, err = g.db.Exec(ctx,query,
		id.String(),
		"Gr-"+pkg.GetSerialId(digit),
		group.Branch_id,
		group.Teacher_id,
		group.Type)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (g *GroupRepo) Update(ctx context.Context,group models.UpdateGroup) (models.UpdateGroup, error) {
	query := `UPDATE "group" SET
		type=$1,
		updated_at=CURRENT_TIMESTAMP
		WHERE id=$2`

	_, err := g.db.Exec(ctx,query, group.Type, group.Id)
	if err != nil {
		return group, nil
	}
	return group, nil
}

func (g *GroupRepo) GetAll(ctx context.Context,req models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error) {
	var (
		resp   = models.GetAllGroupsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := g.db.Query(ctx,`SELECT count (id) OVER(),
        id,
        group_id,
        branch_id,
        teacher,
        type,
        created_at,
        updated_at FROM "group"` + filter + ``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (

			id sql.NullString
			group_id sql.NullString
			branch_id sql.NullString
			teacher_id sql.NullString
			type1 sql.NullString
			create_at sql.NullString
			updateAt   sql.NullString
		
			
		)
		if err := rows.Scan(
			&resp.Count,
			&id,
			&group_id,
			&branch_id,
			&teacher_id,
			&type1,
			&create_at,
			&updateAt); err != nil {
			return resp, err
		}
		resp.Groups = append(resp.Groups, models.Group{
			Id:id.String,
			Group_id: group_id.String,
			Branch_id: branch_id.String,
			Teacher_id: teacher_id.String,
			Type: type1.String,
			Created_at: create_at.String,
			Updated_at: updateAt.String,

		})
	}
	return resp, nil
}

func (g *GroupRepo) GetByID(ctx context.Context,group models.Group) (models.Group, error) {
	var (
		group_id sql.NullString
		branch_id sql.NullString
		teacher_id sql.NullString
		type1 sql.NullString
		create_at sql.NullString
		updateAt   sql.NullString	
	)

	err := g.db.QueryRow(ctx,
	`SELECT  group_id,
	 branch_id, teacher_id, 
	 type, created_at, 
	 updated_at FROM "group" WHERE id = $1`, group.Id).Scan(
		&group_id,
		&branch_id,
		&teacher_id,
		&type1,
		&create_at,
		&updateAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Group{}, fmt.Errorf("student with ID %s not found", group.Id)
			}
		return models.Group{}, err
	}
	group = models.Group{
		Group_id: group_id.String,
		Branch_id: branch_id.String,
		Teacher_id: teacher_id.String,
		Type: type1.String,
		Created_at: create_at.String,
		Updated_at: updateAt.String,

	}
	return group, nil
}

func (g *GroupRepo) Delete(ctx context.Context,id string) error {
	query := `delete from "group" where id = $1`
	_, err := g.db.Exec(context.Background(),query, id)
	if err != nil {
		return err
	}
	return nil
}
