package worker

import (
	"backend/panicHandler"
	"backend/service"
	"sync"

	"github.com/gin-gonic/gin"
)


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
	case 10:
		// get method of fetch banch rule
		routeFunc = service.FetchBanchRule
		break
	case 11:
		// post method of update banch style
		routeFunc = service.UpdateBanchStyle
		break
	case 12:
		// post method of update banch rule
		routeFunc = service.UpdateBanchRule
		break
	case 13:
		// put method of insert banch style
		routeFunc = service.InsertBanchStyle
		break
	case 14:
		// put method of insert banch rule
		routeFunc = service.InsertBanchRule
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
	case 21:
		// delete method of delete banch rule
		routeFunc = service.DeleteBanchRule
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
	default:
		break;
	}
	return func (props *gin.Context)  {
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