package server

import (
	"net/http"

	_ "github.com/tiagompalte/fake-payment-gateway/api"
	"github.com/tiagompalte/fake-payment-gateway/application"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/server/routes"
)

// @title						Fake Payment Gateway API
// @version						1.0
// @description					Fake Payment Gateway API
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url					http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url					http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath					/
// @schemes						http https
// @securityDefinitions.apikey 	apiKey
// @in 							header
// @name 						access_token
func NewServer(app application.App) *http.Server {
	groupRoutes := routes.CreateRoute(app)
	return app.Server().NewServer(groupRoutes)
}
