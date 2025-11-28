package respositories

import (
	"context"

	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/database"
)

type UserRespository struct {
	q *generated.Queries
}

func NewUserRepository() *UserRespository {
	return &UserRespository{
		q: generated.New(database.DB),
	}
}

func (r *UserRespository) CreateUser(ctx context.Context, args generated.CreateUserParams) (generated.User, error) {
	return r.q.CreateUser(ctx, args)
}

func (r *UserRespository) GetUserByEmail(ctx context.Context, email string) (generated.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}
