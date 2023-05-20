package route

import (
	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/handler"
)

func RegisterAllRoutes(e *echo.Echo) {
	handler.RegisterAuthRoutes(e)
	handler.RegisterUserRoutes(e)
	handler.RegisterLocationRoutes(e)
	handler.RegisterFaceShapeRoutes(e)
	handler.RegisterHealthCheckRoutes(e)
	handler.RegisterProductRoutes(e)
	handler.RegisterCategoryRoutes(e)
	handler.RegisterBrandRoutes(e)
	handler.RegisterQuestionnaireRoutes(e)
	handler.RegisterWishlistRoutes(e)
}
