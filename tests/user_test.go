package tests

import (
	"Picture-communicator/handlers"
	// "github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	// "path/filepath"
	// "runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// func init() {
// 	_, file, _, _ := runtime.Caller(1)
// 	_, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
// }

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
