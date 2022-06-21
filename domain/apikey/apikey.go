package apikey

import (
	"errors"
	"fmt"
	"time"
	"yawaraka-tissue/domain/problem"

	"github.com/golang-jwt/jwt"
)

type Verifier struct {
	secrets        map[string]string
	allowedIssuers []string
}

type Encoder struct {
	secrets map[string]string
	issuer  string
}

func NewVerifier(secrets map[string]string, issuers []string) *Verifier {
	return &Verifier{
		secrets:        secrets,
		allowedIssuers: issuers,
	}
}

func NewEncoder(secrets map[string]string, issuer string) *Encoder {
	return &Encoder{
		secrets: secrets,
		issuer:  issuer,
	}
}

func (v *Verifier) Verify(signedString string) error {
	if signedString == "" {
		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail("Auth header is empty")
	}

	token, err := jwt.Parse(signedString, func(token *jwt.Token) (any, error) {
		kid, ok := token.Header["kid"].(string)

		if !ok {
			return nil, problem.NewUnauthorized(
				problem.TypeUnAuthorized,
			).WithDetail("Kid is not found in token's header")
		}

		if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, problem.NewUnauthorized(
				problem.TypeUnAuthorized,
			).WithDetail(fmt.Sprintf("Requested alg is `%v`, but not supported.", token.Header["alg"]))
		}

		key, ok := v.secrets[kid]
		if !ok {
			return nil, problem.NewUnauthorized(
				problem.TypeUnAuthorized,
			).WithDetail(fmt.Sprintf("Requested kid is `%v`, but not found.", kid))
		}

		return []byte(key), nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return problem.NewUnauthorized(
					problem.TypeUnAuthorized,
				).WithDetail("Token is expired").Wrap(err)
			}
		}

		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail("Sign is invalid").Wrap(err)
	}

	if token == nil {
		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail("Missing token in signed string.")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail("Missing claims in token")
	}

	_, ok = claims["iat"].(float64)
	if !ok {
		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail("Missing iat in token")
	}

	iss, ok := claims["iss"].(string)
	if !ok {
		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail("Missing iss in token")
	}

	if !v.isAllowedIssuer(iss) {
		return problem.NewUnauthorized(
			problem.TypeUnAuthorized,
		).WithDetail(fmt.Sprintf("Issuer `%v` is not found in allowed list", iss))
	}

	return nil
}

func (v *Verifier) isAllowedIssuer(iss string) bool {
	for _, allowed := range v.allowedIssuers {
		if allowed == iss {
			return true
		}
	}

	return false
}

func (e *Encoder) Encode(kid string, validMinutes int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": e.issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(validMinutes) * time.Minute).Unix(),
	})
	token.Header["kid"] = kid

	key, ok := e.secrets[kid]
	if !ok {
		return "", fmt.Errorf("Failed to find secret kid=%v", kid)
	}

	return token.SignedString([]byte(key))
}
