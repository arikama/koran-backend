package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/requestresponse"
	"github.com/arikama/koran-backend/services"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func PostAuthGoogleController(
	googleAuthService services.GoogleAuthService,
	userManager managers.UserManager,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := requestresponse.PostAuthGoogleRequest{}

		err := json.NewDecoder(c.Request.Body).Decode(&request)
		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}

		if request.AuthCode == "" {
			utils.JsonError(c, http.StatusBadRequest, errors.New(`missing "auth_code" in body`))
			return
		}

		googleUser, err := googleAuthService.AuthUserCode(request.AuthCode)
		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}
		_, err = userManager.CreateUser(googleUser.Email, googleUser.Token)
		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}
		response := requestresponse.PostAuthGoogleResponse{
			Email:   googleUser.Email,
			Name:    googleUser.Name,
			Token:   googleUser.Token,
			Picture: googleUser.Picture,
		}
		utils.JsonData(c, http.StatusOK, response)
	}
}
