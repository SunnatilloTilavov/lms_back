package controller

import (
	"clone/lms_back/models"
	"encoding/json"
	"fmt"

	// "homework/2-oy/13-dars/rent_car/pkg/check"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Student(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateStudent(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllStudents(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateStudent(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteStudent(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateStudent(w http.ResponseWriter, r *http.Request) {
	Student := models.Student{}

	if err := json.NewDecoder(r.Body).Decode(&Student); err != nil {
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

	id, err := c.Store.Student.Create(Student)
	if err != nil {
		fmt.Println("error while creating Student, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	Student := models.Student{}

	if err := json.NewDecoder(r.Body).Decode(&Student); err != nil {
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
	Student.Id = r.URL.Query().Get("id")

	err := uuid.Validate(Student.Id)
	if err != nil {
		fmt.Println("error while validating, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.Student.Update(Student)
	if err != nil {
		fmt.Println("error while creating Branche, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllStudents(w http.ResponseWriter, r *http.Request) {
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

	branches, err := c.Store.Student.GetAllStudents(search)
	if err != nil {
		fmt.Println("error while getting Students, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, branches)
}

func (c Controller) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Student.Delete(id)
	if err != nil {
		fmt.Println("error while deleting Student, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}
