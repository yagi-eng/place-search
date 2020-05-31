package middlewares

import (
	"virtual-travel/infrastructure/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// DatabaseClient DBクライアント
type DatabaseClient struct {
	DB *gorm.DB
}

// DatabaseService DBサービス
func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := database.Connect()
			d := DatabaseClient{DB: session}

			defer d.DB.Close()

			// output sql query
			d.DB.LogMode(true)

			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
