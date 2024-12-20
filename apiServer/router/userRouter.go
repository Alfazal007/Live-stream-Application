package router

import (
	"net/http"

	"github.com/Alfazal007/apiserver/controllers"
	"github.com/go-chi/chi/v5"
)

func UserRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/sign-up", apiCfg.Signup)
	r.Post("/sign-in", apiCfg.Signin)
	r.Get("/current-user", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.GetCurrentUser)).ServeHTTP)
	r.Post("/user-validate", apiCfg.IsValidUser)
	return r
}
