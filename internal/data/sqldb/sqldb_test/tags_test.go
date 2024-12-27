package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var tagsRepo repo.TagsRepo

func initTagsRepo() {
	dataMem := getDataMem()
	tagsRepo, _ = sqldb.NewTagsRepoGorm(dataMem, logger)
}

func createBaseTags(t *testing.T, data []*repo.Tag) {
	initTagsRepo()
	if data == nil {
		data = []*repo.Tag{
			{Key: "test", Value: "value1"},
			{Key: "test", Value: "value2"},
		}
	}
	if err := tagsRepo.CreateTags(context.Background(), data); err != nil {
		t.Fatal(err)
	}

}

func TestTagsRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateTags_Success", testCreateTagsSuccess},
		{"CreateTags_Error", testCreateTagsError},
		{"UpdateTags_Success", testUpdateTagsSuccess},
		{"UpdateTags_Error", testUpdateTagsError},
		{"DeleteTags_Success", testDeleteTagsSuccess},
		{"DeleteTags_Error", testDeleteTagsError},
		{"GetTags_Success", testGetTagsSuccess},
		{"GetTags_Error", testGetTagsError},
		{"ListTags_emptyFilter_all", testListTags_emptyFilter_all},
		{"ListTags_id_partial", testListTags_id_partial},
		{"ListTags_page_partial", testListTags_page_partial},
		{"ListTags_keys_partial", testListTags_keys_partial},
		{"ListTags_kvs_partial", testListTags_kvs_partial},
		{"ListTags_nil_all", testListTags_nil_all},
		{"CountTags_partial", testCountTags_partial},
		{"CountTags_subpartial", testCountTags_subpartial},
		{"CountTags_all", testCountTags_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateTagsSuccess(t *testing.T) {

	initTagsRepo()

	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	err := tagsRepo.CreateTags(context.Background(), tags)
	assert.NoError(t, err)
}

func testCreateTagsError(t *testing.T) {
	createBaseTags(t, nil)
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value"},
	}
	err := tagsRepo.CreateTags(context.Background(), tags)
	assert.Error(t, err)
}

func testUpdateTagsSuccess(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)

	tags[0].Value = "value1"
	err := tagsRepo.UpdateTags(context.Background(), tags)
	assert.NoError(t, err)
}

func testUpdateTagsError(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)
	tags[0].Value = "value2"
	err := tagsRepo.UpdateTags(context.Background(), tags)
	assert.Error(t, err)
}

func testDeleteTagsSuccess(t *testing.T) {
	createBaseTags(t, nil)
	err := tagsRepo.DeleteTags(context.Background(), []uint32{1})
	assert.NoError(t, err)
}

func testDeleteTagsError(t *testing.T) {

	createBaseTags(t, nil)
	err := tagsRepo.DeleteTags(context.Background(), []uint32{99})
	assert.Error(t, err)
}

func testGetTagsSuccess(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)

	tag, err := tagsRepo.GetTags(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, tag)
	assert.Equal(t, uint32(1), tag.ID)
	assert.Equal(t, "test", tag.Key)
	assert.Equal(t, "value", tag.Value)
}

func testGetTagsError(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)
	tag, err := tagsRepo.GetTags(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func testListTags_emptyFilter_all(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)
	tags, err := tagsRepo.ListTags(context.Background(), nil, &repo.TagsFilter{})
	assert.NoError(t, err)
	assert.Len(t, tags, 2)
}
func testListTags_id_partial(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)
	_tags, err := tagsRepo.ListTags(context.Background(), nil, &repo.TagsFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Len(t, _tags, 1)
}

func testListTags_page_partial(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
	}
	createBaseTags(t, tags)
	filter := &repo.TagsFilter{
		Page:     2,
		PageSize: 1,
	}
	_tags, err := tagsRepo.ListTags(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, tags[1:], _tags)
}

func testListTags_keys_partial(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
		{Key: "test1", Value: "value"},
		{Key: "test1", Value: "value2"},
	}
	createBaseTags(t, tags)
	filter := &repo.TagsFilter{
		Keys: []string{"test"},
	}
	_tags, err := tagsRepo.ListTags(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, tags[:2], _tags)
}
func testListTags_kvs_partial(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
		{Key: "test1", Value: "value"},
		{Key: "test1", Value: "value2"},
	}
	createBaseTags(t, tags)
	filter := &repo.TagsFilter{
		Kvs: []string{"test:value2", "test1:value"},
	}
	_tags, err := tagsRepo.ListTags(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, tags[1:3], _tags)
}
func testListTags_nil_all(t *testing.T) {
	createBaseTags(t, nil)
	_tags, err := tagsRepo.ListTags(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(_tags))
}

func testCountTags_partial(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
		{Key: "test1", Value: "value"},
		{Key: "test1", Value: "value2"},
	}
	createBaseTags(t, tags)
	filter := &repo.TagsFilter{
		Ids: []uint32{2, 3},
	}
	count, err := tagsRepo.CountTags(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
func testCountTags_subpartial(t *testing.T) {
	tags := []*repo.Tag{
		{Key: "test", Value: "value"},
		{Key: "test", Value: "value2"},
		{Key: "test1", Value: "value"},
		{Key: "test1", Value: "value2"},
	}
	createBaseTags(t, tags)
	filter := &repo.TagsFilter{
		Ids: []uint32{2, 99},
	}
	count, err := tagsRepo.CountTags(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func testCountTags_all(t *testing.T) {
	createBaseTags(t, nil)
	count, err := tagsRepo.CountTags(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
