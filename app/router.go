package app

import (
	"pjt1/app/controller"
	"pjt1/app/repo"
	"pjt1/app/service"

	"pjt1/pkg/middleware"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func APIRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	// User part
	urRepo := repo.NewUserRepo(db)
	urService := service.NewUserService(urRepo)
	urController := controller.NewUserController(urService)

	// r.Route("/", func(r chi.Router) {
	// 	r.Get("/hello", api.ExampleHamdler)
	// })

	//user
	r.Route("/user", func(r chi.Router) {
		r.Post("/login", urController.LoginUser)
		r.Post("/create", urController.SaveUserDetails)

		r.With(middleware.JWTAuthMiddleware).Get("/hello", urController.ExampleHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware) // Applying JWT middleware

		//r.Get("/hello", urController.ExampleHandler)
		//r.Post("/create", urController.SaveUserDetails)
	})

	return r
}
