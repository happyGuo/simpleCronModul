package job

import (
	"zycron/app/model"
)
//BootAllJobByDb 启动数据库内符合的所有任务
func BootAllJobByDb() {
	tasks,_ := model.GetAllTask()

	for _, task := range tasks {
		job, err := NewJobFromTask(task)
		if err != nil {
			continue
		}

		if isOK :=AddJob(task.CronSpec, job);!isOK{
			log := new(model.TaskLog)
			log.TaskId = task.Id
			log.Error = "启动失败，请检查表达式"
			log.AddTaskLog()
		}
	}
}