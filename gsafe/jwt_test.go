package gsafe

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"testing"
	"time"
)

func TestMakeJwtToken(t *testing.T) {

	claims := jwt.StandardClaims{
		Audience:  "mobile",
		ExpiresAt: time.Now().Unix() + 3600*24*7,
		Id:        strconv.Itoa(123),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "xytschool",
		Subject:   string(""),
	}

	MakeJwtToken(&claims, "123456")
}
