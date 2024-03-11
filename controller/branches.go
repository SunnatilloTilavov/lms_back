package controller

import (
	"clone/lms_back/models"
	"encoding/json"
	"fmt"

	// "homework/2-oy/13-dars/rent_car/pkg/check"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Branche(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateBranche(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllBranches(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateBranche(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteBranche(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateBranche(w http.ResponseWriter, r *http.Request) {
	branche := models.Branche{}

	if err := json.NewDecoder(r.Body).Decode(&branche); err != nil {
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

	id, err := c.Store.Branche.Create(branche)
	if err != nil {
		fmt.Println("error while creating Branche, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateBranche(w http.ResponseWriter, r *http.Request) {
	branche := models.Branche{}

	if err := json.NewDecoder(r.Body).Decode(&branche); err != nil {
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
	branche.Id = r.URL.Query().Get("id")

	err := uuid.Validate(branche.Id)
	if err != nil {
		fmt.Println("error while validating, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.Branche.Update(branche)
	if err != nil {
		fmt.Println("error while creating Branche, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllBranches(w http.ResponseWriter, r *http.Request) {
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

	branches, err := c.Store.Branche.GetAllBranches(search)
	if err != nil {
		fmt.Println("error while getting Branches, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, branches)
}

func (c Controller) DeleteBranche(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Branche.Delete(id)
	if err != nil {
		fmt.Println("error while deleting Branche, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}
