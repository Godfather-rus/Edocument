package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) CreateEdoc(ctx context.Context, m bson.M) error {
	if m == nil {
		return nil
	}

	_, err := r.edocs.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}
