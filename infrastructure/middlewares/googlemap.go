package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

// GoogleMapClient GoogleMapクライアント
func GoogleMapClient() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := os.Getenv("GMAP_API_KEY")
			cl, err := maps.NewClient(maps.WithAPIKey(apiKey))
			if err != nil {
				logrus.Fatalf("Error creating GoogleMap client: %v", err)
			}

			c.Set("gmc", cl)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}

	}
}
