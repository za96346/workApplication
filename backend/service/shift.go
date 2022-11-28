package service
import (
	"sync"

	"github.com/gin-gonic/gin"
)

// work time
func FetchWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func InsertWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func DeleteWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func UpdateWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}

// paid Vocation
func FetchPaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func InsertPaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func DeletePaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func UpdatePaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}