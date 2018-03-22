package app

import (
	"Picture-communicator/handlers"
	"Picture-communicator/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var err error

func Run() {
	handlers.GormDb, err = gorm.Open("mysql", "root:P@ssw0rd1!@tcp(192.168.74.23:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer handlers.GormDb.Close()
	router := routers.NewRouter(routers.AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
