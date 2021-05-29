package mongo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mattreidarnold/gifter/frameworks/persistence/mongo"
	"github.com/mattreidarnold/gifter/test"
	"github.com/mattreidarnold/gifter/test/stub"
	driver "go.mongodb.org/mongo-driver/mongo"
)

type tearDown func()

const dbPrefix = "groups_test_"

func setUp(t *testing.T) (client *driver.Client, db string, td tearDown) {
	t.Helper()
	db = dbPrefix + test.NewRandomID()
	logger := stub.NewStubLogger()
	client, disconnect, err := mongo.NewClient(logger, mongo.Connection{
		Database: "admin",
		Host:     "localhost",
		Password: "SuperSecret789",
		Port:     "27017",
		Username: "root",
	})
	test.AssertNil(t, err, "failed to create mongo client")

	err = client.Database(db).Drop(context.Background())
	test.AssertNil(t, err, fmt.Sprintf("failed to set up %s db", db))

	td = func() {
		err = client.Database(db).Drop(context.Background())
		test.AssertNil(t, err, fmt.Sprintf("failed to tear down %s db", db))
		disconnect()
	}
	return
}
