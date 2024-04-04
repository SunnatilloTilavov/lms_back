package service

import (
	"clone/lms_back/api/models"
	"clone/lms_back/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
)

type ServiceRepo struct {
	db *pgxpool.Pool
}

func NewPayment(db *pgxpool.Pool) ServiceRepo {
	return ServiceRepo{
		db: db,
	}
}

type PaymentService struct {
	storage storage.IStorage
}

func NewPaymentService(storage storage.IStorage) PaymentService {
	return PaymentService{
		storage: storage,
	}
}

func (u PaymentService) Create(ctx context.Context, req models.CreatePayment) (resp models.Payment, err error) {
	pKey, err := u.storage.Payment().Create(ctx, req)
	if err != nil {
		return models.Payment{}, fmt.Errorf("error in service layer while creating Payment: %w", err)
	}
	student:=models.Student{
		Id: pKey.Student_id,
	}
	payment:=models.Payment{
		Id: pKey.Id,
	}

	if req.Student_id != "" && req.Price > 0 {
		student, err := u.storage.Student().GetByID(ctx, student)
		if err != nil {
			return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
		}
		
		student.PaidSum += req.Price

		_, err = u.storage.Student().Update(ctx, models.UpdateStudent{
			Id:      pKey.Student_id,
			PaidSum: student.PaidSum,
			GroupID: student.GroupID,
			Email: student.Email,
			Age: student.Age,
			Login: student.Login,
			Password: student.Password,
			Status: student.Status,
		})

		if err != nil {
			return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	resp, err = u.storage.Payment().GetByID(ctx,payment)
	if err != nil {
		return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}



func (u PaymentService) Update(ctx context.Context, req models.UpdatePayment) (resp models.Payment,err error) {

	pKey, err := u.storage.Payment().Update(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Payment", err.Error())
		return models.Payment{}, err
	}
	student:=models.Student{
		Id: pKey.Student_id,
	}
	payment:=models.Payment{
		Id: pKey.Id,
	}

	if req.Student_id != "" && req.Price > 0 {
		student, err := u.storage.Student().GetByID(ctx, student)
		if err != nil {
			return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
		}
		
		student.PaidSum += req.Price

		_, err = u.storage.Student().Update(ctx, models.UpdateStudent{
			Id:      pKey.Student_id,
			PaidSum: student.PaidSum,
			GroupID: student.GroupID,
			Email: student.Email,
			Age: student.Age,
			Login: student.Login,
			Password: student.Password,
			Status: student.Status,
		})

		if err != nil {
			return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	resp, err = u.storage.Payment().GetByID(ctx,payment)
	if err != nil {
		return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}



func (u PaymentService) Delete(ctx context.Context, id string) error {

	err := u.storage.Payment().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Payment", err.Error())
		return err
	}

	return nil
}

func (u PaymentService) GetAllPayments(ctx context.Context, Payment models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error) {

	pKey, err := u.storage.Payment().GetAll(ctx, Payment)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Payment", err.Error())
		return models.GetAllPaymentsResponse{}, err
	}

	return pKey, nil
}

func (u PaymentService) GetByIDPayment(ctx context.Context, Id models.Payment) (models.Payment, error) {
	pKey, err := u.storage.Payment().GetByID(ctx, Id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID Payment", err.Error())
		return models.Payment{}, err
	}

	return pKey, nil
}
