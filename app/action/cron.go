package action

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zycron/app/job"
)

type Cron struct {}

//stopCron 重启调度器
func (cr *Cron) restartCron(c *gin.Context)  {
	job.RestartScheduler()

	c.JSON(http.StatusOK,gin.H{
		"isOK": true,
	})
}
//stopCron stopCron
func (cr *Cron) stopCron(c *gin.Context)  {
	job.StopScheduler()

	c.JSON(http.StatusOK,gin.H{
		"isOK": true,
	})
}


