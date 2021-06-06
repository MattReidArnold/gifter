package inmem_test

import (
	"context"
	"testing"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/frameworks/persistence/inmem"
	"github.com/mattreidarnold/gifter/test"
	"github.com/stretchr/testify/assert"
)

func Test_GroupRepository_Add_WhenGroupDoesNotExist(t *testing.T) {

	groupID := test.NewRandomID()

	group := domain.NewGroup(groupID, "test-group-name", 42, []domain.Gifter{
		domain.NewGifter("9876", "Bob"),
	})

	repo := inmem.NewGroupRepository()
	err := repo.Add(context.Background(), group)

	assert.Nil(t, err, "failed saving group")

	inmemRepo, ok := repo.(*inmem.GroupRepo)
	assert.True(t, ok, "inmem.GroupRepo type cast")

	persistedGroup, ok := inmemRepo.Groups[groupID]
	assert.True(t, ok, "group not persisted in repo")

	assert.Equal(t, persistedGroup, group)
}

func Test_GroupRepository_Add_WhenGroupAlreadyIDExists(t *testing.T) {
	want := app.ErrGroupIDAlreadyExists

	groupID := test.NewRandomID()
	existingGroup := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})

	group := domain.NewGroup(groupID, "test-some-other-name", 22, []domain.Gifter{})

	repo := inmem.NewGroupRepository(existingGroup)
	err := repo.Add(context.Background(), group)

	assert.Equal(t, err, want)
}

func Test_GroupRepository_Get_WhenGroupDoesNotExist(t *testing.T) {
	want := app.ErrGroupNotFound

	groupID := test.NewRandomID()

	repo := inmem.NewGroupRepository()
	_, err := repo.Get(context.Background(), groupID)

	assert.Equal(t, err, want)
}

func Test_GroupRepository_Get_WhenGroupExists(t *testing.T) {
	groupID := test.NewRandomID()

	want := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})

	repo := inmem.NewGroupRepository(want)
	got, err := repo.Get(context.Background(), groupID)

	assert.Nil(t, err, "failed getting existing group")
	assert.Equal(t, got, want)
}

func Test_GroupRepository_Save_WhenGroupDoesNotExist(t *testing.T) {
	want := app.ErrGroupNotFound

	groupID := test.NewRandomID()
	repo := inmem.NewGroupRepository()
	group := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})

	err := repo.Save(context.Background(), group)

	assert.ErrorIs(t, err, want)

	inmemRepo, ok := repo.(*inmem.GroupRepo)
	assert.True(t, ok, "inmem.GroupRepo type cast")

	_, persisted := inmemRepo.Groups[groupID]
	assert.False(t, persisted, "group should not be persisted")
}

func Test_GroupRepository_Save_WhenGroupExists(t *testing.T) {
	groupID := test.NewRandomID()

	group := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})
	repo := inmem.NewGroupRepository(group)

	updatedGroup := domain.NewGroup(groupID, "test-updated-name", 250, []domain.Gifter{
		domain.NewGifter("test-gifter-id", "test-name"),
	})

	err := repo.Save(context.Background(), updatedGroup)

	assert.Nil(t, err, "failed saving group")

	inmemRepo, ok := repo.(*inmem.GroupRepo)
	assert.True(t, ok, "inmem.GroupRepo type cast")

	peristedGroup, persisted := inmemRepo.Groups[groupID]
	assert.True(t, persisted, "group should be persisted")

	assert.Equal(t, updatedGroup, peristedGroup)
}
