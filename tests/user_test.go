package tests

import (
	"Picture-communicator/handlers"
	// "github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
用户注册
用户登录
用户登出
用户获取个人信息
用户更新个人信息
用户删除个人信息
*/

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/user/2", nil)
	w := httptest.NewRecorder()
	var user *handlers.User
	// var params = httprouter.Params{
	// 	httprouter.Param{Key: "id", Value: "1"},
	// }
	user.Get(w, r, nil)
	Convey("User: get info\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
