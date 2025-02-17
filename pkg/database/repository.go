package database

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"schaub-dev.xyz/fabrice/fiber-template/pkg/middleware"
)

type Repository[T any] interface {
	Name() string
	Create(item T) error
	Get(id uuid.UUID) (T, error)
	GetAll() ([]T, error)
	Update(item T) error
	Delete(id uuid.UUID) error
}

func InjectRepository[T any](app *fiber.App, repo Repository[T]) {
	app.Use(middleware.InjectItem(repo.Name(), repo))
}
