package routers

import (
	. "Picture-communicator/handlers"
)

func AllRoutes() Routes {
	var user = &User{}
	return Routes{
		//name method patch  handler
		//  /user
		Route{"Index", "GET", "/user", user.Index},
		Route{"Get", "GET", "/user/:id", user.Get},
		Route{"Post", "POST", "/user", user.Post},
		Route{"Put", "PUT", "/user/:id", user.Put},
		Route{"Patch", "PATCH", "/user/:id", user.Patch},
		Route{"Delete", "DELETE", "/user/:id", user.Delete},
		Route{"Login", "POST", "/user/login", user.Login},
		Route{"Logout", "DELETE", "/user/:id/logout", user.Logout},
		Route{"CreateTable", "POST", "/user/createtb", user.CreateTable},
	}

}
