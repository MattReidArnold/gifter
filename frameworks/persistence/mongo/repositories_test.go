package mongo_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/frameworks/persistence/mongo"
)

type stubLogger struct{}

func (l *stubLogger) Info(args ...interface{})             {}
func (l *stubLogger) Error(err error, args ...interface{}) {}

func Test_GroupRepository_Get_WhenGroupDoesNotExist(t *testing.T) {
	logger := &stubLogger{}
	// set up clean db
	client, disconnect, err := mongo.NewClient(logger, mongo.Connection{
		Database: "admin",
		Host:     "localhost",
		Password: "SuperSecret789",
		Port:     "27017",
		Username: "root",
	})
	if err != nil {
		t.Fatal("failed to create mongo client:", err)
	}
	defer disconnect()
	err = client.Database("groups_test").Drop(context.Background())
	if err != nil {
		t.Fatal("failed to drop groups_test db:", err)
	}
	// add group to Mongo with client
	groupID := "8475439584"

	// user repository to retrieve Group
	repo := mongo.NewGroupRepository(client, "groups_test")
	_, err = repo.Get(context.Background(), groupID)
	// assert nil error
	AssertEqual(t, err, app.ErrGroupNotFound)
}
func Test_GroupRepository_Get_WhenGroupExists(t *testing.T) {
	logger := &stubLogger{}
	// set up clean db
	client, disconnect, err := mongo.NewClient(logger, mongo.Connection{
		Database: "admin",
		Host:     "localhost",
		Password: "SuperSecret789",
		Port:     "27017",
		Username: "root",
	})
	if err != nil {
		t.Fatal("failed to create mongo client:", err)
	}
	defer disconnect()
	err = client.Database("groups_test").Drop(context.Background())
	if err != nil {
		t.Fatal("failed to drop groups_test db:", err)
	}
	// add group to Mongo with client
	groupID := "8475439584"
	groupDoc := mongo.Group{
		ID:     groupID,
		Name:   "test-group-name",
		Budget: 249.99,
		Gifters: []mongo.Gifter{
			{ID: "6789", Name: "Jim"},
			{ID: "6790", Name: "Joe"},
			{ID: "6791", Name: "Jill"},
		},
	}
	_, err = client.Database("groups_test").Collection("groups").InsertOne(context.Background(), groupDoc)
	if err != nil {
		t.Fatal("failed inserting group doc:", err)
	}
	// user repository to retrieve Group
	repo := mongo.NewGroupRepository(client, "groups_test")
	got, err := repo.Get(context.Background(), groupID)
	// assert nil error
	if err != nil {
		t.Fatal("failed getting group:", err)
	}

	// assert correct group
	AssertEqual(t, got.ID(), groupID)
	AssertEqual(t, got.Name(), groupDoc.Name)
	AssertEqual(t, got.Budget(), groupDoc.Budget)
	AssertEqual(t, len(got.Gifters()), len(groupDoc.Gifters))
	for i, gifter := range got.Gifters() {
		doc := groupDoc.Gifters[i]
		AssertEqual(t, gifter.ID(), doc.ID)
		AssertEqual(t, gifter.Name(), doc.Name)
	}
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	gotType := reflect.TypeOf(got)
	wantType := reflect.TypeOf(want)
	if gotType != wantType {
		t.Errorf("got type: %v, want type: %v", gotType, wantType)
		return
	}
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
