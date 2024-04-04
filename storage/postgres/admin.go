package postgres

import (
	"clone/lms_back/api/models"
	// "clone/lms_back/pkg"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type adminRepo struct {
	db *pgxpool.Pool
}

func NewAdmin(db *pgxpool.Pool) adminRepo {
	return adminRepo{
		db: db,
	}
}

func (c *adminRepo) Create(ctx context.Context,admin models.CreateAdmin) (models.CreateAdmin, error) {
	id := uuid.New()
	query := `INSERT INTO "admin" (
		id,
		full_name,
		email,
		age,
		status,
		login,
		password,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(ctx, query,
		id.String(),
		admin.Full_Name,
		admin.Email,
		admin.Age,
		admin.Status,
		admin.Login,
		admin.Password)

	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (c *adminRepo) Update(ctx context.Context,admin models.UpdateAdmin) (models.UpdateAdmin, error) {
	query := `update "admin" set 
	full_name=$1,
	email=$2,
	age=$3,
	status=$4,
	login=$5,
	password=$6,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $7
	`
	_, err := c.db.Exec(ctx, query,
		admin.Full_Name,
		admin.Email,
		admin.Age,
		admin.Status,
		admin.Login,
		admin.Password,
		admin.Id,
	)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (c *adminRepo) GetAll(ctx context.Context,req models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error) {
	var (
		resp   = models.GetAllAdminsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(ctx, `select count(id) over(),
        id,
        full_name,
        email,
		age,
		status,
		login,
		password,
        created_at,
        updated_at
        FROM "admin"`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			id        sql.NullString
			full_name sql.NullString
			email     sql.NullString
			age       sql.NullInt64
			login     sql.NullString
			status    sql.NullString
			pasword   sql.NullString
			create_at sql.NullString
			updatedAt sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&email,
			&age,
			&status,
			&login,
			&pasword,
			&create_at,
			&updatedAt); err != nil {
			return resp, err
		}
		resp.Admins = append(resp.Admins, models.Admin{

			Id:         id.String,
			Full_Name:  full_name.String,
			Email:      email.String,
			Age:        int(age.Int64),
			Status:     status.String,
			Login:      login.String,
			Password:   pasword.String,
			Created_at: create_at.String,
			Updated_at: updatedAt.String,
		})
	}
	return resp, nil
}

func (c *adminRepo) GetByID(ctx context.Context,admin  models.Admin) (models.Admin, error) {
	var (
		full_name sql.NullString
		email     sql.NullString
		age       sql.NullInt64
		login     sql.NullString
		status    sql.NullString
		pasword   sql.NullString
		create_at sql.NullString
	)
	err := c.db.QueryRow(ctx,
		`select full_name, email,
	age, status, login, 
	password, created_at
	from "admin" where id = $1`, admin.Id).Scan(
		&full_name,
		&email,
		&age,
		&status,
		&login,
		&pasword,
		&create_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Admin{}, fmt.Errorf("student with ID %s not found", admin.Id)
		}
		return models.Admin{}, err
	}
	admin = models.Admin{
		Full_Name:  full_name.String,
		Email:      email.String,
		Age:        int(age.Int64),
		Status:     status.String,
		Login:      login.String,
		Password:   pasword.String,
		Created_at: create_at.String,
	}
	return admin, nil
}

func (c *adminRepo) Delete(ctx context.Context,id string) error {
	query := `delete from "admin" where id = $1`
	_, err := c.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
