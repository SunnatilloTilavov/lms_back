package service

import (
	"context"
	"fmt"
	"clone/lms_back/api/models"
	"clone/lms_back/storage"
)

type TeacherService struct {
	storage storage.IStorage
}

func NewTeacherService(storage storage.IStorage) TeacherService {
	return TeacherService{
		storage: storage,
	}
}
func (u TeacherService) Create(ctx context.Context, Teacher models.CreateTeacher) (models.CreateTeacher, error) {

	pKey, err := u.storage.Teacher().Create(ctx, Teacher)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Teacher", err.Error())
		return models.CreateTeacher{}, err
	}

	return pKey, nil
}

func (u TeacherService) Update(ctx context.Context, Teacher models.UpdateTeacher) (models.UpdateTeacher, error) {

	pKey, err := u.storage.Teacher().Update(ctx, Teacher)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Teacher", err.Error())
		return models.UpdateTeacher{}, err
	}

	return pKey, nil
}

func (u TeacherService) Delete(ctx context.Context,id string) error {

	err := u.storage.Teacher().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Teacher", err.Error())
		return err
	}

	return nil
}

func (u TeacherService) GetAllTeachers(ctx context.Context, Teacher models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {

	pKey, err := u.storage.Teacher().GetAll(ctx, Teacher)
	if err != nil {
	 fmt.Println("ERROR in service layer while getalling Teacher", err.Error())
	 return models.GetAllTeachersResponse{}, err
	}
   
	return pKey, nil
   }

   func (u TeacherService) GetByIDTeacher(ctx context.Context,Id models.Teacher) (models.Teacher, error) {
	pKey, err := u.storage.Teacher().GetByID(ctx,Id)
	if err != nil {
	 fmt.Println("ERROR in service layer while getbyID Teacher", err.Error())
	 return models.Teacher{}, err
	}
   
	return pKey, nil
   }