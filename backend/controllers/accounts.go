package controllers

import (
	"api/teamdev/controllers/responses"
	auth "api/teamdev/controllers/token"
	"api/teamdev/errors"
	"api/teamdev/models"
	"api/teamdev/objects"
	"time"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type account struct {
	model *models.AccountM
}

func InitAccount(r *mux.Router, model *models.AccountM) {
	ctrl := &account{model}
	r.HandleFunc("/accounts/login", ctrl.LogIn).Methods("POST")
	r.HandleFunc("/accounts/logout", ctrl.LogOut).Methods("POST")
	r.HandleFunc("/accounts", ctrl.add).Methods("POST")
	r.HandleFunc("/accounts/{login}", ctrl.get).Methods("GET")
	r.HandleFunc("/accounts/{login}/role", ctrl.patch).Methods("PATCH")
}

// @Tags Accounts
// @Router /accounts/login [post]
// @Param account body objects.AccountDTO false "Authentication data"
// @Summary User's authorization
// @Produce json
// @Success 200 {object} objects.AccountDTO
// @Failure 400 Invalid value
// @Failure 401 Authentication failed
func (ctrl *account) LogIn(w http.ResponseWriter, r *http.Request) {
	accDTO := &objects.AccountDTO{}

	err := json.NewDecoder(r.Body).Decode(accDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	acc, err := ctrl.model.LogIn(accDTO.Login, accDTO.Password)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	auth.FillCookie(w, acc.Login)
}

// @Tags Accounts
// @Router /accounts/logout [post]
// @Summary Logging out
// @Produce json
// @Success 200 Successful opeartion
func (ctrl *account) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",

		HttpOnly: true,
	})

	responses.TextSuccess(w, "Logout was successful")
}

// @Tags Accounts
// @Router /accounts [post]
// @Param account body objects.AccountDTO true "Account data"
// @Summary Creates a new account
// @Produce json
// @Success 201 {object} objects.AccountDTO
// @Failure 400 Invalid value
func (ctrl *account) add(w http.ResponseWriter, r *http.Request) {
	accDTO := new(objects.AccountDTO)
	err := json.NewDecoder(r.Body).Decode(accDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	err = ctrl.model.Create(accDTO.ToModel())
	switch err {
	case nil:
		responses.TextSuccess(w, "Account creation was successful")
	case errors.AccountExists:
		responses.BadRequest(w, "Account associated with such login is already exists")
	case errors.DBAdditionError:
		responses.BadRequest(w, "Error in addition to DB")
	default:
		responses.BadRequest(w, "Error in addition the account")
	}
}

// @Tags Accounts
// @Router /accounts/{login} [get]
// @Summary Retrieves account
// @Param login path string true "Account login"
// @Produce json
// @Success 200 {object} objects.AccountDTO
func (ctrl *account) get(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	login := urlParams["login"]

	data, _ := ctrl.model.Find(login)
	responses.JsonSuccess(w, data.ToDTO())
}

// @Tags Accounts
// @Router /accounts/{login}/role [patch]
// @Param login path string true "Account login"
// @Param account body objects.AccountDTO true "Account data"
// @Summary Change user's role
// @Produce json
// @Success 200 Successful operation
// @Failure 400 Invalid value
// @Failure 403 Access denied
func (ctrl *account) patch(w http.ResponseWriter, r *http.Request) {
	accDTO := new(objects.AccountDTO)
	err := json.NewDecoder(r.Body).Decode(accDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	urlParams := mux.Vars(r)
	login := urlParams["login"]

	cur_login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	err = ctrl.model.UpdateRole(cur_login, login, accDTO.Role)
	switch err {
	case nil:
		responses.TextSuccess(w, "Account updation was successful")
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "user")
	case errors.AccessDenied:
		responses.AccessDenied(w)
	case errors.UnknownRole:
		responses.BadRequest(w, "Wrong role")
	default:
		responses.BadRequest(w, "Error in changing role")
	}
}
