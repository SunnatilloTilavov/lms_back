package storage

import (
	"clone/lms_back/api/models"
	"context"
)

type IStorage interface {
	CloseDB()
	Admin() IAdminStorage
	Branches() IBranchStorage
	Group() IGroupStorage
	Payment() IPaymentStorage
	Student() IStudentStorage
	Teacher() ITeacherStorage
}

type IAdminStorage interface {
	Create(context.Context,models.CreateAdmin) (models.CreateAdmin, error)
	GetAll(context.Context,models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error)
	GetByID(context.Context,models.Admin) (models.Admin, error)
	Update(context.Context,models.UpdateAdmin) (models.UpdateAdmin, error)
	Delete(context.Context,string) error
}

type IBranchStorage interface {
	Create(context.Context,models.CreateBranches) (models.CreateBranches, error)
	GetAll(context.Context,models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error)
	GetByID(context.Context,models.Branches) (models.Branches, error)
	Update(context.Context,models.UpdateBranches) (models.UpdateBranches, error)
	Delete(context.Context,string) error
}

type IGroupStorage interface {
	Create(context.Context,models.CreateGroup) (models.CreateGroup, error)
	GetAll(context.Context,models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error)
	GetByID(context.Context,models.Group) (models.Group, error)
	Update(context.Context,models.UpdateGroup) (models.UpdateGroup, error)
	Delete(context.Context,string) error
}

type IPaymentStorage interface {
	Create(context.Context,models.CreatePayment) (models.Payment, error)
	GetAll(context.Context, models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error)
	GetByID(context.Context,models.Payment) (models.Payment, error)
	Update(context.Context,models.UpdatePayment) (models.UpdatePayment, error)
	Delete(context.Context,string) error
}

type IStudentStorage interface {
	Create(context.Context,models.CreateStudent) (models.CreateStudent, error)
	GetAll(context.Context,models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetByID(context.Context,models.Student) (models.Student, error)
	Update(context.Context,models.UpdateStudent) (models.Student, error)
	Delete(context.Context,string) error
}

type ITeacherStorage interface {
	Create(context.Context,models.CreateTeacher) (models.CreateTeacher, error)
	GetAll( context.Context,models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetByID(context.Context,models.Teacher) (models.Teacher, error)
	Update(context.Context,models.UpdateTeacher) (models.UpdateTeacher, error)
	Delete(context.Context,string) error
}