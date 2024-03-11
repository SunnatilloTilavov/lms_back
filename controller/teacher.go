package controller

import (
	"clone/lms_back/models"
	"encoding/json"
	"fmt"

	// "homework/2-oy/13-dars/rent_car/pkg/check"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Teacher(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateTeacher(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllTeachers(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateTeacher(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteTeacher(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	teacher := models.Teacher{}

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	// if err := check.ValidateCarYear(car.Year); err != nil {
	// 	fmt.Println("error while validating year: ", car.Year)
	// 	handleResponse(w, http.StatusBadRequest, err)
	// 	return
	// }

	id, err := c.Store.Teacher.Create(teacher)
	if err != nil {
		fmt.Println("error while creating Teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	teacher := models.Teacher{}

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	// if err := check.ValidateCarYear(car.Year); err != nil {
	// 	fmt.Println("error while validating year: ", car.Year)
	// 	handleResponse(w, http.StatusBadRequest, err)
	// 	return
	// }
	teacher.Id = r.URL.Query().Get("id")

	err := uuid.Validate(teacher.Id)
	if err != nil {
		fmt.Println("error while validating, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.Teacher.Update(teacher)
	if err != nil {
		fmt.Println("error while creating Teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	// if _, ok := values["search"]; ok {
	// 	search = values["search"][0]
	// }

	teachers, err := c.Store.Teacher.GetAllTeachers(search)
	if err != nil {
		fmt.Println("error while getting Branches, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, teachers)
}

func (c Controller) DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Teacher.Delete(id)
	if err != nil {
		fmt.Println("error while deleting Teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}
