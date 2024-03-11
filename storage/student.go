package storage

import (
	"clone/lms_back/models"
	"clone/lms_back/pkg"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudent(db *sql.DB) studentRepo {
	return studentRepo{
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

func (c *studentRepo) Create(student models.Student) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO Student (
		"id",
		"full_name",
		"email",
		"age",
		"paid_sum",
		"status",
		"login",
		"password",
		"group_id")
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		student.Id, student.Full_Name,
		student.Email, student.Age,
		student.Paid_sum, student.Status,
		student.Login, student.Password)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *studentRepo) Update(student models.Student) (string, error) {

	query := ` UPDATE Student set
		"full_name"=$1
		"email"=$2
		"age"=$3
		"paid_sum"=$4
		"status"=$5
		"login"=$6
		"password"=$7
		"group_id"=$8
		"updated_at"=CURRENT_TIMESTAMP
		WHERE "id" = $9 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query,
		student.Full_Name,
		student.Email, student.Age,
		student.Paid_sum, student.Status,
		student.Login, student.Password,
		student.Id)

	if err != nil {
		return "", err
	}

	return student.Id, nil
}

func (c studentRepo) GetAllStudents(search string) (models.GetAllStudentsResponse, error) {
	var (
		resp   = models.GetAllStudentsResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and "full_name" ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count("id") OVER(),
				"id",
				"full_name",
				"email",
				"age",
				"paid_sum",
				"status",
				"login",
				"password",
				"group_id",
				"created_at"::date,
				"updated_at"
	  FROM Student WHERE "deleted_at" = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			student  = models.Student{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&student.Id,
			&student.Full_Name,
			&student.Email,
			&student.Age,
			&student.Paid_sum,
			&student.Login,
			&student.Password,
			&student.Group_id,
			&student.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		student.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Students = append(resp.Students, student)
	}
	return resp, nil
}

func (c *studentRepo) Delete(id string) error {

	query := ` UPDATE student set
			"deleted_at" = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE "id" = $1 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
