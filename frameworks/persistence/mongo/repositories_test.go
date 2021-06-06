package mongo_test

import (
	"context"
	"testing"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/frameworks/persistence/mongo"
	"github.com/mattreidarnold/gifter/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	driver "go.mongodb.org/mongo-driver/mongo"
)

const groupsCollection = "groups"

func Test_GroupRepository_Add_WhenGroupDoesNotExist(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	groupID := test.NewRandomID()

	group := domain.NewGroup(groupID, "test-group-name", 42, []domain.Gifter{domain.NewGifter("9876", "Bob")})

	repo := mongo.NewGroupRepository(client, db)
	err := repo.Add(context.Background(), group)

	assert.Nil(t, err, "failed saving group")

	groupDoc := mongo.Group{}
	err = client.Database(db).Collection(groupsCollection).FindOne(context.Background(), byIdFilter(groupID)).Decode(&groupDoc)
	assert.Nil(t, err, "failed reloading saved group")

	assertModelEqualsDoc(t, group, groupDoc)

	assertCountIs(t, client, db, groupID, 1)
}

func Test_GroupRepository_Add_WhenGroupAlreadyIDExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	want := app.ErrGroupIDAlreadyExists

	groupID := test.NewRandomID()
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
	_, err := client.Database(db).Collection(groupsCollection).InsertOne(context.Background(), groupDoc)
	assert.Nil(t, err, "failed inserting group doc")

	group := domain.NewGroup(groupID, "test-some-other-name", 22, []domain.Gifter{})

	repo := mongo.NewGroupRepository(client, db)
	err = repo.Add(context.Background(), group)

	assert.ErrorIs(t, err, want)
}
func Test_GroupRepository_Get_WhenGroupDoesNotExist(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	want := app.ErrGroupNotFound

	groupID := test.NewRandomID()

	repo := mongo.NewGroupRepository(client, db)
	_, err := repo.Get(context.Background(), groupID)

	assert.Equal(t, err, want)
}

func Test_GroupRepository_Get_WhenGroupExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	groupID := test.NewRandomID()
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
	_, err := client.Database(db).Collection(groupsCollection).InsertOne(context.Background(), groupDoc)
	assert.Nil(t, err, "failed inserting group doc")

	repo := mongo.NewGroupRepository(client, db)
	got, err := repo.Get(context.Background(), groupID)
	assert.Nil(t, err, "failed getting Group")

	assertModelEqualsDoc(t, got, groupDoc)
}

func Test_GroupRepository_Save_WhenGroupDoesNotExist(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	want := app.ErrGroupNotFound

	groupID := test.NewRandomID()

	repo := mongo.NewGroupRepository(client, db)
	group := domain.NewGroup(groupID, "test-group-name", 42, []domain.Gifter{domain.NewGifter("9876", "Bob"), domain.NewGifter("3452", "Daryl")})

	err := repo.Save(context.Background(), group)
	assert.Equal(t, err, want)

	assertCountIs(t, client, db, groupID, 0)
}

func Test_GroupRepository_Save_WhenGroupExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	groupID := test.NewRandomID()
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
	_, err := client.Database(db).Collection(groupsCollection).InsertOne(context.Background(), groupDoc)
	assert.Nil(t, err, "failed inserting group doc")

	repo := mongo.NewGroupRepository(client, db)
	group := domain.NewGroup(groupID, "test-updated-group-name", 42, []domain.Gifter{domain.NewGifter("9876", "Bob")})

	err = repo.Save(context.Background(), group)
	assert.Nil(t, err, "failed saving Group")

	updatedGroupDoc := mongo.Group{}
	err = client.Database(db).Collection(groupsCollection).FindOne(context.Background(), byIdFilter(groupID)).Decode(&updatedGroupDoc)
	assert.Nil(t, err, "failed reloading group doc")

	assertModelEqualsDoc(t, group, updatedGroupDoc)

	assertCountIs(t, client, db, groupID, 1)
}

func byIdFilter(id string) bson.D {
	return bson.D{{"identifier", id}}
}

func assertCountIs(t *testing.T, client *driver.Client, db string, groupID string, want int64) {
	t.Helper()
	count, err := client.Database(db).Collection(groupsCollection).CountDocuments(context.Background(), byIdFilter(groupID))
	assert.Nil(t, err, "failed counting docs")
	assert.Equal(t, count, want, "doc count")
}

func assertModelEqualsDoc(t *testing.T, m domain.Group, d mongo.Group) {
	t.Helper()
	assert.Equal(t, m.ID(), d.ID, "group.ID()")
	assert.Equal(t, m.Name(), d.Name, "group.Name()")
	assert.Equal(t, m.Budget(), d.Budget, "group.Budget()")
	assert.Equal(t, len(m.Gifters()), len(d.Gifters), "len(group.Gifters())")
	for i, gifter := range m.Gifters() {
		doc := d.Gifters[i]
		assert.Equal(t, gifter.ID(), doc.ID, "gifter.ID()")
		assert.Equal(t, gifter.Name(), doc.Name, "gifter.Name()")
	}
}
