package views

import (
	"encoding/json"
	"net/http"

	"github.com/knid/timezen/internal/models"
	"github.com/knid/timezen/internal/requests"
	"github.com/knid/timezen/internal/utils"
)

func (vw *Views) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	user, err := vw.Controller.GetUserFromRequest(r)
	if err != nil {
		utils.JSONResponse(w, 401, err)
		return 
	}

	json.NewEncoder(w).Encode(user)
}

func (vw *Views) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCreateRequest requests.UserCreateRequest
	json.NewDecoder(r.Body).Decode(&userCreateRequest)

	user := models.User{
		Name: userCreateRequest.Name,
		Surname: userCreateRequest.Surname,
		Email: userCreateRequest.Email,
		Password: userCreateRequest.Password,
	}

	result := vw.Controller.DB.Create(&user)
	if result.Error != nil {
		utils.JSONResponse(w, 401, result.Error)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (vw *Views) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest requests.UserLoginRequests
	json.NewDecoder(r.Body).Decode(&userLoginRequest)

	user, err := vw.Controller.GetUserByEmailPassword(userLoginRequest.Email, userLoginRequest.Password)
	if err != nil {
		utils.JSONResponse(w, 401, err)
		return
	}

	token, err := vw.Controller.CreateToken(user)
	if err != nil {
		utils.JSONResponse(w, 401, err)
		return
	}

	json.NewEncoder(w).Encode(token)
}


