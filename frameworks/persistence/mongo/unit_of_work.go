package mongo

import (
	"context"

	"github.com/mattreidarnold/gifter/app"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUOW struct {
	*groupRepository
	client *mongo.Client
}

func MongoUnitOfWork(client *mongo.Client, db string) app.UseUnitOfWork {
	return func(c context.Context, f func(context.Context, app.UnitOfWork) error) error {
		uow := &mongoUOW{
			client:          client,
			groupRepository: NewGroupRepository(client, db),
		}
		return client.UseSession(c, func(sc mongo.SessionContext) error {
			return f(sc, uow)
		})
	}
}

func (uow *mongoUOW) Groups() app.GroupRepository {
	return uow.groupRepository
}
