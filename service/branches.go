package service

import (
	"clone/lms_back/api/models"
	"clone/lms_back/storage"
	"context"
	"fmt"
)

type BranchesService struct {
	storage storage.IStorage
}

func NewBranchesService(storage storage.IStorage) BranchesService {
	return BranchesService{
		storage: storage,
	}
}
func (u BranchesService) Create(ctx context.Context, Branches models.CreateBranches) (models.CreateBranches, error) {

	pKey, err := u.storage.Branches().Create(ctx, Branches)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Branches", err.Error())
		return models.CreateBranches{}, err
	}

	return pKey, nil
}

func (u BranchesService) Update(ctx context.Context, Branches models.UpdateBranches) (models.UpdateBranches, error) {

	pKey, err := u.storage.Branches().Update(ctx, Branches)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Branches", err.Error())
		return models.UpdateBranches{}, err
	}

	return pKey, nil
}

func (u BranchesService) Delete(ctx context.Context, id string) error {

	err := u.storage.Branches().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Branches", err.Error())
		return err
	}

	return nil
}

func (u BranchesService) GetAllBranches(ctx context.Context, Branches models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {

	pKey, err := u.storage.Branches().GetAll(ctx, Branches)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Branches", err.Error())
		return models.GetAllBranchesResponse{}, err
	}

	return pKey, nil
}

func (u BranchesService) GetByIDBranches(ctx context.Context, Id models.Branches) (models.Branches, error) {
	pKey, err := u.storage.Branches().GetByID(ctx, Id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID Branches", err.Error())
		return models.Branches{}, err
	}

	return pKey, nil
}
