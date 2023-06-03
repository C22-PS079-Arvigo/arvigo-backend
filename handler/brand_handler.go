package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/middleware"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterBrandRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	brandGroup := v1Group.Group("/brands", middleware.AuthMiddleware)
	brandGroup.GET("", getBrands)
	brandGroup.POST("", createBrand)
	brandGroup.PUT("/:id", updateBrand)
	brandGroup.GET("/category/:id", getBrandByCategory)

	brandGroup.GET("/:id/list-product", getListProductByBrand)

}

func getBrands(c echo.Context) error {
	data, statusCode, err := repository.GetBrands()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func getBrandByCategory(c echo.Context) error {
	categoryID := utils.StrToUint64(c.Param("id"), 0)
	if categoryID == 0 {
		return utils.ResponseJSON(c, "Invalid category ID", nil, http.StatusBadRequest)
	}

	data, statusCode, err := repository.GetBrandByCategory(categoryID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func createBrand(c echo.Context) error {
	var data datastruct.BrandInput
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

	images := form.File["image"]
	if len(images) == 0 {
		return utils.ResponseJSON(c, "Images must be filled", nil, http.StatusBadRequest)
	}

	data.Image = images[0]
	statusCode, err := repository.CreateBrand(data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed create brand", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Brand created", nil, statusCode)
}

func updateBrand(c echo.Context) error {
	brandID := utils.StrToUint64(c.Param("id"), 0)
	if brandID == 0 {
		return utils.ResponseJSON(c, "Invalid brand ID", nil, http.StatusBadRequest)
	}

	var data datastruct.BrandInput
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

	data.Image = images[0]
	statusCode, err := repository.UpdateBrand(brandID, data)
	if err != nil {
		return utils.ResponseJSON(c, "Failed update brand", err.Error(), statusCode)
	}

	return utils.ResponseJSON(c, "Brand updated", nil, statusCode)
}

func getListProductByBrand(c echo.Context) error {
	brandID := utils.StrToUint64(c.Param("id"), 0)
	if brandID == 0 {
		return utils.ResponseJSON(c, "Invalid brand ID", nil, http.StatusBadRequest)
	}

	data, statusCode, err := repository.GetListProductByBrand(brandID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, statusCode)
}
