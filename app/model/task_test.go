package model

import (
	"testing"

	"fmt"
)

func TestGetTaskById(t *testing.T) {

	res,err := GetTaskById("id")
	fmt.Println(res.Id,err)

}

func TestGetAllTask(t *testing.T) {

	//res,err := GetAllTask()
	//fmt.Println(res[0].Id,err)

}
