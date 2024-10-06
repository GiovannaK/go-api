package controller

import (
	"net/http"
	"net/mail"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("FindUserByID function called", zap.String("journey", "FindUserByID"))
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error parsing UUID", err, zap.String("journey", "FindUserByID"))
		errorMessage := rest_err.NewBadRequestError("UserID must be a valid UUID")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	user, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error finding user by ID", err, zap.String("journey", "FindUserByID"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("User founded by ID", zap.String("journey", "FindUserByID"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(user))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("FindUserByEmail function called", zap.String("journey", "FindUserByEmail"))

	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error parsing email", err, zap.String("journey", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("Email must be a valid email address")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	user, err := uc.service.FindUserByEmailServices(email)

	if err != nil {
		logger.Error("Error finding user by email", err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User founded by email", zap.String("journey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(user))
}
