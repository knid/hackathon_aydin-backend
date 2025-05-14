package main

import (
	"fmt"
	"net/http"

	"github.com/knid/timezen/internal/controllers"
	"github.com/knid/timezen/internal/models"
	"github.com/knid/timezen/internal/views"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	dsn := "host=localhost user=timezen password=timezen dbname=timezen port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Err", err)
		return
	}

	db.AutoMigrate(&models.User{})	
	db.AutoMigrate(&models.Token{})
	// db.AutoMigrate(&models.Chat{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Project{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
	    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	    AllowedOrigins:   []string{"https://*", "http://*"},
	    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	    ExposedHeaders:   []string{"Link"},
	    AllowCredentials: false,
	    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	vw := views.Views{
		Controller: &controllers.Controller{
			DB: db,
		},
	}

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", vw.GetAPIInfo)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", vw.GetUserInfo)
			r.Post("/login", vw.LoginUser)
			r.Post("/", vw.CreateUser)
			// r.Put("/", vw.UpdateUser)
			// r.Delete("/", vw.DeleteUser)
		})
		r.Route("/projects", func(r chi.Router) {
			r.Get("/", vw.GetProjects)
			r.Post("/", vw.CreateProject)
			r.Route("/{projectId}", func(r chi.Router){
				r.Get("/", vw.GetProject)
		// 		r.Put("/", vw.UpdateProject)
		// 		r.Delete("/", vw.DeleteProject)
		//
		// 		r.Route("/users", func(r chi.Router) {
		// 			r.Get("/", vw.GetProjectUsers)
		// 			r.Post("/invite", vw.InviteUserToProject)
		// 			r.Delete("/", vw.DeleteUserFromProject)
		// 		})
				r.Route("/tasks", func(r chi.Router) {
					r.Get("/", vw.GetProjectTasks)
		// 			r.Post("/", vw.CreateProjectTask)
		// 			r.Route("/{taskId}", func(r chi.Router) {
		// 				r.Get("/", vw.GetTask)
		// 				r.Put("/", vw.UpdateTask)
		// 				r.Delete("/", vw.DeleteTask)
					})
		// 		})
			})
		})


	})

	http.ListenAndServe(":3993", r)
}
