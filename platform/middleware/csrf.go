package middleware

import (
	"crypto/sha1"
	"encoding/base64"
	"io"

	"github.com/gin-gonic/gin"
)

type CSRFMiddleware struct {
	secret string
	token  string
}

func NewCSRFMiddleware(secret string) *CSRFMiddleware {
	return &CSRFMiddleware{
		secret: secret,
		token:  tokenize(secret),
	}
}

func tokenize(secret string) string {
	h := sha1.New()
	io.WriteString(h, secret)
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return hash
}

func (m *CSRFMiddleware) GetToken() string {
	return m.token
}

func (m *CSRFMiddleware) ValidateCSRF() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("X-CSRF-TOKEN")

		if token != m.token {
			ctx.JSON(400, "CSRF token mismatch")
            ctx.Abort()
		}

		ctx.Next()
	}
}
