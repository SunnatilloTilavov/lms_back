package service

import (
	"clone/lms_back/api/models"
	"clone/lms_back/storage"
	"context"
	"fmt"
)

type StudentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) StudentService {
	return StudentService{
		storage: storage,
	}
}
func (u StudentService) Create(ctx context.Context, Student models.CreateStudent) (models.CreateStudent, error) {

	pKey, err := u.storage.Student().Create(ctx, Student)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Student", err.Error())
		return models.CreateStudent{}, err
	}

	return pKey, nil
}

func (u StudentService) Update(ctx context.Context, Student models.UpdateStudent) (models.Student, error) {

	pKey, err := u.storage.Student().Update(ctx, Student)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Student", err.Error())
		return models.Student{}, err
	}

	return pKey, nil
}

func (u StudentService) Delete(ctx context.Context, id string) error {

	err := u.storage.Student().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Student", err.Error())
		return err
	}

	return nil
}

func (u StudentService) GetAllStudents(ctx context.Context, Student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {

	pKey, err := u.storage.Student().GetAll(ctx, Student)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Student", err.Error())
		return models.GetAllStudentsResponse{}, err
	}

	return pKey, nil
}

func (u StudentService) GetByIDStudent(ctx context.Context, Id models.Student) (models.Student, error) {
	pKey, err := u.storage.Student().GetByID(ctx, Id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID Student", err.Error())
		return models.Student{}, err
	}

	return pKey, nil
}
