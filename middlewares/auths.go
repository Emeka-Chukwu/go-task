package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"go-task/token"
	"go-task/util"

	"fmt"
	"net/http"

	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker, config util.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(authorizationHeaderKey)
		if authorizationHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifiyToken(accessToken)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		tokenDiff := payload.ExpiredAt.Sub(payload.IssuedAt).Minutes() + 0.005
		if (tokenDiff - config.AccessTokenDuration.Minutes()) > 1 {
			err := errors.New("invalid token")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}
		reqWithStore := r.WithContext(context.WithValue(r.Context(), authorizationPayloadKey, payload))
		next.ServeHTTP(w, reqWithStore)
	})

}

func GetCurrentUserID(ctx context.Context) (*token.Payload, bool) {
	payload, ok := ctx.Value(authorizationPayloadKey).(*token.Payload)
	return payload, ok
}
