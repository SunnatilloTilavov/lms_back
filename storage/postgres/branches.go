package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"clone/lms_back/api/models"
	// "clone/lms_back/pkg"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranch(db *pgxpool.Pool) branchRepo {
	return branchRepo{
		db: db,
	}
}

func (c *branchRepo) Create(branch models.CreateBranches) (models.CreateBranches, error) {

	id := uuid.New()
	query := `INSERT INTO branches (
		id,
		name,
		address,
		created_at)
		VALUES($1,$2,$3,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(context.Background(),query,
		id.String(),
		branch.Name,
		branch.Address)

	if err != nil {
		return branch, err
	}
	return branch, nil
}

func (c *branchRepo) Update(branch models.UpdateBranches) (models.UpdateBranches, error) {
	query := `update branches set 
	name=$1,
	address=$2,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $3 
	`
	_, err := c.db.Exec(context.Background(),query,
		branch.Name,
		branch.Address,
		branch.Id,
	)
	if err != nil {
		return branch, err
	}
	return branch, err
}

func (c *branchRepo) GetAll(req models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {
	var (
		resp   = models.GetAllBranchesResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(context.Background(),`select count(id) over(),
        id,
        name,
        address,
        created_at,
        updated_at,
        deleted_at FROM branches ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			id sql.NullString
			name sql.NullString
			address sql.NullString
			create_at sql.NullString
			updateAt sql.NullString
			deleted_at sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&create_at,
			&updateAt,
			&deleted_at); err != nil {
			return resp, err
		}

		resp.Branches = append(resp.Branches, models.Branches{
			Id: id.String,
			Name: name.String,
			Address: address.String,
			CreatedAt: create_at.String,
			UpdatedAt: updateAt.String,
			DeletedAt: deleted_at.String,
		})
	}
	return resp, nil
}

func (c *branchRepo) GetByID(id string) (models.Branches, error) {
	branch := models.Branches{}

	if err := c.db.QueryRow(context.Background(),`select id, name, address, created_at from branches where id = $1`, id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt); err != nil {
		return models.Branches{}, err
	}
	return branch, nil
}

func (c *branchRepo) Delete(id string) error {
	query := `delete from branches where id = $1`
	_, err := c.db.Exec(context.Background(),query, id)
	if err != nil {
		return err
	}
	return nil
}
