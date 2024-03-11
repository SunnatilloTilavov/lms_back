package storage

import (
	"clone/lms_back/models"
	"clone/lms_back/pkg"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type groupRepo struct {
	db *sql.DB
}

func NewGroup(db *sql.DB) groupRepo {
	return groupRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *groupRepo) Create(group models.Group) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO group (
		"id",
		"groupId",
		"brancheId",
		"teacher",
		"type")
		VALUES($1,$2,$3,$4,$5) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		group.GroupId,group.BrancheId,
	group.Teacher,group.Type)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *groupRepo) Update(group models.Group) (string, error) {

	query := ` UPDATE Group set
			"groupId"=$1,
			"brancheId"=$2,
			"name"=$3,
			"teacher"=$4,
			"type"=$5,
			"updated_at"=CURRENT_TIMESTAMP
		WHERE "id" = $6 
	`

	_, err := c.db.Exec(query,
		group.GroupId,group.BrancheId,
	group.Teacher,group.Type,
	group.Id)

	if err != nil {
		return "", err
	}

	return group.Id, nil
}

func (c groupRepo) GetAllGroups(search string) (models.GetAllGroupsResponse, error) {
	var (
		resp   = models.GetAllGroupsResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and "groupid" ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count("id") OVER(),
				"id",
				"groupId",
				"brancheId",
				"teacher",
				"type",
				"created_at"::date,
				"updated_at"
	  FROM Group WHERE "deleted_at" = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			group  = models.Group{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&group.Id,
			&group.GroupId,
			&group.BrancheId,
			&group.Teacher,
			&group.Type,
			&group.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		group.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Groups = append(resp.Groups, group)
	}
	return resp, nil
}

func (c *groupRepo) Delete(id string) error {

	query := ` UPDATE Goup set
			"deleted_at" = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE "id" = $1 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
