package handler

import (
	"backend/table"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

//宣告JWT 結構
type Token struct {
    User table.UserTable
    Company table.CompanyTable
	jwt.StandardClaims
}

var SECRETKEY = os.Getenv("TOKEN_PASSWORD")

func(tokenStruct *Token) GetLoginToken() string {

	//宣告使用 HS256 與加入Payload 的聲明內容
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenStruct) //宣告使用 HS256 與加入Payload 的聲明內容

	//將 token_pwd 設定為 secret 並產生 jwt
	tokenString, _ := token.SignedString([]byte(SECRETKEY))

	return tokenString
}

//解析 token
func ParseToken(tokenString string) (jwt.MapClaims, any)  {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Don't forget to validate the alg is what you expect:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
        return []byte(SECRETKEY), nil
    })
    if err != nil {
        return jwt.MapClaims{}, "error"
    }
	
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}