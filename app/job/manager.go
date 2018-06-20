//定时任务管理器
package job

import (
	"sync"
	"zycron/app/cron"
	"log"

)

var (
	mainCron *cron.Cron
	lock     sync.Mutex
)
//init  初始化
func init() {
	mainCron = cron.New()
	mainCron.Start()
}

//StartAllJobByDb 开始所有状态status==1 的任务
func StartAllJobByDb()  {
	BootAllJobByDb()
}
//AddJob 添加任务到调度器
func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.id) != nil {
		return false
	}

	err := mainCron.AddJob(spec, job)

	if err != nil {
		log.Println("AddJob: ", err.Error())
		return false
	}
	return true
}

//RemoveJob 已id来移除job
func RemoveJob(id int) {
	mainCron.RemoveJob(func(e *cron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return true
			}
		}
		return false
	})
}

//GetEntryById  获取cron执行体
func GetEntryById(id int) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}

//RestartScheduler  重启调度器
func RestartScheduler()  {
	mainCron.Stop()
	mainCron.Start()
}

//RestartScheduler  stop调度器
func StopScheduler()  {
	mainCron.Stop()
}

//IsRunScheduler 调度器是否运行
func IsRunScheduler() bool {
	return mainCron.IsRunIng()
}
