package job

import (

	"fmt"
	"html/template"
	"time"
	"log"
	"net/url"
	"strconv"

	"zycron/app/model"
	"zycron/app/util/httpclient"
)

var mailTpl *template.Template

func init() {

}

type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	task       *model.Task                                      // 任务对象
	runFunc    func(int64) (string,error) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	reqUrl     string                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(task *model.Task) (*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := makeHttpJob(task.Id, task.ReqUrl)
	job.task = task
	job.reqUrl = task.ReqUrl
	//job.Concurrent = task.Concurrent == 1
	return job, nil
}

//makeHttpJob 构造http触发式任务
func makeHttpJob(id int, reqUrl string) *Job {
	job := &Job{
		id:   id,
		reqUrl: reqUrl,
	}
	job.runFunc = func(logid int64) (string,error) {
		valus,err :=url.Parse(reqUrl)
		if err!=nil {
			log.Println(err)
		}

		req := reqUrl + "?task_id=" + strconv.Itoa(id)
		if len(valus.Query())>0 {
			req = reqUrl + "&task_id=" + strconv.Itoa(id)
		}
		req += "&log_id="+strconv.FormatInt(logid,10)

		fmt.Println(req)
		resp,err := httpclient.Get(req,0)
		return resp.Body,err
	}
	return job
}

//Run  实现定时执行接口
func (j *Job) Run() {
	task, _ := model.GetTaskById(strconv.Itoa(j.id))

	if (!j.Concurrent && j.status > 0) || task.ExecStatus==1 {
		log.Printf(fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。%s", j.id,
			time.Now().Format("2006-01-02 15:04:05")))
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()


	j.status++
	defer func() {
		j.status--
	}()

	// 产生日志id
	log := new(model.TaskLog)

	log.TaskId = j.id
	//1 已响应 具体执行结果由定时任务执行项目写入
	logId,_ := log.AddTaskLog()

	//执行触发
	body, err := j.runFunc(logId)

	fmt.Println(body,err)

	//触发后	

	log.Output =  body
	log.Error = ""
	if err != nil {
		log.Error = err.Error()
		//触发失败
		log.Status = 1
	}


	log.UpdateTaskLogById(logId)

	// 更新上次执行时间
	j.task.PrevTime = time.Now().Format("2006-01-02 15:04:05")
	j.task.ExecuteTimes++
	//j.task.ExecStatus = 1 //正在执行
	j.task.Update()

}


