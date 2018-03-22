package handlers

import (
	// "Picture-communicator/models"
	// "bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

var GormDb *gorm.DB

type Test struct {
	gorm.Model
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
type User struct {
	PathFunc map[string]string //path:funcname
}

func (u *User) CreateTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !GormDb.HasTable(&Test{}) {
		if err := GormDb.CreateTable(&Test{}).Error; err != nil {
			fmt.Println("create user err:", err)
		}
	}
	fmt.Fprint(w, "CreateTable!\n")
}
func (u *User) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var users []Test
	if err := GormDb.Find(&users).Error; err != nil {
		fmt.Println("getall err:", err)
	}
	fmt.Println("get all:", users)
	fmt.Fprint(w, "Index!\n")
}
func (u *User) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	var user Test
	if err := GormDb.First(&user, id).Error; err != nil {
		fmt.Println("get err:", err)
	}
	fmt.Println("get :", user)
	fmt.Fprint(w, "Get!\n")
}
func (u *User) Put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Put!\n")
}
func (u *User) Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Post!\n")
	fmt.Println("p:", p)
	respBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var tests Test

	fmt.Println("respBytes:", string(respBytes))
	json.Unmarshal(respBytes, &tests)

	if err := GormDb.Where(Test{Name: tests.Name}).First(&tests).Error; err == nil {
		fmt.Println("create user has exist:", err)
	} else {
		if err := GormDb.Create(&tests).Error; err != nil {
			fmt.Println("create user:", err)
		}
	}

	//byte数组直接转成string，优化内存
	// str := (*string)(unsafe.Pointer(&respBytes))
	// str := bytes.NewBuffer(respBytes).Bytes() //.String()
}
func (u *User) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Delete!\n")
}
func (u *User) Patch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Delete!\n")
}
func (u *User) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Login!\n")
}
func (u *User) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Logout!\n")
}
