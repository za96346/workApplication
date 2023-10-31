package tasktimer
import (
	"fmt"
	"github.com/robfig/cron/v3"
)
func AddDailyTask(task ...func()) {
	c := cron.New()
	for _, v := range task {
		_, err := c.AddFunc("@daily", v)
		if err != nil {
			fmt.Println("AddFunc error :", err)
			return
		 }
	}
	c.Start()
	defer c.Stop()
	select {}
}