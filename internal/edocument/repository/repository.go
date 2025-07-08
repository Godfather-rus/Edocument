package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	edocs *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		edocs: client.Database("core").Collection("edocuments"),
	}
}
