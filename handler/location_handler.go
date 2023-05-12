package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterLocationRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	locationGroup := v1Group.Group("/location")
	locationGroup.GET("/provinces", getProvince)
	locationGroup.GET("/cities", getCity)               // province_id
	locationGroup.GET("/districts", getDistricts)       // city_id
	locationGroup.GET("/subdistricts", getSubDistricts) // district_id
	locationGroup.GET("/postal_codes", getPostalCode)   // subdistrict_id
}

func getProvince(c echo.Context) error {
	data, statusCode, err := repository.GetProvinces()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func getCity(c echo.Context) error {
	provinceID := utils.StrToUint64(c.QueryParam("province_id"), 0)

	data, statusCode, err := repository.GetCities(provinceID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func getDistricts(c echo.Context) error {
	cityID := utils.StrToUint64(c.QueryParam("city_id"), 0)

	data, statusCode, err := repository.GetDistricts(cityID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func getSubDistricts(c echo.Context) error {
	districtID := utils.StrToUint64(c.QueryParam("district_id"), 0)

	data, statusCode, err := repository.GetSubDistricts(districtID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}

func getPostalCode(c echo.Context) error {
	subDistrictID := utils.StrToUint64(c.QueryParam("subdistrict_id"), 0)

	data, statusCode, err := repository.GetPostalCodes(subDistrictID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success get data", data, http.StatusOK)
}
