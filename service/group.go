package service

import (
	"clone/lms_back/api/models"
	"clone/lms_back/storage"
	"context"
	"fmt"
)

type GroupService struct {
	storage storage.IStorage
}

func NewGroupService(storage storage.IStorage) GroupService {
	return GroupService{
		storage: storage,
	}
}
func (u GroupService) Create(ctx context.Context, Group models.CreateGroup) (models.CreateGroup, error) {

	pKey, err := u.storage.Group().Create(ctx, Group)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Group", err.Error())
		return models.CreateGroup{}, err
	}

	return pKey, nil
}

func (u GroupService) Update(ctx context.Context, Group models.UpdateGroup) (models.UpdateGroup, error) {

	pKey, err := u.storage.Group().Update(ctx, Group)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Group", err.Error())
		return models.UpdateGroup{}, err
	}

	return pKey, nil
}

func (u GroupService) Delete(ctx context.Context, id string) error {

	err := u.storage.Group().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Group", err.Error())
		return err
	}

	return nil
}

func (u GroupService) GetAllGroups(ctx context.Context, Group models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error) {

	pKey, err := u.storage.Group().GetAll(ctx, Group)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Group", err.Error())
		return models.GetAllGroupsResponse{}, err
	}

	return pKey, nil
}

func (u GroupService) GetByIDGroup(ctx context.Context, Id models.Group) (models.Group, error) {
	pKey, err := u.storage.Group().GetByID(ctx, Id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID Group", err.Error())
		return models.Group{}, err
	}

	return pKey, nil
}
