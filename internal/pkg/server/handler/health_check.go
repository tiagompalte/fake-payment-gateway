package handler

import (
	"net/http"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

// @Summary Health Check
// @Description Verify health check application
// @Tags Health Check
// @Produce json
// @Success 204
// @Router /api/health-check [get]
func HealthCheckHandler(healthCheck usecase.HealthCheckUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		err := healthCheck.Execute(ctx)
		if err != nil {
			return errors.Wrap(err)
		}

		server.RespondNoContent(w)

		return nil
	}
}
