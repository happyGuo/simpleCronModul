package action

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"

	"zycron/app/model"
	"zycron/app/job"

)

type Task struct {

}

//start 加入一个任务到定时任务
func (t *Task)start(c *gin.Context) {
	prevCheck(c)

	id := c.Query("id")
	task,err := model.GetTaskById(id)

	if err !=nil {
		c.JSON(http.StatusOK,gin.H{
			"isOK": false,
		})
		StopRun()
	}
	jobObj, err :=job.NewJobFromTask(task)

	if isOk :=job.AddJob(task.CronSpec, jobObj); !isOk {
		c.JSON(http.StatusOK,gin.H{
			"isOK": false,
			"msg":"启动失败，请检查表达式",
		})
		StopRun()

	}

	task.Status = 1
	task.UpdateStatus("status")

	c.JSON(http.StatusOK,gin.H{
		"isOK": true,
	})
}

// stop 从任务调度中取出 并且修改任务状态
func (t *Task)stop(c *gin.Context) {
	prevCheck(c)

	id := c.Query("id")
	task, err := model.GetTaskById(id)
	if err !=nil {
		c.JSON(http.StatusOK,gin.H{
			"isOK": false,
		})
	}
	idInt, _ := strconv.Atoi(id)
	job.RemoveJob(idInt)
	task.Status = 0
	task.UpdateStatus("status")

	c.JSON(http.StatusOK,gin.H{
		"isOK": true,
	})
}

//setStatus 设置任务执行状态
/*func (t *Task)setStatus(c *gin.Context)  {
	id := c.Query("id")
	status := c.Query("status")
	errStr := c.Query("err")
	out := c.Query("out")
	execTime := c.Query("process_time")

	fmt.Println(out,errStr,status,id,execTime)
	task, err := model.GetTaskById(id)
	if err !=nil {

	}
	//啊啊啊
	intId, _ := strconv.Atoi(id)
	intStatus, _ := strconv.Atoi(status)
	intexecTime, _ := strconv.Atoi(execTime)




	task.Status = intStatus
	task.UpdateStatus("exec_status")

	//fmt.Println(intId,intStatus, intexecTime)

	log := new(model.TaskLog)
	log.TaskId = intId
	log.Output = out
	//2 触发日志
	log.Status = 2
	log.ProcessTime = intexecTime
	log.Error = errStr
	log.AddTaskLog()


	c.JSON(http.StatusOK,gin.H{
		"isOK": true,
	})


}*/
//prevCheck 前置检查
func prevCheck(c *gin.Context)  {
	if (!job.IsRunScheduler()){
		c.JSON(http.StatusOK,gin.H{
			"isOK": false,
			"msg":"调度器已关闭",
		})
		StopRun()
	}
}
