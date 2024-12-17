package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Alfazal007/apiserver/controllers"
	"github.com/Alfazal007/apiserver/internal/database"
	"github.com/Alfazal007/apiserver/router"
	"github.com/Alfazal007/apiserver/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	// load env variables
	envVariables := utils.LoadEnvFiles()
	conn, err := sql.Open("postgres", envVariables.DatabaseUrl)
	if err != nil {
		log.Fatal("Issue connecting to the database", err)
	}

	apiCfg := controllers.ApiConf{DB: database.New(conn)}

	// add middlewares
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)

	r.Mount("/api/v1/user", router.UserRouter(&apiCfg))

	log.Println("Starting the server at port", envVariables.PortNumber)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", envVariables.PortNumber), r)
	if err != nil {
		log.Fatal("There was an error starting the server", err)
	}
}
