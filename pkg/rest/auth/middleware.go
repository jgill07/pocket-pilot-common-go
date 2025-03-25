package auth

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

var (
	cacheTTL = 1 * time.Hour
)

func NewAuthMiddleware(config *Config) (gin.HandlerFunc, error) {
	if len(config.IssueUrl) == 0 || len(config.Audiences) == 0 {
		return nil, errors.New("Audience and Issuer URL are required")
	}

	parsedIssuerUrl, err := url.Parse(config.IssueUrl)
	if err != nil {
		return nil, err
	}

	provider := jwks.NewCachingProvider(parsedIssuerUrl, cacheTTL)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		parsedIssuerUrl.String(),
		config.Audiences,
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &PocketUser{}
			},
		),
		// Allow for a 1 minute clock skew
		// E.g. if the system clock is off, this will allow for a 1 minute leeway
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	// TODO: Add custom error handler
	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return adapter.Wrap(middleware.CheckJWT), nil
}

func PocketUserFromContext(ctx context.Context) (*PocketUser, error) {
	verifyUserErr := fmt.Errorf("unable to verify user")
	validClaims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(validator.ValidatedClaims)
	if !ok {
		return nil, verifyUserErr
	}
	user, ok := validClaims.CustomClaims.(*PocketUser)
	if !ok {
		return nil, verifyUserErr
	}
	user.Subject = validClaims.RegisteredClaims.Subject
	return user, nil
}
