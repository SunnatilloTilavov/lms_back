package postgres

// import (
// 	"clone/lms_back/api/models"
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/go-faker/faker/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateTeacher(t *testing.T) {
// 	TeacherRepo := NewTeacher(db)

// 	reqTeacher := models.CreateTeacher{
// //
// 	}

// 	id, err := TeacherRepo.Create(context.Background(), reqTeacher)
// 	if assert.NoError(t, err) {
// 		models.Teacher{
// 			Id: id.Id,
// 		}
// 		createdTeacher, err := TeacherRepo.GetByID(context.Background(),models.Teacher{})
// 		if assert.NoError(t, err) {
// 			assert.Equal(t, reqTeacher.Full_name, createdTeacher.Full_name)
// 			assert.Equal(t, reqTeacher.Age, createdTeacher.Age)
// 		} else {
// 			return
// 		}
// 		fmt.Println("Created Teacher", createdTeacher)
// 	}
// }