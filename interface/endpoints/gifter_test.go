package endpoints_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/interface/endpoints"
	"github.com/mattreidarnold/gifter/interface/presenters"
	"github.com/mattreidarnold/gifter/test"
	"github.com/mattreidarnold/gifter/test/mocks"
)

func Test_MakeAddGifter_WhenIDGeneratorFails(t *testing.T) {
	ErrGenerateIDNeverWorks := errors.New("generate id never works")

	d := &app.Dependencies{
		GenerateID: func() (string, error) {
			return "", ErrGenerateIDNeverWorks
		},
	}

	addGifter := endpoints.MakeAddGifter(d)
	req := presenters.AddGifterRequest{}

	_, err := addGifter(context.Background(), req)

	test.AssertErrorIs(t, err, ErrGenerateIDNeverWorks)
}

func Test_MakeAddGifter_WhenIDGeneratorSucceeds(t *testing.T) {
	gifterID := "test-gifter-id"
	name := "test-name"
	groupID := "test-group-id"

	want := presenters.AddGifterResponse{
		GifterID: gifterID,
		GroupID:  groupID,
		Name:     name,
	}

	ctx := context.Background()

	req := presenters.AddGifterRequest{
		GroupID: groupID,
		Name:    name,
	}

	cmd := app.NewCommandMessage(domain.AddGifterCommand{
		Name:     name,
		GifterID: gifterID,
		GroupID:  groupID,
	})

	mockMB := &mocks.MessageBus{}
	mockMB.On("Handle", ctx, cmd).Return(nil).Once()
	defer mockMB.AssertExpectations(t)

	d := &app.Dependencies{
		GenerateID: func() (string, error) { return gifterID, nil },
		MessageBus: mockMB,
	}

	addGifter := endpoints.MakeAddGifter(d)

	got, err := addGifter(context.Background(), req)

	test.AssertNil(t, err)
	test.AssertEqual(t, got, want)

}
