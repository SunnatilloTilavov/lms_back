package storage

import (
	"clone/lms_back/models"
	"clone/lms_back/pkg"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type brancheRepo struct {
	db *sql.DB
}

func NewBranche(db *sql.DB) brancheRepo {
	return brancheRepo{
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

func (c *brancheRepo) Create(branche models.Branche) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO Branches (
		"id",
		"name",
		"address")
		VALUES($1,$2,$3) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		branche.Name, branche.Address)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *brancheRepo) Update(branche models.Branche) (string, error) {

	query := ` UPDATE Branches set
			"name"=$1,
			"addess"=$2,
			"updated_at"=CURRENT_TIMESTAMP
		WHERE "id" = $3 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query,
		branche.Name, branche.Address, branche.Id)

	if err != nil {
		return "", err
	}

	return branche.Id, nil
}

func (c brancheRepo) GetAllBranches(search string) (models.GetAllBranchesResponse, error) {
	var (
		resp   = models.GetAllBranchesResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and "name" ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count("id") OVER(),
				"id", 
				"name",
				"address",
				"created_at"::date,
				"updated_at"
	  FROM Branches WHERE "deleted_at" = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			branche  = models.Branche{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&branche.Id,
			&branche.Name,
			&branche.Address,
			&branche.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		branche.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Branches = append(resp.Branches, branche)
	}
	return resp, nil
}

func (c *brancheRepo) Delete(id string) error {

	query := ` UPDATE Branches set
			"deleted_at" = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE "id" = $1 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
