package mongo

import (
	"context"
	"fmt"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type groupRepository struct {
	client *mongo.Client
}

type Gifter struct {
	ID   string `bson:"identifier"`
	Name string `bson:"name"`
}

type Group struct {
	ID      string   `bson:"identifier"`
	Name    string   `bson:"name"`
	Budget  float64  `bson:"budget"`
	Gifters []Gifter `bson:"gifters"`
}

func NewGroupRepository(c *mongo.Client) app.GroupRepository {
	return &groupRepository{
		client: c,
	}
}

func (gr *groupRepository) Get(ctx context.Context, id string) (domain.Group, error) {
	doc := Group{}
	err := gr.collection().FindOne(ctx, bson.D{{"identifier", id}}).Decode(&doc)
	if err != nil {
		return nil, err
	}
	return doc.ToModel(), nil
}

func (gr *groupRepository) Add(ctx context.Context, g domain.Group) error {
	return nil
}

func (gr *groupRepository) Save(ctx context.Context, g domain.Group) error {
	doc := &Group{}
	doc.FromModel(g)
	fmt.Printf("saving %+v", doc)
	err := gr.collection().FindOneAndReplace(ctx, bson.D{{"identifier", g.ID()}}, doc).Decode(&Group{})
	return err
}

func (gr *groupRepository) collection() *mongo.Collection {
	return gr.client.Database("groups").Collection("groups")
}

func (doc Group) ToModel() domain.Group {
	gifters := []domain.Gifter{}
	for _, gifter := range doc.Gifters {
		gifters = append(gifters, domain.NewGifter(gifter.ID, gifter.Name))
	}
	return domain.NewGroup(doc.ID, doc.Name, doc.Budget, gifters)
}

func (doc *Group) FromModel(model domain.Group) {
	doc.ID = model.ID()
	doc.Name = model.Name()
	doc.Budget = model.Budget()
	doc.Gifters = []Gifter{}
	for _, gifter := range model.Gifters() {
		doc.Gifters = append(doc.Gifters, Gifter{
			ID:   gifter.ID(),
			Name: gifter.Name(),
		})
	}
}
