package db
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"zycron/app/util"
)

//getInstance 获取db对象
func GetMysqlInstance() *sql.DB  {
	host := util.GetConfig("mysql","host")
	port := util.GetConfig("mysql","port")
	user := util.GetConfig("mysql","user")
	password := util.GetConfig("mysql","password")
	name := util.GetConfig("mysql","name")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
