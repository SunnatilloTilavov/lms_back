package postgres

import (
	"clone/lms_back/api/models"
	"database/sql"
	"fmt"
	"time"

	// "clone/lms_back/pkg"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudent(db *pgxpool.Pool) StudentRepo {
	return StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(ctx context.Context,student models.CreateStudent) (models.CreateStudent, error) {
	id := uuid.New()
	query := `INSERT INTO student (
		id,
		full_name,
		email,
		age,
		paid_sum,
		status,
		login,
		password,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(ctx,query,
		id.String(),
		student.Full_Name,
	    student.Email,
		student.Age,
		student.PaidSum,
		student.Status,
		student.Login,
		student.Password,
	)

	if err != nil {
		return student, err
	}
	return student, nil
}

func (c *StudentRepo) Update(ctx context.Context,student models.UpdateStudent) (models.Student, error) {
	fmt.Println("id==",student.Id)
	// id:=student.Id[1:]
	query := `update student set 
	full_name=$1,
	email=$2,
	age=$3,
	paid_sum=$4,
	login=$5,
	password=$6,
	group_id=$7,
	status=$8,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $9`

	_, err := c.db.Exec(ctx,query,
		student.Full_Name,
		student.Email,
		student.Age,
		student.PaidSum,
		student.Login,
		student.Password,
		student.GroupID,
		student.Status,
		student.Id,
    )
	student1:=models.Student{
		Id: student.Id,
		Full_Name: student.Full_Name,
		Email: student.Email,
		Age: student.Age,
		PaidSum: student.PaidSum,
		Login: student.Login,
		Password: student.Password,
		Status: student.Status,
		Updated_At: time.Now().String(),

	}
	if err != nil {
		return student1, err
	}
	return student1, nil
}

func (c *StudentRepo) GetAll(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
    var (
		resp   = models.GetAllStudentsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

    rows, err := c.db.Query(ctx,`select count(id) over(),
        id,
        full_name,
		email,
		age,
		paid_sum,
		status,
        login,
		password,
		group_id,
        created_at,
        updated_at FROM student` + filter + ``)
    if err != nil {
        return resp, err
    }

    for rows.Next() {
        var (
            updateAt  sql.NullString
			group_id  sql.NullString
			id sql.NullString
			full_name sql.NullString
			email sql.NullString
			age sql.NullInt64
			paid_sum sql.NullFloat64
			status sql.NullString
			login sql.NullString
			password sql.NullString
			created_at sql.NullString
        )
        if err := rows.Scan(
            &resp.Count,
            &id,
            &full_name,
			&email,
			&age,
			&paid_sum,
			&status,
            &login,
			&password,
			&group_id,
            &created_at,
            &updateAt);
			 err != nil {
            return resp, err
        }
        resp.Students = append(resp.Students, models.Student{
			Id:id.String,
			Full_Name: full_name.String,
			Email: email.String,
			Age: int(age.Int64),
			PaidSum: float64(paid_sum.Float64),
			Status: status.String,
			Login: login.String,
			Password: password.String,
			GroupID: group_id.String,
			Created_At: created_at.String,
			Updated_At: updateAt.String, 
		})
    }
    return resp, nil
}

func (c *StudentRepo) GetByID(ctx context.Context,student models.Student) (models.Student, error) {
	var(
		updateAt  sql.NullString
		group_id  sql.NullString
		full_name sql.NullString
		email sql.NullString
		age sql.NullInt64
		paid_sum sql.NullFloat64
		status sql.NullString
		login sql.NullString
		password sql.NullString
		created_at sql.NullString
		
	)
	 err := c.db.QueryRow(ctx,
	`select full_name, email,
	age, paid_sum, status, login,
	password, group_id, created_at,
	updated_at from student where id = $1`, student.Id).Scan(
		&full_name,
		&email,
		&age,
		&paid_sum,
		&status,
		&login,
		&password,
		&group_id,
		&created_at,
		&updateAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Student{}, fmt.Errorf("student with ID %s not found", student.Id)
			}
			return models.Student{}, err
		}
	
		student = models.Student{
			Full_Name:   full_name.String,
			Email:      email.String,
			Age:        int(age.Int64),
			PaidSum:    float64(paid_sum.Float64),
			Status:     status.String,
			Login:      login.String,
			Password:   password.String,
			GroupID:    group_id.String,
			Created_At:  created_at.String,
			Updated_At:  updateAt.String,
		}
		return student, nil
	}
	
func (c *StudentRepo) Delete(ctx context.Context,id string) error {
	query := `delete from student where id = $1`
	_, err := c.db.Exec(ctx,query, id)
	if err != nil {
		return err
	}
	return nil
}