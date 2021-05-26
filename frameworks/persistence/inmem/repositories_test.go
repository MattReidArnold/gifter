package inmem_test

import (
	"context"
	"testing"

	"github.com/mattreidarnold/gifter/app"
	"github.com/mattreidarnold/gifter/domain"
	"github.com/mattreidarnold/gifter/frameworks/persistence/inmem"
	"github.com/mattreidarnold/gifter/test"
)

func Test_GroupRepository_Add_WhenGroupDoesNotExist(t *testing.T) {

	groupID := test.NewRandomID()

	group := domain.NewGroup(groupID, "test-group-name", 42, []domain.Gifter{
		domain.NewGifter("9876", "Bob"),
	})

	repo := inmem.NewGroupRepository()
	err := repo.Add(context.Background(), group)

	if err != nil {
		t.Fatal("failed saving group:", err)
	}

	inmemRepo, ok := repo.(*inmem.GroupRepo)
	if !ok {
		t.Fatal("wrong type of GroupRepository")
	}

	persistedGroup, ok := inmemRepo.Groups[groupID]
	if !ok {
		t.Fatal("group not persisted in repo")
	}

	test.AssertEqual(t, persistedGroup, group)
}

func Test_GroupRepository_Add_WhenGroupAlreadyIDExists(t *testing.T) {
	want := app.ErrGroupIDAlreadyExists

	groupID := test.NewRandomID()
	existingGroup := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})

	group := domain.NewGroup(groupID, "test-some-other-name", 22, []domain.Gifter{})

	repo := inmem.NewGroupRepository(existingGroup)
	err := repo.Add(context.Background(), group)

	test.AssertEqual(t, err, want)
}

func Test_GroupRepository_Get_WhenGroupDoesNotExist(t *testing.T) {
	want := app.ErrGroupNotFound

	groupID := test.NewRandomID()

	repo := inmem.NewGroupRepository()
	_, err := repo.Get(context.Background(), groupID)

	test.AssertEqual(t, err, want)
}

func Test_GroupRepository_Get_WhenGroupExists(t *testing.T) {
	groupID := test.NewRandomID()

	want := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})

	repo := inmem.NewGroupRepository(want)
	got, err := repo.Get(context.Background(), groupID)

	if err != nil {
		t.Fatalf("failed getting existing group: %e", err)
	}
	test.AssertEqual(t, got, want)
}

func Test_GroupRepository_Save_WhenGroupDoesNotExist(t *testing.T) {
	want := app.ErrGroupNotFound

	groupID := test.NewRandomID()
	repo := inmem.NewGroupRepository()
	group := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})

	err := repo.Save(context.Background(), group)

	test.AssertEqual(t, err, want)

	inmemRepo, ok := repo.(*inmem.GroupRepo)
	if !ok {
		t.Fatal("wrong type of GroupRepository")
	}

	_, persisted := inmemRepo.Groups[groupID]
	if persisted {
		t.Fatal("group should not be persisted")
	}
}

func Test_GroupRepository_Save_WhenGroupExists(t *testing.T) {
	groupID := test.NewRandomID()

	group := domain.NewGroup(groupID, "test-group-name", 100, []domain.Gifter{})
	repo := inmem.NewGroupRepository(group)

	updatedGroup := domain.NewGroup(groupID, "test-updated-name", 250, []domain.Gifter{
		domain.NewGifter("test-gifter-id", "test-name"),
	})

	err := repo.Save(context.Background(), updatedGroup)

	if err != nil {
		t.Fatal("failed saving group:", err)
	}

	inmemRepo, ok := repo.(*inmem.GroupRepo)
	if !ok {
		t.Fatal("wrong type of GroupRepository")
	}

	peristedGroup, persisted := inmemRepo.Groups[groupID]
	if !persisted {
		t.Fatal("group should be persisted")
	}

	test.AssertEqual(t, updatedGroup, peristedGroup)
}
