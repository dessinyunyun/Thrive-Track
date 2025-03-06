package repository

import (
	"context"
	"go-gin/database/ent"
	entActivationToken "go-gin/database/ent/activation_token"

	"github.com/google/uuid"
	googleUUID "github.com/google/uuid"
)

type AuthRepository struct {
	db *ent.Client
}

func NewAuthRepository(db *ent.Client) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) GetDetailAT(ctx context.Context, userID googleUUID.UUID) (*ent.Activation_token, error) {
	atQuery := r.db.Activation_token.Query().
		Where(
			entActivationToken.UserID(userID),
		)

	if count, _ := atQuery.Count(ctx); count == 0 {
		return nil, nil
	}

	exec, _ := atQuery.First(ctx)

	return exec, nil
}

func (r *AuthRepository) CreateAT(ctx context.Context, token string, userID string) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	userIdParse, err := uuid.Parse(userID)
	_, err = tx.Activation_token.Create().
		SetToken(token).
		SetUserID(userIdParse).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) UsedAT(ctx context.Context, userID googleUUID.UUID) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	err = tx.Activation_token.Update().
		Where(entActivationToken.UserIDEQ(userID)).
		SetIsused(true).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
