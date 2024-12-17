package router

import (
	"github.com/Alfazal007/apiserver/controllers"
	"github.com/go-chi/chi/v5"
)

func UserRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/sign-up", apiCfg.Signup)
	r.Post("/sign-in", apiCfg.Signin)
	return r
}
