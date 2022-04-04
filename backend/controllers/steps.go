package controllers

import (
	"api/teamdev/controllers/responses"
	auth "api/teamdev/controllers/token"
	"api/teamdev/errors"
	"api/teamdev/models"
	"api/teamdev/objects"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type stepCtrl struct {
	model *models.StepM
}

func InitSteps(r *mux.Router, model *models.StepM) {
	ctrl := &stepCtrl{model}
	r.HandleFunc("/recipes/{id}/steps", ctrl.getAll).Methods("GET")
	r.HandleFunc("/recipes/{id}/steps", ctrl.post).Methods("POST")

	r.HandleFunc("/recipes/{id}/steps/{step}", ctrl.get).Methods("GET")
	r.HandleFunc("/recipes/{id}/steps/{step}", ctrl.patch).Methods("PATCH")
	r.HandleFunc("/recipes/{id}/steps/{step}", ctrl.del).Methods("DELETE")
}

// @Tags Steps
// @Router /recipes/{id}/steps [get]
// @Summary Retrieves all recipe's steps
// @Param id path int true "Recipe's id"
// @Produce json
// @Success 200 {object} []objects.StepDTO
// @Failure 400 Invalid value
func (this *stepCtrl) getAll(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	data, err := this.model.GetSteps(id_rcp)

	switch err {
	case nil:
		responses.JsonSuccess(w, objects.Step{}.ArrToDTO(data))
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	default:
		responses.BadRequest(w, "Error in getting steps")
	}
}

// @Tags Steps
// @Router /recipes/{id}/steps [post]
// @Summary Posts step
// @Param id path int true "Recipe's id"
// @Param step body objects.StepDTO true "Step"
// @Produce json
// @Success 201 {object} objects.StepDTO
// @Failure 400 Invalid value
// @Failure 401 Authentication failed
// @Failure 403 Access denied
func (this *stepCtrl) post(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	stepDTO := new(objects.StepDTO)
	err = json.NewDecoder(r.Body).Decode(stepDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	err = this.model.AddStep(id_rcp, stepDTO.ToModel(), login)
	switch err {
	case nil:
		data, _ := this.model.GetStepLast(id_rcp)
		responses.JsonSuccess(w, data.ToDTO())
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "user")
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	case errors.StepExists:
		responses.BadRequest(w, "Step with this number is already exists")
	case errors.AccessDenied:
		responses.AccessDenied(w)
	default:
		responses.BadRequest(w, "Error in creation a new step")
	}
}

// @Tags Steps
// @Router /recipes/{id}/steps/{step} [get]
// @Summary Retrieves recipe's step by its number
// @Param id path int true "Recipe's id"
// @Param step path int true "Step's number"
// @Produce json
// @Success 200 {object} objects.StepDTO
// @Failure 400 Invalid value
func (this *stepCtrl) get(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]
	strStep := urlParams["step"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	step, err := strconv.Atoi(strStep)
	if err != nil {
		responses.BadRequest(w, "Wrong step's number")
		return
	}

	data, err := this.model.GetStepByNum(id_rcp, step)

	switch err {
	case nil:
		responses.JsonSuccess(w, data.ToDTO())
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	case errors.UnknownStep:
		responses.RecordNotFound(w, "step")
	default:
		responses.BadRequest(w, "Error in getting step")
	}
}

// @Tags Steps
// @Router /recipes/{id}/steps/{step} [patch]
// @Summary Updates step
// @Param id path int true "Recipe's id"
// @Param step path int true "Step's number"
// @Param value body objects.StepDTO true "Step data"
// @Produce json
// @Success 200 Successful operation
func (this *stepCtrl) patch(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]
	strStep := urlParams["step"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	step, err := strconv.Atoi(strStep)
	if err != nil {
		responses.BadRequest(w, "Wrong step's number")
		return
	}

	stepDTO := new(objects.StepDTO)
	err = json.NewDecoder(r.Body).Decode(stepDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	cur_login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	err = this.model.UpdateStep(cur_login, id_rcp, step, stepDTO.ToModel())
	switch err {
	case nil:
		responses.TextSuccess(w, "Account updation was successful")
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "user")
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	case errors.UnknownStep:
		responses.RecordNotFound(w, "step")
	case errors.AccessDenied:
		responses.AccessDenied(w)
	default:
		responses.BadRequest(w, "Error in changing role")
	}
}

// @Tags Steps
// @Router /recipes/{id}/steps/{step} [delete]
// @Summary Removes step
// @Param id path int true "Recipe's id"
// @Param step path int true "Step's number"
// @Produce json
// @Success 200 Successful operation
// @Failure 400 Invalid value
// @Failure 401 Authentication failed
func (this *stepCtrl) del(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]
	strStep := urlParams["step"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	step, err := strconv.Atoi(strStep)
	if err != nil {
		responses.BadRequest(w, "Wrong step's number")
		return
	}

	login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	err = this.model.DeleteStep(id_rcp, step, login)
	switch err {
	case nil:
		responses.TextSuccess(w, "Delete operation was successful")
	case errors.AccessDeleteDenied:
		responses.JwtAccessDenied(w)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, "user")
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "author")
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	case errors.UnknownStep:
		responses.RecordNotFound(w, "step")
	default:
		responses.BadRequest(w, "Delete operation was successful")
	}
}
