package middlewares

import "github.com/labstack/echo/v4"

func Validate(schema interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := c.Bind(schema); err != nil {
				return err
			}

			if err := c.Validate(schema); err != nil {
				return err
			}

			c.Set("payload", schema)

			return next(c)
		}
	}
}
