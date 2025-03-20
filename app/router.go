package app

import (
	"pjt1/app/controller"
	"pjt1/app/repo"
	"pjt1/app/service"

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
		r.Post("/create", urController.SaveUserDetails)
		r.Post("/login", urController.LoginUser)
	})

	return r
}
