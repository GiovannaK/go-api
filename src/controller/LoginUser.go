package controller

import (
	"fmt"
	"net/http"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/controller/model/request"
	"github.com/GiovannaK/go-api/src/model"
	"github.com/GiovannaK/go-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Login function called", zap.String("journey", "Login"))

	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error while binding JSON info", err)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid JSON body: %s", err))

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(userLogin.Email, userLogin.Password)
	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error while login", err, zap.String("email", userLogin.Email))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Login successful", zap.String("email", userLogin.Email))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
