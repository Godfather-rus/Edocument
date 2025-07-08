package app

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (a *App) initDBConn() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	a.closers = append(a.closers, func() error {
		err := dbClient.Disconnect(context.Background())
		if err != nil {
			return err
		}
		return nil
	})

	a.dbClient = dbClient

	return nil
}
