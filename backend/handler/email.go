package handler

import (
	panichandler "backend/panicHandler"
	"backend/redis"
	"backend/mysql/table"
	"encoding/csv"
	"fmt"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	// "backend/redis"

	"backend/mysql"

	"github.com/jordan-wright/email"
)
var p *email.Pool
var pMux = new(sync.Mutex)

var fileAll = "工程文件(所有).csv"
var fileLogin = "工程文件(登入).csv"

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
	 
	v := Rand(100000, 999999)
	(*redis.Singleton()).InsertEmailCaptcha(emailAdd, v)
	
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

// 寄送每天的伺服器狀態
func SendDailyInfo(emailAdd string) bool {
	defer panichandler.Recover()
	emailAccount := os.Getenv("EMAIL_ACCOUNT")
	// emailAuth := os.Getenv("EMAIL_AUTH_TOKEN")
	em := email.NewEmail()
	
	em.From = emailAccount
	em.To = []string{emailAdd}
	 
	// title of email
	em.Subject = "work App 每日伺服器狀態"

	// 時間計算
	yesYear, yesMonth, yesDay := time.Now().AddDate(0, 0, -1).Date()
	yesterday := time.Date(yesYear, yesMonth, yesDay, 0, 0, 0, 0, time.Now().Location())
	todYear, todMonth, todDay := time.Now().Date()
	today := time.Date(todYear, todMonth, todDay, 0, 0, 0, 0, time.Now().Location())
	fmt.Print(yesterday, today)

	// get log data
	data := (*mysql.Singleton()).SelectLog(0, yesterday, today)
	dailyInfoIntoCSV(data, fileAll)
	dailyInfoIntoCSV(data, fileLogin)
	em.AttachFile(fileAll)
	em.AttachFile(fileLogin)

	//设置服务器相关的配置
	err := pool().Send(em, 10 * time.Second)
	if err != nil {
	   fmt.Println(err)
	   return false
	}
	fmt.Println("send successfully ... ")
	return true
}
func htmlDailyInfoBoard(data *[]table.Log) string {
	s := `
	<table>
		<thead>
			<tr>
				<td>姓名</td>
				<td>ip</td>
				<td>路由</td>
				<td>時間</td>
			</tr>
		</thead>
		<tbody>
	`
	for _, v := range *data {
		s += "<tr>"
		s += "<td>" + v.UserName + "</td>"
		s += "<td>" + v.Ip + "</td>"
		s += "<td>" + v.Routes + "</td>"
		s += "<td>" + v.CreateTime.Local().String() + "</td>"
		s += "</tr>"
	}
	return s + "</tbody></table>"
}

// 每日資訊寫入 csv
func dailyInfoIntoCSV (data *[]table.Log, fileName string) {
	csvFile, err := os.Create(fileName)

	if err != nil {
		fmt.Println("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	// 加入 column
	columns := []string{
		"id",
		"公司碼",
		"使用者",
		"路由",
		"ip位置",
		"參數",
		"創建時間",
	}
	csvwriter.Write(columns)

	// 去跑每一個 row
	for _, log := range *data {
		if fileName == fileLogin && log.Routes != "POST/workApp/entry/login" {
			continue
		}
		arr := []string{}
		arr = append(arr,
			strconv.FormatInt(log.LogId, 10),
			log.CompanyCode,
			log.UserName,
			log.Routes,
			log.Ip,
			log.Params,
			log.CreateTime.Local().String()[0: 20],
		)
		_ = csvwriter.Write(arr)
	}

	csvwriter.Flush()
	csvFile.Close()
}
func VerifyEmailFormat(email string) bool {
    pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配電子郵箱
	reg := regexp.MustCompile(pattern)
    return reg.MatchString(email)
}