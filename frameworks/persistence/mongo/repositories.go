package mongo

import (
	"context"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type groupRepository struct {
	client *mongo.Client
	db     string
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

func NewGroupRepository(c *mongo.Client, db string) app.GroupRepository {
	return &groupRepository{
		client: c,
		db:     db,
	}
}

func (gr *groupRepository) Add(ctx context.Context, g domain.Group) error {
	doc := &Group{}
	doc.FromModel(g)
	filter := bson.D{{"identifier", g.ID()}}
	err := gr.collection().FindOne(ctx, filter).Err()
	if err == nil {
		return app.ErrGroupIDAlreadyExists
	}
	if err != mongo.ErrNoDocuments {
		return err
	}
	_, err = gr.collection().InsertOne(ctx, doc)
	return err
}

func (gr *groupRepository) Get(ctx context.Context, id string) (domain.Group, error) {
	doc := Group{}
	err := gr.collection().FindOne(ctx, bson.D{{"identifier", id}}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return nil, app.ErrGroupNotFound
	}
	if err != nil {
		return nil, err
	}
	return doc.ToModel(), nil
}

func (gr *groupRepository) Save(ctx context.Context, g domain.Group) error {
	doc := &Group{}
	doc.FromModel(g)
	filter := bson.D{{"identifier", g.ID()}}
	err := gr.collection().FindOneAndReplace(ctx, filter, doc).Decode(&Group{})
	if err == mongo.ErrNoDocuments {
		return app.ErrGroupNotFound
	}
	return err
}

func (gr *groupRepository) collection() *mongo.Collection {
	return gr.client.Database(gr.db).Collection("groups")
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
