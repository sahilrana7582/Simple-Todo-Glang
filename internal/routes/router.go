package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/simple-todo-golang/internal/handlers"
)

func InitializeRouter(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", userHandler.CreateUser).Methods(http.MethodPost)                 // POST /users
	userRouter.HandleFunc("/{id}", userHandler.GetUserByID).Methods(http.MethodGet)            // GET /users/{id}
	userRouter.HandleFunc("/{id}", userHandler.UpdateUser).Methods(http.MethodPut)             // PUT /users/{id}
	userRouter.HandleFunc("/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)          // DELETE /users/{id}
	userRouter.HandleFunc("/username", userHandler.GetUserByUsername).Methods(http.MethodPost) // POST /users/username
	userRouter.HandleFunc("/email", userHandler.GetUserByEmail).Methods(http.MethodPost)       // POST /users/email

	return r
}
