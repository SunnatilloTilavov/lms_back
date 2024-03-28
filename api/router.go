package api

import (
   "clone/lms_back/api/handler"
	"clone/lms_back/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)
	r := gin.Default()

	r.POST("/branches", h.CreateBranch)
	r.GET("/branches", h.GetAllBranches)
	r.PUT("/branches/:id", h.UpdateBranch)
	r.DELETE("/branches/:id", h.DeleteBranch)

	r.POST("/teacher", h.CreateTeacher)
	r.GET("/teacher", h.GetAllTeacher)
	r.PUT("/teacher/:id", h.UpdateTeacher)
	r.DELETE("/teacher/:id", h.DeleteTeacher)

	r.POST("/admin", h.CreateAdmin)
	r.GET("/admin", h.GetAllAdmins)
	r.PUT("/admin/:id", h.UpdateAdmin)
	r.DELETE("/admin/:id", h.DeleteAdmin)

	r.POST("/group", h.CreateGroup)
	r.GET("/group", h.GetAllGroups)
	r.PUT("/group/:id", h.UpdateGroup)
	r.DELETE("/group/:id", h.DeleteGroup)

	r.POST("/student", h.CreateStudent)
	r.GET("/student", h.GetAllStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)

	r.POST("/payment", h.CreatePayment)
	r.GET("/payment", h.GetAllPayment)
	r.PUT("/payment/:id", h.UpdatePayment)
	r.DELETE("/payment/:id", h.DeletePayment)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
