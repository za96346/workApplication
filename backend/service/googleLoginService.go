package service

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
var redirectURL = os.Getenv("REDIRECT_URL")
var clientId = os.Getenv("CLIENT_ID")
var ClientSecret = os.Getenv("CLIENT_SECRET")
var googleTokenRoute = os.Getenv("GOOGLE_TOKEN_ROUTE")
func CreateGoogleOAuthURL() string {
	// 使用 lib 產生一個特定 config instance
	config := &oauth2.Config{
		//憑證的 client_id
		ClientID: clientId,
		//憑證的 client_secret
		ClientSecret: ClientSecret,
		//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
		RedirectURL:  redirectURL,
		//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		//指的是 Google auth server 的 endpoint，用 lib 預設值即可
		Endpoint: google.Endpoint,
	}
   //產生出 config instance 後，就可以使用 func AuthCodeURL 建立請求網址
   return config.AuthCodeURL(uuid.New().String())
}
func GetGoogleOAuth(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	redirectURL := CreateGoogleOAuthURL()
	props.JSON(http.StatusOK, gin.H{
		"data": redirectURL,
		"message": StatusText().FindSuccess,
	})
}
func LoginGoogle (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	req := new(struct {
		Code string `json:"Code"`
	})
	(*props).ShouldBindJSON(req)

	postData := url.Values{}
	postData.Add("client_id", clientId)
	postData.Add("client_secret", ClientSecret)
	postData.Add("code", req.Code)
	postData.Add("grant_type", "authorization_code")
	postData.Add("redirect_uri", redirectURL)

	response, err := http.Post(googleTokenRoute, "application/x-www-form-urlencoded", strings.NewReader(postData.Encode()))
	Body, _ := ioutil.ReadAll(response.Body)
	(*props).JSON(response.StatusCode, gin.H {
		"body": Body,
		"err": err,
	})
}