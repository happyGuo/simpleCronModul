package main

import (
	"zycron/app/action"
	"zycron/app/job"
)

func main()  {
	//启动所有任务
	job.StartAllJobByDb()
	//启动HTTP服务
	action.HttpServerRun()
}
