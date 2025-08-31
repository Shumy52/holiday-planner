package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Verifier *oidc.IDTokenVerifier
}

func NewAuth(issuer, audience string) (*Auth, error) {
	p, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		return nil, err
	}
	v := p.Verifier(&oidc.Config{ClientID: audience})
	return &Auth{Verifier: v}, nil
}

func (a *Auth) Require() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		raw := strings.TrimPrefix(h, "Bearer ")
		token, err := a.Verifier.Verify(c, raw)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// stash claims for handlers
		var claims map[string]any
		_ = token.Claims(&claims)
		c.Set("claims", claims)
		c.Next()
	}
}
