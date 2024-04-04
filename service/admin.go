package service

import (
	"clone/lms_back/api/models"
	"clone/lms_back/storage"
	"context"
	"fmt"
)

type AdminService struct {
	storage storage.IStorage
}

func NewAdminService(storage storage.IStorage) AdminService {
	return AdminService{
		storage: storage,
	}
}
func (u AdminService) Create(ctx context.Context, Admin models.CreateAdmin) (models.CreateAdmin, error) {

	pKey, err := u.storage.Admin().Create(ctx, Admin)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Admin", err.Error())
		return models.CreateAdmin{}, err
	}

	return pKey, nil
}

func (u AdminService) Update(ctx context.Context, Admin models.UpdateAdmin) (models.UpdateAdmin, error) {

	pKey, err := u.storage.Admin().Update(ctx, Admin)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Admin", err.Error())
		return models.UpdateAdmin{}, err
	}

	return pKey, nil
}

func (u AdminService) Delete(ctx context.Context, id string) error {

	err := u.storage.Admin().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Admin", err.Error())
		return err
	}

	return nil
}

func (u AdminService) GetAllAdmins(ctx context.Context, Admin models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error) {

	pKey, err := u.storage.Admin().GetAll(ctx, Admin)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Admin", err.Error())
		return models.GetAllAdminsResponse{}, err
	}

	return pKey, nil
}

func (u AdminService) GetByIDAdmin(ctx context.Context, Id models.Admin) (models.Admin, error) {
	pKey, err := u.storage.Admin().GetByID(ctx, Id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID Admin", err.Error())
		return models.Admin{}, err
	}

	return pKey, nil
}
