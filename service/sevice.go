package service

import (
	"clone/lms_back/storage"
)

type IServiceManager interface {
	Student() StudentService
	Teacher() TeacherService 
	Payment() PaymentService 
	Group() GroupService 
	Branches() BranchesService 
	Admin() AdminService 
}

type Service struct {
	StudentService StudentService
	TeacherService TeacherService
	PaymentService PaymentService
	GroupService GroupService
	BranchesService BranchesService
	AdminService AdminService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.StudentService = NewStudentService(storage)
	services.TeacherService = NewTeacherService(storage)
	services.PaymentService = NewPaymentService(storage)
	services.GroupService = NewGroupService(storage)
	services.BranchesService = NewBranchesService(storage)
	services.AdminService = NewAdminService(storage)
	return services
}

func (s Service) Student() StudentService {
	return s.StudentService
}

func(s Service) Teacher()TeacherService{
	return s.TeacherService
}
func(s Service) Payment()PaymentService{
	return s.PaymentService
}

func(s Service) Group()GroupService{
	return s.GroupService
}

func(s Service) Branches()BranchesService{
	return s.BranchesService
}

func(s Service) Admin()AdminService{
	return s.AdminService
}