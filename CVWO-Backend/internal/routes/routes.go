package routes

import (
	"github.com/XKMai/CVWO-React/CVWO-Backend/internal/handlers/posts"
	"github.com/XKMai/CVWO-React/CVWO-Backend/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() func(chi.Router) {
	r := chi.NewRouter()
	userhandler := users.UserHandler{}
	r.Get("/", userhandler.ListUsers)
	r.Post("/", userhandler.CreateUser)
	r.Get("/{id}", userhandler.GetUser)
	r.Put("/{id}", userhandler.UpdateUser)
	r.Delete("/{id}", userhandler.DeleteUser)
	return func(r chi.Router) {
		r.Mount("/", r)
	}
}

func PostRoutes() func(chi.Router) {
	r := chi.NewRouter()
	userhandler := posts.PostHandler{}
	r.Get("/", userhandler.ListPosts)
	r.Post("/", userhandler.CreatePost)
	r.Get("/{id}", userhandler.GetPost)
	r.Put("/{id}", userhandler.UpdatePost)
	r.Delete("/{id}", userhandler.DeletePost)
	return func(r chi.Router) {
		r.Mount("/", r)
	}
}