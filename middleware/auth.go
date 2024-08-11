package middleware

import (
  "context"
  "go-jwt/helpers"
  "net/http"
)

func Auth(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    accessToken := r.Header.Get("Authorization")

    if accessToken == "" {
      helpers.Response(w, http.StatusUnauthorized, "Authorization is not found", nil)
      return
    }

    user, err := helpers.ValidateToken(accessToken)
    if err != nil {
      helpers.Response(w, http.StatusUnauthorized, "Authorization is invalid", nil)
      return
    }

    ctx := context.WithValue(r.Context(), "userinfo", user)

    next.ServeHTTP(w, r.WithContext(ctx))
  })
}
