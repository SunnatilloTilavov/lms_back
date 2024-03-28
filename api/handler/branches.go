package handler

import (
	"clone/lms_back/api/models"
	"fmt"
	_ "clone/lms_back/api/docs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

// CreateBranch godoc
// @Router 		   /branches [POST]
// @Summary 	   create a branch
// @Description    This api is creates a new branch and returns its id
// @Tags 		   branch
// @Accept		   json
// @Produce		   json
// @Param		   branch body    models.CreateBranches true "car"
// @Success		   200  {object}  models.CreateBranches
// @Failure		   400  {object}  models.Response
// @Failure		   404  {object}  models.Response
// @Failure		   500  {object}  models.Response
func (h Handler) CreateBranch(c *gin.Context) {
	branch := models.CreateBranches{}

	if err := c.ShouldBindJSON(&branch); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.Branches().Create(branch)
	if err != nil {
		handleResponse(c, "error while creating branch", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "created successfully", http.StatusOK, id)
}

// UpdateBranch godoc
// @Router                /branches/{id} [PUT]
// @Summary 			  update a branch
// @Description:          this api updates branch information
// @Tags 			      branch
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Branch ID"
// @Param       		  car body models.UpdateBranches true "branch"
// @Success 		      200 {object} models.UpdateBranches
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateBranch(c *gin.Context) {

	branch := models.UpdateBranches{}
	if err := c.ShouldBindJSON(&branch); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	branch.Id = c.Param("id")
	err := uuid.Validate(branch.Id)

	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.Branches().Update(branch)
	if err != nil {
		handleResponse(c, "error while updating branch", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "updated successfully", http.StatusOK, id)
}

// GetAllBranch godoc
// @Router 			/branches [GET]
// @Summary 		get all branch
// @Description 	This API returns branch list
// @Tags 			branch
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllBranchesResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllBranches(c *gin.Context) {

	request := models.GetAllBranchesRequest{}

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	branches, err := h.Store.Branches().GetAll(request)
	if err != nil {
		handleResponse(c, "error while getting branches", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, branches)
}

// GetByIDBranch godoc
// @Router       /branches/{id} [GET]
// @Summary      return a branch by ID
// @Description  Retrieves a branch by its ID
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id path string true "Branch ID"
// @Success      200 {object} models.Branches
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDBranch(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	Branch, err := h.Store.Branches().GetByID(id)
	if err != nil {
		handleResponse(c, "error while getting branch by id", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "", http.StatusOK, Branch)
}

// DeleteBranch godoc
// @Router          /branches/{id} [DELETE]
// @Summary         delete a branch by ID
// @Description     Deletes a branch by its ID
// @Tags            branch
// @Accept          json
// @Produce         json
// @Param           id path string true "Branch ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteBranch(c *gin.Context) {
	
	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Store.Branches().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting branch", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "", http.StatusOK, id)
}

