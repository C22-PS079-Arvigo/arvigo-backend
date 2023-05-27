package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterProductRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	v1Group.GET("/product-recommendation", getRecommendationProduct, middleware.ApiKeyMiddleware)

	productGroup := v1Group.Group("/products", middleware.AuthMiddleware)
	initialProductGroup := productGroup.Group("/initials")
	initialProductGroup.GET("/:id", getInitalProductByCategoryID)             // TODO: develop
	initialProductGroup.GET("/marketplace/:id", getInitalProductByCategoryID) // TODO: develop

	initialProductGroup.POST("", createInitialProductHandler)
	initialProductGroup.GET("/category/:id", getInitalProductByCategoryID)

	merchantProductGroup := productGroup.Group("/merchants")
	merchantProductGroup.POST("", createMerchantProductHandler)
	merchantProductGroup.PUT("", updateMerchantProduct)
	merchantProductGroup.PUT("/verify", verifyMerchantProduct)
}

func createInitialProductHandler(c echo.Context) error {
	var data datastruct.CreateInitialProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, "Failed to parse form data", nil, http.StatusBadRequest)
	}

	images := form.File["images"]
	if len(images) == 0 {
		return utils.ResponseJSON(c, "Images must be filled", nil, http.StatusBadRequest)
	}

	data.Images = images
	statusCode, err := repository.CreateInitialProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed create product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product created", nil, statusCode)
}

func createMerchantProductHandler(c echo.Context) error {
	var data datastruct.CreateMerchantProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, "Failed to parse form data", nil, http.StatusBadRequest)
	}

	images := form.File["images"]
	if len(images) == 0 {
		return utils.ResponseJSON(c, "Images must be filled", nil, http.StatusBadRequest)
	}

	data.Images = images

	statusCode, err := repository.CreateMerchantProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed create product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product created", nil, statusCode)
}

func getRecommendationProduct(c echo.Context) error {
	data, statusCode, err := repository.GetProductRecommendationMachineLearningDummy()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func getInitalProductByCategoryID(c echo.Context) error {
	categoryID := utils.StrToUint64(c.Param("id"), 0)
	if categoryID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	data, statusCode, err := repository.GetInitialProductByCategoryID(categoryID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}

func verifyMerchantProduct(c echo.Context) error {
	var data datastruct.VerifyProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.VerifyMerchantProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed update product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product updated", nil, statusCode)
}

func updateMerchantProduct(c echo.Context) error {
	var data datastruct.UpdateProductInput
	if err := c.Bind(&data); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(data)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.UpdateMerchantProduct(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed update product", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Product updated", nil, statusCode)
}
