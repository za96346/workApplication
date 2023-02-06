package tasktimer
import (
	"fmt"
	"github.com/robfig/cron/v3"
)
func AddTask(task ...func()) {
	c := cron.New()
	for _, v := range task {
		_, err := c.AddFunc("@every 10s", v)
		if err != nil {
			fmt.Println("AddFunc error :", err)
			return
		 }
	}
	c.Start()
	defer c.Stop()
	select {}
}