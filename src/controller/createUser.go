package controller

import (
	"fmt"
	"net/http"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/controller/model/request"
	"github.com/GiovannaK/go-api/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("CreateUser function called", zap.String("journey", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while binding JSON info", err)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid JSON body: %s", err))

		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "123",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)

}
