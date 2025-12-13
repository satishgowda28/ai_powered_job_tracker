package respositories

import (
	"context"

	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/database"
)

type UserRepository struct {
	q *generated.Queries
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		q: generated.New(database.DB),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, args generated.CreateUserParams) (generated.User, error) {
	return r.q.CreateUser(ctx, args)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (generated.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}
