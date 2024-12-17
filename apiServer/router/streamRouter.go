package router

import (
	"net/http"

	"github.com/Alfazal007/apiserver/controllers"
	"github.com/go-chi/chi/v5"
)

func StreamRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create-stream", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CreateStream)).ServeHTTP)
	r.Put("/start-stream", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.StartStream)).ServeHTTP)
	r.Put("/end-stream", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.EndStream)).ServeHTTP)
	return r
}
