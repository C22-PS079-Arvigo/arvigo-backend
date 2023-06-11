package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterSubscriptionRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	subsGroup := v1Group.Group("/subscription", middleware.AuthMiddleware)

	subsGroup.POST("/user", userCreatePayment)
	subsGroup.POST("/merchant", partnerCreatePayment)

	subsGroup.PUT("/user/verify/:id", verifyPaymentUser)
	subsGroup.PUT("/merchant/verify/:id", verifyPaymentMerchant)

	subsGroup.GET("/user", getAllUserPayment)
	subsGroup.GET("/merchant", getAllMerchantPayment)
}

func getAllUserPayment(c echo.Context) error {
	user, statusCode, err := repository.GetListPaymentUser()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", user, statusCode)
}

func getAllMerchantPayment(c echo.Context) error {
	user, statusCode, err := repository.GetListPaymentMerchant()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", user, statusCode)
}

func userCreatePayment(c echo.Context) error {
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)
	var data datastruct.UserCreatePaymentInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.UserCreatePayment(userAuth.ID, data)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}
	return utils.ResponseJSON(c, "Success", nil, http.StatusOK)
}

func partnerCreatePayment(c echo.Context) error {
	userAuth := c.Get("userAuth").(*datastruct.UserAuth)
	var data datastruct.PartnerCreatePaymentInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.PartnerCreatePayment(userAuth.ID, data)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}
	return utils.ResponseJSON(c, "Success", nil, http.StatusOK)
}

func verifyPaymentUser(c echo.Context) error {
	pID := utils.StrToUint64(c.Param("id"), 0)
	if pID == 0 {
		return utils.ResponseJSON(c, "invalid subscription id", nil, http.StatusBadRequest)
	}
	var data datastruct.VerifyPaymentUser
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.VerifyPaymentUser(pID, data)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}
	return utils.ResponseJSON(c, "Success", nil, http.StatusOK)
}

func verifyPaymentMerchant(c echo.Context) error {
	pID := utils.StrToUint64(c.Param("id"), 0)
	if pID == 0 {
		return utils.ResponseJSON(c, "invalid subscription id", nil, http.StatusBadRequest)
	}
	var data datastruct.VerifyPaymentMerchant
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.VerifyPaymentMerchant(pID, data)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}
	return utils.ResponseJSON(c, "Success", nil, http.StatusOK)
}
