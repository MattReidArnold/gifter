package mongo_test

import (
	"context"
	"testing"

	"github.com/mattreidarnold/gifter/frameworks/persistence/mongo"
	driver "go.mongodb.org/mongo-driver/mongo"
)

type tearDown func()

const dbPrefix = "groups_test_"

func setUp(t *testing.T) (client *driver.Client, db string, td tearDown) {
	t.Helper()
	db = dbPrefix + NewRandomID()
	logger := &stubLogger{}
	client, disconnect, err := mongo.NewClient(logger, mongo.Connection{
		Database: "admin",
		Host:     "localhost",
		Password: "SuperSecret789",
		Port:     "27017",
		Username: "root",
	})
	if err != nil {
		t.Fatalf("failed to create mongo client: %e", err)
	}

	err = client.Database(db).Drop(context.Background())
	if err != nil {
		t.Fatalf("failed to set up %s db: %e", db, err)
	}
	td = func() {
		err = client.Database(db).Drop(context.Background())
		if err != nil {
			t.Fatalf("failed to tear down %s db: %e", db, err)
		}
		disconnect()
	}
	return
}
