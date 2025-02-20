package middleware

import (
	contextNative "context"
	"net/http"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/server/constant"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

func ValidateExtractAccountTokenMiddleware(header string, findAccountByTokenUseCase usecase.FindAccountByTokenUseCase) server.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(header)
			if len(token) == 0 || token == "" {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			ctx := r.Context()

			account, err := findAccountByTokenUseCase.Execute(ctx, token)
			if err != nil {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			ctx = contextNative.WithValue(ctx, constant.ContextAccount, account)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
