package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/mattreidarnold/gifter/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Disconnect func() error

type Connection struct {
	Database string
	Host     string
	Password string
	Port     string
	Username string
}

func (c Connection) ToURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
}

func NewClient(logger app.Logger, conn Connection) (*mongo.Client, Disconnect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Info("connecting to mongo")
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://groupsUser:Password123@localhost:27017/groups")))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn.ToURI()))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	logger.Info("testing mongo connetion")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	return client,
		func() error {
			logger.Info("disconnecting from mongo")
			return client.Disconnect(context.Background())
		}, nil
}
