package storage

import (
	"clone/lms_back/models"
	"clone/lms_back/pkg"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type teacherRepo struct {
	db *sql.DB
}

func NewTeacher(db *sql.DB) teacherRepo {
	return teacherRepo{
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

func (c *teacherRepo) Create(teacher models.Teacher) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO Teacher (
		"id",
		"full_name",
		"email",
		"age",
		"status",
		"login",
		"password",)
		VALUES($1,$2,$3,$4,$5,$6,$7) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		teacher.Full_Name, 
		teacher.Email	,
		teacher.Age	,
		teacher.Status    ,
		teacher.Login	,
		teacher.Password)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *teacherRepo) Update(teacher models.Teacher) (string, error) {

	query := ` UPDATE Teacher set
	"full_name"=$1,
	"email"=$2    ,
	"age"=$3      ,
	"status"=$4,
	"login"=$5,
	"password"=$6,
		"updated_at"=CURRENT_TIMESTAMP
		WHERE "id" = $7 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query,
		teacher.Full_Name, 
		teacher.Email	,
		teacher.Age	,
		teacher.Status    ,
		teacher.Login	,
		teacher.Password)

	if err != nil {
		return "", err
	}

	return teacher.Id, nil
}

func (c teacherRepo) GetAllTeachers(search string) (models.GetAllTeachersResponse, error) {
	var (
		resp   = models.GetAllTeachersResponse{}
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
	 			"status",
	 			"login",
	 			"password",
	 			"created_at"::date,
				"updated_at"
	FROM Teacher WHERE "deleted_at" = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			teacher  = models.Teacher{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&teacher.Id,
			&teacher.Full_Name, 
			&teacher.Email	,
			&teacher.Age	,
			&teacher.Status    ,
			&teacher.Login	,
			&teacher.Password,
			&teacher.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		teacher.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Teachers = append(resp.Teachers, teacher)
	}
	return resp, nil
}

func (c *teacherRepo) Delete(id string) error {

	query := ` UPDATE Teacher set
			"deleted_at" = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE "id" = $1 AND "deleted_at"=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
