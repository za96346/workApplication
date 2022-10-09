package handler

import (
	panichandler "backend/panicHandler"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"sync"
	"time"

	// "backend/redis"

	"github.com/jordan-wright/email"
)
var p *email.Pool
var pMux = new(sync.Mutex)

func pool() *email.Pool {
	if p == nil {
		pMux.Lock()
		defer pMux.Unlock()
		if p == nil {
			
			p, _ =  email.NewPool(
				"smtp.gmail.com:587",
				3,    // 数量设置成 3 个
				smtp.PlainAuth("", os.Getenv("EMAIL_ACCOUNT"), os.Getenv("EMAIL_AUTH_TOKEN"), "smtp.gmail.com"),
			)
			
		}
	}
	return p
}

func SendEmail(emailAdd string) bool {
	defer panichandler.Recover()
	emailAccount := os.Getenv("EMAIL_ACCOUNT")
	// emailAuth := os.Getenv("EMAIL_AUTH_TOKEN")
	em := email.NewEmail()
	
	em.From = emailAccount
	em.To = []string{emailAdd}
	 
	// title of email
	em.Subject = "work App 電子信箱驗證"
	 
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(999999)
	(*Singleton()).Redis.InsertEmailCaptcha(emailAdd, v)
	
	em.HTML = []byte(htmlBoard(v))
	 
	//设置服务器相关的配置
	err := pool().Send(em, 10 * time.Second)
	if err != nil {
	   fmt.Println(err)
	   return false
	}
	fmt.Println("send successfully ... ")
	return true
}
func htmlBoard(number int) string {
	return fmt.Sprintf(`
		<div>驗證碼為<span style="color: blue;">%d</span>  時效為三分鐘 請儘速完成註冊</div>
			<a href='#'>前往登入註冊頁面</a>`, number)
}