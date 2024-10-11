package controller

import (
	"fmt"
	"net/http"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/controller/model/request"
	"github.com/GiovannaK/go-api/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("UpdateUser function called", zap.String("journey", "UpdateUser"))
	var userUpdateRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		logger.Error("Error while binding JSON info", err)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid JSON body: %s", err))

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid ID, must be a hexadecimal")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(userUpdateRequest.Name, userUpdateRequest.Age)
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error while trying to update user", err, zap.String("journey", "UpdateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully", zap.String("journey", "UpdateUser"), zap.String("userId", userId))

	c.Status(http.StatusNoContent)
}
