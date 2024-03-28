package storage

import "clone/lms_back/api/models"

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
	Create(models.CreateAdmin) (models.CreateAdmin, error)
	GetAll(request models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error)
	GetByID(id string) (models.Admin, error)
	Update(models.UpdateAdmin) (models.UpdateAdmin, error)
	Delete(string) error
}

type IBranchStorage interface {
	Create(models.CreateBranches) (models.CreateBranches, error)
	GetAll(request models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error)
	GetByID(id string) (models.Branches, error)
	Update(models.UpdateBranches) (models.UpdateBranches, error)
	Delete(string) error
}

type IGroupStorage interface {
	Create(models.CreateGroup) (models.CreateGroup, error)
	GetAll(request models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error)
	GetByID(id string) (models.Group, error)
	Update(models.UpdateGroup) (models.UpdateGroup, error)
	Delete(string) error
}

type IPaymentStorage interface {
	Create(models.CreatePayment) (models.CreatePayment, error)
	GetAll(request models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error)
	GetByID(id string) (models.Payment, error)
	Update(models.UpdatePayment) (models.UpdatePayment, error)
	Delete(string) error
}

type IStudentStorage interface {
	Create(models.CreateStudent) (models.CreateStudent, error)
	GetAll(request models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetByID(id string) (models.Student, error)
	Update(models.UpdateStudent) (models.UpdateStudent, error)
	Delete(string) error
}

type ITeacherStorage interface {
	Create(models.CreateTeacher) (models.CreateTeacher, error)
	GetAll(request models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetByID(id string) (models.Teacher, error)
	Update(models.UpdateTeacher) (models.UpdateTeacher, error)
	Delete(string) error
}