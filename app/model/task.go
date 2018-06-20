package model

import (
	"zycron/app/db"
)

//查询的字段
var fields ="id,exec_type,command,req_url,cron_spec,execute_times,exec_status"

type Task struct{
	Id int
	ExecType int  //执行类型 1 req_url 2 command
	Command string //当执行类型为登陆到服务器时 执行的命令
	ReqUrl string  //执行类型为请求url时 使用
	CronSpec string  //crontab 格式描述
	PrevTime string  //上次执行的时间
	ExecuteTimes int  //执行次数
	Status int //状态
	ExecStatus int //执行状态
}
//GetTaskById 获取任务
func GetTaskById(id string) (*Task, error) {

	var t Task
	db := db.GetMysqlInstance()

	queryStr := "SELECT "+fields+" FROM cms_task WHERE id=?"
	err := db.QueryRow(queryStr, id).Scan(&t.Id,
										&t.ExecType,
										&t.Command,
										&t.ReqUrl,
										&t.CronSpec,
										&t.ExecuteTimes,
										&t.ExecStatus,
										)
	if err!=nil {
		return nil, err
	}

	return &t, nil
}

//GetAllTask  获取所有的任务
func GetAllTask() ([]*Task, error)  {
	var taskList []*Task

	db := db.GetMysqlInstance()
	queryStr := "SELECT "+fields+" FROM cms_task WHERE is_del=1 and status=1"
	rows,err := db.Query(queryStr)
	if err!=nil {
		return nil,err
	}
	for rows.Next() {
		var t Task
		if err := rows.Scan(
			&t.Id,
			&t.ExecType,
			&t.Command,
			&t.ReqUrl,
			&t.CronSpec,
			&t.ExecuteTimes,
			&t.ExecStatus);err !=nil {
			return nil, err
		}
		taskList = append(taskList,&t)
	}
	return taskList, nil
}

func (t *Task)Update() (int64,error) {

	db := db.GetMysqlInstance()
	sql := "update cms_task set prev_time=?,execute_times=? where id=?"
	stmt,err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(t.PrevTime,t.ExecuteTimes,t.Id)
	if err != nil {
		return 0 ,err
	}
	return res.RowsAffected()    //返回影响的条数,注意有两个返回值

}

//UpdateStatus 更新状态
func (t *Task)UpdateStatus(field string) (int64,error) {

	db := db.GetMysqlInstance()
	sql := "update cms_task set " + field +"=? where id=?"
	stmt,err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(t.Status,t.Id)
	if err != nil {
		return 0 ,err
	}
	return res.RowsAffected()    //返回影响的条数,注意有两个返回值

}

