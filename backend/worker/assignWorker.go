package worker

import (
	"backend/logger"
	"backend/mysql"
	"backend/panicHandler"
	"backend/restFul/service"
	"backend/mysql/table"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)
var Log = logger.Logger()
var bd []byte
func AssignWorker(routerMethod int) func(props *gin.Context) {
	defer panichandler.Recover()
	routeFunc := service.FindSingleUser

	switch routerMethod {
	case 0:
		//get method of fetch single user
		routeFunc = service.FindSingleUser
		break
	case 1:
		// get method of fetch my data
		routeFunc = service.FindMine
		break
	case 2:
		//put method of update user
		routeFunc = service.UpdateUser
		break
	case 3:
		// post method of login
		routeFunc = service.Login
		break
	case 4:
		//post method of register
		routeFunc = service.Register
		break
	case 5:
		// get method of check user access
		routeFunc = service.CheckAccess
		break
	case 6:
		// post method of send email captcha
		routeFunc = service.EmailCaptcha
		break;
	case 7:
		// get method of fetch all user data
		routeFunc = service.GetAllUser
		break
	case 8:
		// get method of getch banch all
		routeFunc = service.FetchBanchAll
		break
	case 9:
		// get method of fetch banch style
		routeFunc = service.FetchBanchStyle
		break
	case 11:
		// post method of update banch style
		routeFunc = service.UpdateBanchStyle
		break
	case 13:
		// put method of insert banch style
		routeFunc = service.InsertBanchStyle
		break
	case 15:
		// put method of insert banch
		routeFunc = service.InsertBanch
		break
	case 16:
		// post method of update banch
		routeFunc = service.UpdateBanch
		break
	case 17:
		// get method of fetching company info
		routeFunc = service.FetchCompany
		break
	case 18:
		// post method of update company
		routeFunc = service.UpdateCompany
		break
	case 19:
		// post method of update password
		routeFunc = service.ChangePassword
		break
	case 20:
		// delete method of delete banch style
		routeFunc = service.DeleteBanchStyle
		break
	case 22:
		// post method of update myself user data
		routeFunc = service.UpdateMine
		break
	case 23:
		// put method of insert company
		routeFunc = service.InsertCompany
		break
	case 24:
		// post method of update password
		routeFunc = service.ForgetPassword
		break
	case 25:
		routeFunc = service.DeleteBanch
		break
	case 26:
		// get method of fetch wait reply
		routeFunc = service.FetchWaitReply
		break
	case 27:
		// post method of update wait company reply
		routeFunc = service.UpdateWaitCompanyReply
		break
	case 28:
		// put method of insert wait company reply
		routeFunc = service.InsertWaitCompanyReply
		break
	case 32:
		// get method of fetch workTime
		routeFunc = service.FetchWorkTime
		break
	case 33:
		// put method of insert workTime
		routeFunc = service.InsertWorkTime
		break
	case 34:
		// post method of update workTime
		routeFunc = service.UpdateWorkTime
		break
	case 35:
		// delete method of delete workTime
		routeFunc = service.DeleteWorkTime
		break
	case 36:
		// get method of fetch paidVocation
		routeFunc = service.FetchPaidVocation
		break
	case 37:
		// put method of insert paidVocation
		routeFunc = service.InsertPaidVocation
		break
	case 38:
		// post method of update paidVocation
		routeFunc = service.UpdatePaidVocation
		break
	case 39:
		// delete method of delete paidVocation
		routeFunc = service.DeletePaidVocation
		break
	case 40:
		routeFunc = service.FetchPerformance
		break
	case 41:
		routeFunc = service.UpdatePerformance
		break
	case 42:
		routeFunc = service.InsertPerformance
		break
	case 43:
		routeFunc = service.DeletePerformance
		break
	case 44:
		routeFunc = service.GetGoogleOAuth
		break
	case 45:
		routeFunc = service.LoginGoogle
		break
	case 46:
		routeFunc = service.CopyPerformance
		break
	case 47:
		routeFunc = service.InsertUser
		break
	case 48:
		routeFunc = service.FetchYearPerformance
		break
	default:
		break;
	}
	return func (props *gin.Context)  {
		// 紀錄參數
		params := "params => "
		for i, v := range (*props).Request.URL.Query() {
			if i != "token" {
				params += fmt.Sprintf("%s : %s ,", i, v)	
			}
		}

		// 紀錄 body
		bd, _ = ioutil.ReadAll(props.Request.Body)
		body := make(map[string]interface{})
		json.Unmarshal(bd, &body)
		bodyString := "body => "
		for i, v := range body {
			if i != "token" {
				bodyString += fmt.Sprintf("%s : %s ,", i, v)
			}
		}
		// 儲存回去
		props.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bd))

		// 儲存進入資料庫
		user, company, err := service.CheckUserAndCompany(props)
		if err {return}
		now := time.Now()
		logStruct := table.Log{
			UserId: user.UserId,
			UserName: user.UserName,
			CompanyId: company.CompanyId,
			CompanyCode: company.CompanyCode,
			Permession: user.Permession,
			Ip: (*props).ClientIP(),
			Params: params + bodyString,
			Routes: (*props).Request.Method + (*props).Request.URL.Path,
			CreateTime: now,
			LastModify: now,
		}
		// 紀錄log
		(*mysql.Singleton()).InsertLog(&logStruct)
		// 分配
		waitJob := new(sync.WaitGroup)
		waitJob.Add(1)
		(*props).Writer.Header().Set("Transfer-Encoding", "chunked")
		// fmt.Println(*props)
		(*WorkerSingleton()).JobChan <- func ()  {
			routeFunc(props, waitJob)
		}
		waitJob.Wait()
	}
}