package auth

import (
	pkgzauth "github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/token"
)

var AuthService *pkgzauth.Service

func InitAuthService(baseURL string, authSecret string) {
	AuthService = pkgzauth.NewService(pkgzauth.Opts{
		SecretReader: token.SecretFunc(func(aud string) (string, error) {
			return authSecret, nil
		}),
	})
}
