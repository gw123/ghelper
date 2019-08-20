package gsafe

import (
	"github.com/dgrijalva/jwt-go"
	"errors"
	"crypto/rsa"
	"encoding/json"
	"strings"
	"io/ioutil"
	"fmt"
)

func ParseJwtToken(jwtToken string, secret string) (*jwt.StandardClaims, error) {
	if jwtToken == "" {
		return nil, errors.New("octoken is nil")
	}

	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		err = claims.Valid()
		if err != nil {
			return nil, err
		}

		return claims, nil
	}

	return nil, errors.New("token is invalid")
}

func MakeJwtToken(claims *jwt.StandardClaims, secret string) (string, error) {
	if claims == nil {
		return "", errors.New("MakeJwtToken claims  is nil")
	}

	mySigningKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func MakeRSAJwtToken(claims *jwt.StandardClaims, alg string, privateKey *rsa.PrivateKey) (string, error) {
	if claims == nil {
		return "", errors.New("MakeJwtToken claims  is nil")
	}
	claimsStr, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	method := jwt.GetSigningMethod(alg)
	return method.Sign(string(claimsStr), privateKey)
}

func ParseRSAJwtToken(jwtToken string, alg string, publicKey *rsa.PublicKey) (*jwt.StandardClaims, error) {
	method := jwt.GetSigningMethod(alg)
	parts := strings.Split(jwtToken, ".")
	fmt.Println("parts",len(parts))
	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		err = claims.Valid()
		if err != nil {
			return nil, err
		}

		return claims, nil
	}
	return nil, errors.New("token is invalid")
}

func LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		return nil, e
	}

	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		return nil, e
	}
	return key, nil
}

func LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		return nil, e
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		return nil, e
	}
	return key, nil
}
