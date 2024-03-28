package postgres

import (
	"database/sql"
	"fmt"
	"clone/lms_back/api/models"
	// "clone/lms_back/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudent(db *pgxpool.Pool) StudentRepo {
	return StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(student models.CreateStudent) (models.CreateStudent, error) {

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

	_, err := c.db.Exec(context.Background(),query,
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

func (c *StudentRepo) Update(student models.UpdateStudent) (models.UpdateStudent, error) {
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
	WHERE id = $9
	`
	_, err := c.db.Exec(context.Background(),query,
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
	if err != nil {
		return student, err
	}
	return student, nil
}

func (c *StudentRepo) GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
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

    rows, err := c.db.Query(context.Background(),`select count(id) over(),
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

func (c *StudentRepo) GetByID(id string) (models.Student, error) {
	var(
		student = models.Student{}
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
	 err := c.db.QueryRow(context.Background(),
	`select full_name, email,
	age, paid_sum, status, login,
	password, group_id, created_at,
	updated_at from student where id = $1`, id).Scan(
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
			Id: id,
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
	
func (c *StudentRepo) Delete(id string) error {
	query := `delete from student where id = $1`
	_, err := c.db.Exec(context.Background(),query, id)
	if err != nil {
		return err
	}
	return nil
}