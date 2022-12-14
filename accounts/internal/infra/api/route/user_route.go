package route

import (
	"net/http"

	"github.com/julianojj/microservices/accounts/internal/infra/api/controller"
)

type UserRoute struct {
	app            *http.ServeMux
	userController *controller.CreateUserController
}

func NewUserRoute(
	app *http.ServeMux,
	userController *controller.CreateUserController,
) *UserRoute {
	return &UserRoute{
		app,
		userController,
	}
}

func (u *UserRoute) Init() {
	u.app.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" {
			u.userController.Handle(w, r)
			return
		}
	})
}
