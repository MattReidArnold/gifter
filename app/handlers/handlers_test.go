package handlers_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/app/handlers"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_MakeAddGifter_ReturnsCorrectMessageType(t *testing.T) {
	expected := reflect.TypeOf(domain.AddGifterCommand{})
	actual, _ := handlers.MakeAddGifter(&app.Dependencies{})

	assert.Equal(t, expected, actual)
}
func Test_MakeAddGifter_WhenMessageIsNotAddGifterCommand(t *testing.T) {
	expected := app.ErrInvalidMessageTypeForHandler

	_, handler := handlers.MakeAddGifter(&app.Dependencies{})

	msg := app.NewCommandMessage(struct {
		key string
	}{
		key: "value",
	})
	err := handler(context.Background(), msg)

	assert.ErrorIs(t, err, expected)
}

func Test_MakeAddGifter_WhenGroupNotFound(t *testing.T) {
	expected := app.ErrGroupNotFound

	groupID := "test-group-id"
	gifterID := "test-gifter-id"
	name := "test-gifter-name"

	ctx := context.Background()

	repo := &mocks.GroupRepository{}
	repo.On("Get", ctx, groupID).Return(nil, expected).Once()

	uow := &mocks.UnitOfWork{}
	uow.On("Groups").Return(repo)

	d := &app.Dependencies{
		GroupRepository: repo,
		UseUnitOfWork: func(c context.Context, f func(context.Context, app.UnitOfWork) error) error {
			return f(c, uow)
		},
	}

	_, handler := handlers.MakeAddGifter(d)

	msg := app.NewCommandMessage(domain.AddGifterCommand{
		GroupID:  groupID,
		GifterID: gifterID,
		Name:     name,
	})

	err := handler(ctx, msg)

	assert.ErrorIs(t, err, expected)
}
