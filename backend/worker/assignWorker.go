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
		// routeFunc = service.FetchBanchRule
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