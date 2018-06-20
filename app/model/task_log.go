package model

import (
     "zycron/app/db"
	"fmt"
)

type TaskLog struct {
	TaskId      int
	Output      string
	Error       string
	Status      int
	ProcessTime int
}

//AddTaskLog 加入日志
func (t *TaskLog)AddTaskLog() (int64, error)  {
	insertSql := "insert into cms_task_log (task_id,output,error,status,process_time) values (?,?,?,?,?)"
	db := db.GetMysqlInstance()
	stmt, _ :=db.Prepare(insertSql)
	defer stmt.Close()
	result, _ := stmt.Exec(t.TaskId, t.Output, t.Error, t.Status, t.ProcessTime)
	return result.LastInsertId()
}

//UpdateTaskLogById 更新任务状态
func (t *TaskLog)UpdateTaskLogById(id int64) (int64, error) {
	db := db.GetMysqlInstance()
	sql := "update cms_task_log set status=?,output=?,error=?  where id=?"
	stmt,err := db.Prepare(sql)
	fmt.Println(t.Status,t.Output,id,err)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(t.Status,t.Output,t.Error,id)
	if err != nil {
		return 0 ,err
	}
	return res.RowsAffected()    //返回影响的条数,注意有两个返回值
}