package app

import (
	"github.com/jmoiron/sqlx"

	"github.com/hoachnt/go-todo-app/internal/handler"
	"github.com/hoachnt/go-todo-app/internal/repository"
	"github.com/hoachnt/go-todo-app/internal/service"
	"github.com/hoachnt/go-todo-app/pkg/auth"
	"github.com/hoachnt/go-todo-app/pkg/swagger"
)

// 1. Repositories: Interface to interact with the database.
// 2. Services: Business logic layer using the initialized repositories.
// 3. Auth: Authentication module using the services for user management.
// 4. Swagger: API documentation module.
// 5. Handlers: HTTP handlers that combine all services, auth, and swagger.
func InitializeModules(db *sqlx.DB) *handler.Handler {
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	auth := auth.NewUser(repos)
	swagger := swagger.NewSwagger()

	return handler.NewHandler(services, auth, swagger)
}
