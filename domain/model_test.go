package domain_test

import (
	"testing"

	"github.com/mattreidarnold/gifter/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Group_AddGifter_WhenGifterIsNotAlreadyInGroup(t *testing.T) {
	assert := assert.New(t)
	otherGifter := domain.NewGifter("test-other-gifter-id", "Old MacDonald")
	gifter := domain.NewGifter("test-gifter-id", "Ba Ba Blacksheep")
	want := []domain.Gifter{otherGifter, gifter}
	group := domain.NewGroup("test-group-id", "test-group-name", 100, []domain.Gifter{otherGifter})

	err := group.AddGifter(gifter)
	assert.NoError(err)

	got := group.Gifters()
	assert.Equal(got, want)
}
func Test_Group_AddGifter_WhenGifterIsAlreadyInGroup(t *testing.T) {
	assert := assert.New(t)
	gifter := domain.NewGifter("test-gifter-id", "Ba Ba Blacksheep")
	expGifters := []domain.Gifter{gifter}
	expErr := domain.ErrGifterAlreadyInGroup
	group := domain.NewGroup("test-group-id", "test-group-name", 100, []domain.Gifter{gifter})

	err := group.AddGifter(gifter)
	assert.ErrorIs(err, expErr)

	actGifters := group.Gifters()
	assert.Equal(expGifters, actGifters)
}
