package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func NewMiddleware(f *fiber.App) {
	//f.Use(middleware.Recover())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: strings.Split(util.Env("CLIENT_WEB_URL").GetString("*"), ","),
	//	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	//	AllowHeaders: []string{
	//		"Access-Control-Allow-Credentials",
	//		"Access-Control-Allow-Headers",
	//		"Content-Type",
	//		"Content-Length",
	//		"Accept-Encoding",
	//		"Authentication",
	//	},
	//	AllowCredentials: false,
	//	MaxAge:           24 * int(time.Hour),
	//}))

	//f.Use(middleware.RequestID())
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: logger.AccessLogFormat,
	//	Output: logger.NewAccessLogger(),
	//}))

	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		reqID := c.Response().Header().Get(echo.HeaderXRequestID)
	//		c.Set("RequestID", reqID)
	//		logger.With("request_id", reqID)
	//		return next(c)
	//	}
	//})
}
