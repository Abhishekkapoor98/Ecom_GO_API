package products

import (
	"context"

	repo "github.com/Abhishekkapoor98/Ecom_GO_API/internal/adapter/postgresql/sqlc"
)

type Service interface {
	// Define service methods here, e.g., ListProducts, GetProductByID, etc.

	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	// respository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	// Implement the logic to list products here.
	return s.repo.ListProducts(ctx)

}
