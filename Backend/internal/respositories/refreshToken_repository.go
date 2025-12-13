package respositories

import (
	"context"
	"fmt"

	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/database"
)

type RefreshTokenRepository struct {
	q *generated.Queries
}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{
		q: generated.New(database.DB),
	}
}

func (rtkn *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, args generated.CreateRefreshTokenParams) (generated.UserRefreshToken, error) {
	fmt.Printf("%+v\n", args)
	return rtkn.q.CreateRefreshToken(ctx, args)
}

func (rtkn *RefreshTokenRepository) RevokeRefreshToken(ctx context.Context, token string) (generated.UserRefreshToken, error) {
	return rtkn.q.RevokeRefreshToken(ctx, token)
}

func (rtkn *RefreshTokenRepository) GetUserFromRefreshToken(ctx context.Context, token string) (generated.User, error) {
	return rtkn.q.GetUserFromRefreshToken(ctx, token)
}
