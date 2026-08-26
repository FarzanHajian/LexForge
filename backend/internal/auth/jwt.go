// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v3"
	"github.com/auth0/go-jwt-middleware/v3/jwks"
	"github.com/auth0/go-jwt-middleware/v3/validator"
)

/* Builds an Auth0 JWT bearer-token validator for the given tenant
 * domain and API audience. The returned middleware validates a
 * token's signature (against Auth0's JWKS), issuer, audience and expiry;
 * it does not perform any authorization (scope/permission) checks. */
func NewJWTMiddleware(domain, audience string) (*jwtmiddleware.JWTMiddleware, error) {
	issuerUrl, err := url.Parse("https://" + domain + "/")
	if err != nil {
		return nil, err
	}

	provider, err := jwks.NewCachingProvider(
		jwks.WithIssuerURL(issuerUrl),
		jwks.WithCacheTTL(5*time.Minute),
	)
	if err != nil {
		return nil, err
	}

	jwtValidator, err := validator.New(
		validator.WithKeyFunc(provider.KeyFunc),
		validator.WithAlgorithm(validator.RS256),
		validator.WithIssuer(issuerUrl.String()),
		validator.WithAudience(audience),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return jwtmiddleware.New(
		jwtmiddleware.WithValidator(jwtValidator),
		jwtmiddleware.WithErrorHandler(handleAuthError),
	)
}

func handleAuthError(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Missing authorization token."})
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"message": "Invalid or expired token."})
}

/* Returns the validated token's subject ("sub" claim) from a request that has already
 * passed through the JWT middleware. It errors if no validated claims are present on
 * the request's context. */
func Subject(r *http.Request) (string, error) {
	claims, err := jwtmiddleware.GetClaims[*validator.ValidatedClaims](r.Context())
	if err != nil {
		return "", err
	}
	return claims.RegisteredClaims.Subject, nil
}
