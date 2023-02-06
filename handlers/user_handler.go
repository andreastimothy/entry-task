package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var input *dtos.Login
	err := c.ShouldBindJSON(&input)

	inputLogin := dtos.Login{
		Email:    input.Email,
		Password: input.Password,
	}

	if err != nil {
		err := helpers.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		c.Error(err)
		return
	}

	login, err := h.userService.Login(&inputLogin)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(http.StatusOK, "SUCCESS", login))
}

func (h *Handler) LoginWithGoogleHandler(c *gin.Context) {
	var request dtos.LoginWithGoogleRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Error(err)
		return
	}

	client := &http.Client{}
	googleRequest, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/oauth2/v3/userinfo", nil)

	if err != nil {
		c.Error(err)
		return
	}

	googleRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %v", request.Token))

	googleResponse, err := client.Do(googleRequest)
	if err != nil {
		c.Error(err)
		return
	}

	var responseObject dtos.LoginWithGoogleResponse
	err = json.NewDecoder(googleResponse.Body).Decode(&responseObject)

	if err != nil {
		c.Error(err)
		return
	}

	if responseObject.Email == "" {
		err := helpers.FailedResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Error(err)
		return
	}

	user, err := h.userService.GetByEmail(responseObject.Email)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := helpers.GenerateJwtToken(&dtos.JwtToken{
		ID:       user.ID,
		Role:     user.Role,
		FullName: user.Name,
	})

	if err != nil {
		c.Error(err)
		return
	}

	response := &models.TokenResponse{
		IDToken: token,
		Name:    user.Name,
		Email:   user.Email,
		Photo:   user.Photo,
		Role:    user.Role,
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(http.StatusOK, "SUCCESS", response))
}
