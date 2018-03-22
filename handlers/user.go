package handlers

import (
	"Picture-communicator/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct {
	PathFunc map[string]string //path:funcname
}

func (u *User) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user *models.User
	if users, err := user.GetAll(); err != nil {
		w.WriteHeader(404)
	} else {
		resp, _ := json.Marshal(users)
		w.Write(resp)
		w.WriteHeader(200)
	}
}
func (u *User) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprint(w, "Get!\n")
}
func (u *User) Put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Put!\n")
}
func (u *User) Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Post!\n")
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
