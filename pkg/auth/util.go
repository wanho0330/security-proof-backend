package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/Netflix/go-env"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"security-proof/pkg/constants"
)

var jwtRole = "role"

// jwtConfig struct composed of a header, a secret key, an access token time and a refresh token time.
type jwtConfig struct {
	Header           string        `env:"JWT_HEADER,default=Bearer "`
	SecretKey        string        `env:"JWT_SECRET_KEY,default=secret-key"`
	AccessTokenTime  time.Duration `env:"JWT_ACCESS_TOKEN_EXPIRED,default=1h"`
	RefreshTokenTime time.Duration `env:"JWT_REFRESH_TOKEN_EXPIRED,default=72h"`
}

// parseToken function is returning an index, a role and an error, accepting signed token.
func parseToken(signedToken string) (idx string, role int32, err error) {
	config := jwtConfig{}
	_, err = env.UnmarshalFromEnviron(&config)
	if err != nil {
		return "", 0, errors.Join(constants.ErrTokenParse, err)
	}

	key := []byte(config.SecretKey)
	token, err := jwt.Parse([]byte(signedToken), jwt.WithKey(jwa.HS256, key))
	if err != nil {
		return "", 0, errors.Join(constants.ErrTokenParse, err)
	}

	roleAny, _ := token.Get(jwtRole)

	return token.Subject(), int32(roleAny.(float64)), nil
}

// Pint32ToStr function is returning a string accepting an int32 pointer.
func Pint32ToStr(pint *int32) string {
	if pint == nil {
		return ""
	}

	return strconv.Itoa(int(*pint))
}

// StrToInt32 function is returning an int32 accepting a string.
func StrToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}

	return int32(i)
}
