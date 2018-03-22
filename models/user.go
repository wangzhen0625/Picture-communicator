package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func (u *User) GetAll() (users []User, err error) {
	err = GormDb.Find(&users).Error
	return
}
func (u *User) Get(id int) (user User, err error) {
	err = GormDb.First(&user, id).Error
	return
}

func (u *User) Post() (user User, err error) {
	if err = GormDb.Where(User{Name: u.Name}).First(&user).Error; err == nil {
		return
	} else {
		if err = GormDb.Create(&user).Error; err != nil {
			return
		}
	}
	return

	//byte数组直接转成string，优化内存
	// str := (*string)(unsafe.Pointer(&respBytes))
	// str := bytes.NewBuffer(respBytes).Bytes() //.String()
}
