package worker
import (
	"backend/service"
	"github.com/gin-gonic/gin"
	"sync"
)


func AssignWorker(routerMethod int) func(props *gin.Context) {
	routeFunc := service.FindSingleUser
	switch routerMethod {
	case 0:
		//get method of fetch single user
		routeFunc = service.FindSingleUser
		break;
	case 1:
		//post method of create user
		routeFunc = service.CreateUser
		break;
	case 2:
		//put method of update user
		routeFunc = service.UpdateUser
		break;
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