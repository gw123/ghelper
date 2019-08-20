package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gw123/ghelper/gsafe"
	"fmt"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
	"crypto/rsa"
	"errors"
	"strconv"
)

func token() {
	key := "xytschool"
	var timeout int64 = 3600 * 24 * 100
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = time.Now().Unix() + timeout
	claims.Subject = `{"name":"client01","version":"1.0.1"}`
	claims.Issuer = "xytschool.com"
	claims.IssuedAt = time.Now().Unix()

	token, err := gsafe.MakeJwtToken(claims, key)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("token : ", token)
	}

	res, err := gsafe.ParseJwtToken(token, key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("res   : ", res)
	}
}

func t2() {
	var timeout int64 = 3600 * 24 * 100
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = time.Now().Unix() + timeout
	claims.Subject = `{"name":"client01","version":"1.0.1"}`
	claims.Issuer = "xytschool.com"
	claims.IssuedAt = time.Now().Unix()

	privatekey, err := gsafe.LoadRSAPrivateKeyFromDisk("./keys/sample_key")
	if err != nil {
		fmt.Println("LoadRSAPrivateKeyFromDisk", err)
		return
	}
	//fmt.Println("privatekey : ", privatekey)
	jwt_token, err := gsafe.MakeRSAJwtToken(claims, "RS256", privatekey)
	if err != nil {
		fmt.Println("MakeRSAJwtToken", err)
		return
	}
	fmt.Println("jwt_token: ", jwt_token)
	fmt.Println("================")

	pubkey, err := gsafe.LoadRSAPublicKeyFromDisk("./keys/sample_key.pub")
	if err != nil {
		fmt.Println("LoadRSAPublicKeyFromDisk", err)
		return
	}
	fmt.Println("pubkey : ", pubkey)
	claims1, err := gsafe.ParseRSAJwtToken(jwt_token, "RS256", pubkey)
	if err != nil {
		fmt.Println("ParseRSAJwtToken", err)
		return
	}
	fmt.Println("claims: ", claims1)
	fmt.Println("================")
}

func parseJwsToken(token []byte, pubkey *rsa.PublicKey) (*jwt.StandardClaims, error) {
	jwstoken, err := jws.Parse(token)
	if err != nil {
		return nil, err
	}
	err = jwstoken.Verify(pubkey, crypto.SigningMethodRS512)
	if err != nil {
		return nil, err
	}
	maps, ok := jwstoken.Payload().(map[string]interface{})
	if !ok {
		return nil, errors.New("转换失败")
	}

	sub := strconv.FormatFloat(maps["sub"].(float64), 'f', -1, 64)
	authInfo := &jwt.StandardClaims{
		Audience: maps["aud"].(string),
		Subject:  sub,
		Issuer:   maps["iss"].(string),
	}
	return authInfo, nil
}

func parseJswtokenExample() {
	token2 := `eyJhbGciOiJSUzUxMiJ9.eyJhdWQiOiJwb3MiLCJleHAiOjE1NjIyNTAwNTIsImlzcyI6Im9jdG9wYXNzIiwiamlkIjoiYmtmMGs2NWNkY2tncHBjaHR0aWciLCJzdWIiOjM5NX0.yNJJZlpbJsXJkJR_RclLyqWorzpOr7GFdvzx9-0TrdVJ46iqIZBx8D325brCiHtPUCKvAv2iHO83iLX_Xotb4SOOHuhnb1b0yQGDyP8WTsOaTl4cJuShQX2gdGJwpuX-Q_gjPYwYx0-P1rV2Xwkjret3OnC7rdNRK9PPdDZHD1Rlu5hvc_u5E9NN2TsFk0Yw2SDhyGlm2v7uppWmpKCWrqbURmIJ2o7jryLeoAHMNqVk5LlJhoRfT2ZpUeNW-9xjzEr_TWeRLxBBEJBdTlMyDafn7hkjMifjhi-w3w8mFZpp0vp6scBsSOPwScJg9wnC7zQ7Sai2d58Um9N554O3STZuay-NkbLqbJml3DHDgxQdFMl3pWBIDSAU0QQ0ndg8X2JWTUK2zCLMhdihYSLunPgZhNVnKaioVt4cUKpQbCmQwsdwuaqXSo1Gy-c1V9tJTlTlbkMlmtBYNHq2IJWamNc8gwefomQ20KCUm15r0S-OAGWnyZv9wz5h4dJWb629T_kl_yDQ6suE93YESf1AmyIsU9bzv1NmgQ2Nv7bGstvlF7DK6FzqBVX1PzgkxA_ZgEKA1HjrU09jjmdr1yfA52PjP7z6ieR59ja0xBOuuRTQQmNWaDz6QO1D-iwThrW9gLFgkv5Crs4UQK4ePTn4g--I0Sccv8hwgttTfcgspbg`
	//gsafe.ParseRSAJwtToken(string(token),"RS256",)
	pubkey, err := gsafe.LoadRSAPublicKeyFromDisk("./keys/occ.pem")
	if err != nil {
		fmt.Println("LoadRSAPublicKeyFromDisk", err)
		return
	}

	claims2, err := parseJwsToken([]byte(token2), pubkey)
	if err != nil {
		fmt.Println("parseJwsToken ", err)
		return
	}
	fmt.Println(claims2)
}

func main() {
	//privatekey, err := gsafe.LoadRSAPrivateKeyFromDisk("./keys/sample_key")
	//if err != nil {
	//	fmt.Println("LoadRSAPrivateKeyFromDisk", err)
	//	return
	//}
	//
	//claims := jws.Claims{
	//	"sub": "123",
	//	"iss": "xyt",
	//	"aud": "xxxx",
	//	"jid": xid.New().String(),
	//	"exp": time.Now().Unix() + 1000,
	//}
	//
	//t := jws.NewJWT(claims, crypto.SigningMethodRS256)
	//token, err := t.Serialize(privatekey)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(token))

}
