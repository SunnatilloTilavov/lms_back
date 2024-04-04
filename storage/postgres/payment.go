package postgres

import (
	"fmt"
	"clone/lms_back/api/models"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)

type paymentRepo struct {
	db *pgxpool.Pool
}

func NewPayment(db *pgxpool.Pool) paymentRepo {
	return paymentRepo{
		db: db,
	}
}

func (p *paymentRepo) Create(ctx context.Context,payment models.CreatePayment) (models.Payment, error) {
	id := uuid.New()
	query := `INSERT INTO payment(id, price,student_id,
		 branch_id,admin_id,created_at) 
	          VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP);		  
		`
	_, err := p.db.Exec(ctx,query,id.String(), payment.Price, 
	payment.Student_id, payment.Branch_id, payment.Admin_id)

	carpayment:=models.Payment{
		Id: id.String(),
		Price: payment.Price,
		Student_id: payment.Student_id,
		Admin_id: payment.Admin_id,
	}
	if err != nil {
		return carpayment, err
	}
	return carpayment, nil
}

func (p *paymentRepo) Update(ctx context.Context,payment models.UpdatePayment) (models.UpdatePayment, error) {
	query := `UPDATE payment SET price=$1, student_id=$2, branch_id=$3, 
	admin_id=$4, updated_at=CURRENT_TIMESTAMP WHERE id=$5`

	_, err := p.db.Exec(ctx,query, payment.Price, payment.Student_id, 
	payment.Branch_id, payment.Admin_id, payment.Id)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (p *paymentRepo) GetAll(ctx context.Context,req models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error) {
	var (
		resp   = models.GetAllPaymentsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := p.db.Query(ctx,`SELECT id, price, student_id, branch_id, admin_id, created_at, updated_at FROM payment`)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			id sql.NullString
			price sql.NullFloat64
			student_id sql.NullString
			branch_id sql.NullString
			admin_id sql.NullString
			created_at sql.NullString
			updated_at sql.NullString			
		)
		if err := rows.Scan(
			&id,
			&price,
			&student_id,
			&branch_id,
			&admin_id,
			&created_at,
			&updated_at,
		); err != nil {
			return resp, err
		}
		resp.Payments = append(resp.Payments, models.Payment{
			Id: id.String,
			Price: float64(price.Float64),
			Student_id: student_id.String,
			Branch_id: branch_id.String,
			Admin_id: admin_id.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}
	return resp, nil
}

func (p *paymentRepo) GetByID(ctx context.Context,payment models.Payment) (models.Payment, error) {
	var (
		price sql.NullFloat64
		student_id sql.NullString
		branch_id sql.NullString
		admin_id sql.NullString
		created_at sql.NullString
		updated_at sql.NullString			
	)

	row := p.db.QueryRow(ctx,
	`SELECT  price, student_id,
	 branch_id, admin_id, created_at,
	  updated_at FROM payment WHERE id = $1`, payment.Id)
 	err := row.Scan(
			&price,
			&student_id,
			&branch_id,
			&admin_id,
			&created_at,
			&updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Payment{}, fmt.Errorf("student with ID %s not found",payment.Id)
		}
		return payment, err
	}
	payment =models.Payment{
		Price: float64(price.Float64),
		Student_id: student_id.String,
		Branch_id: branch_id.String,
		Admin_id: admin_id.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}

	return payment, nil
}

func (p *paymentRepo) Delete(ctx context.Context,id string) error {
	_, err := p.db.Exec(context.Background(),`DELETE FROM payment WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
