package handler

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

//宣告JWT 結構
type Token struct {
	UserId int64
	Account string
	jwt.StandardClaims
}

func(tokenStruct *Token) GetLoginToken() string {

	//宣告使用 HS256 與加入Payload 的聲明內容
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenStruct) //宣告使用 HS256 與加入Payload 的聲明內容

	//將 token_pwd 設定為 secret 並產生 jwt
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	return tokenString
}
