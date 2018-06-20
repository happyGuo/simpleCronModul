package action

import (
	"github.com/gin-gonic/gin"
	"errors"
)

var ErrAbort = errors.New("User stop run")

//HttpServerRun runHttpServer
func HttpServerRun() {
	r := gin.Default()
	r.GET("/task/start", new(Task).start)
	r.GET("/task/stop", new(Task).stop)
	//r.GET("/task/setStatus", new(Task).setStatus)
	r.GET("/cron/restartCron", new(Cron).restartCron)
	r.GET("/cron/stop", new(Cron).stopCron)

	r.Run(":9999")
}

//StopRun 停止执行
func StopRun() {
	panic(ErrAbort)
}

