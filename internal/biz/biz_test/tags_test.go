package biz_test

import (
	"appix/internal/biz"
	"appix/internal/data/repo"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTags(t *testing.T) {
	ctx := context.Background()
	tagsrepo := new(MockTagsRepo)
	apptagrepo := new(MockAppTagsRepo)
	hgtagrepo := new(MockHostgroupTagsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTagsUsecase(
		tagsrepo,
		nil,
		apptagrepo,
		hgtagrepo,
		txm,
	)

	// Test case: Validation fails
	tags := []*biz.Tag{{Key: "Key", Value: "validcode"}}
	err := usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Validation fails
	tags = []*biz.Tag{{Key: "0-name", Value: "validcode"}}
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	tags = []*biz.Tag{{Key: "-name", Value: "validcode"}}
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	tags = []*biz.Tag{{Key: "name-", Value: "validcode"}}
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	tags = []*biz.Tag{{Key: "nAme", Value: "validcode"}}
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Validation fails
	tags = []*biz.Tag{{Key: "Invalid Key", Value: "validcode"}}
	// should not call repo
	// repoCall := tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Conversion fails
	tags = []*biz.Tag{{Key: "ValidTag", Value: "invalid code"}}
	// should not call repo
	// repoCall = tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Invalid Leader
	tags = []*biz.Tag{{Key: "ValidTag", Value: "validcode"}}
	// not call repo
	// repoCall = tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Creation fails
	tags = []*biz.Tag{{Key: "valid", Value: "validcode"}}
	repoCall := tagsrepo.On("CreateTags", ctx, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.CreateTags(ctx, tags)
	assert.Error(t, err)
	repoCall.Unset()

	// Test case: Successful creation
	tagsrepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	tags = []*biz.Tag{{Key: "name", Value: "validcode"}}
	err = usecase.CreateTags(ctx, tags)
	assert.NoError(t, err)

	tags = []*biz.Tag{{Key: "name0", Value: "validcode0"}}
	err = usecase.CreateTags(ctx, tags)
	assert.NoError(t, err)

	tags = []*biz.Tag{{Key: "name-0", Value: "validcode-0"}}
	err = usecase.CreateTags(ctx, tags)
	assert.NoError(t, err)

	tags = []*biz.Tag{{Key: "name-0", Value: "1-validcode-0"}}
	err = usecase.CreateTags(ctx, tags)
	assert.NoError(t, err)
}

func TestUpdateTags(t *testing.T) {
	ctx := context.Background()
	tagsrepo := new(MockTagsRepo)
	apptagrepo := new(MockAppTagsRepo)
	hgtagrepo := new(MockHostgroupTagsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTagsUsecase(
		tagsrepo,
		nil,
		apptagrepo,
		hgtagrepo,
		txm,
	)

	// Test case: Validation fails
	tags := []*biz.Tag{{Id: 1, Key: "Name", Value: "validcode"}}
	err := usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Validation fails
	tags = []*biz.Tag{{Id: 1, Key: "0-name", Value: "validcode"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	tags = []*biz.Tag{{Id: 1, Key: "-name", Value: "validcode"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	tags = []*biz.Tag{{Id: 1, Key: "name-", Value: "validcode"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	tags = []*biz.Tag{{Id: 1, Key: "nAme", Value: "validcode"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Validation fails
	tags = []*biz.Tag{{Id: 1, Key: "Invalid Name", Value: "validcode"}}
	// should not call repo
	// repoCall := tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Conversion fails
	tags = []*biz.Tag{{Id: 1, Key: "Validtag", Value: "invalid code"}}
	// should not call repo
	// repoCall = tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Invalid Leader
	tags = []*biz.Tag{{Id: 1, Key: "Validtag", Value: "validcode"}}
	// not call repo
	// repoCall = tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	// no id
	tags = []*biz.Tag{{Key: "name", Value: "validcode"}}
	// not call repo
	// repoCall = tagRepo.On("CreateTags", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)

	// Test case: Creation fails
	tags = []*biz.Tag{{Id: 1, Key: "valid", Value: "validcode"}}
	repoCall := tagsrepo.On("UpdateTags", ctx, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.UpdateTags(ctx, tags)
	assert.Error(t, err)
	repoCall.Unset()

	// Test case: Successful creation
	tagsrepo.On("UpdateTags", ctx, mock.Anything).Return(nil)
	tags = []*biz.Tag{{Id: 1, Key: "name", Value: "validcode"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.NoError(t, err)

	tags = []*biz.Tag{{Id: 1, Key: "name0", Value: "validcode0"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.NoError(t, err)

	tags = []*biz.Tag{{Id: 1, Key: "name-0", Value: "validcode-0"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.NoError(t, err)

	tags = []*biz.Tag{{Id: 1, Key: "name-0", Value: "1-validcode-0"}}
	err = usecase.UpdateTags(ctx, tags)
	assert.NoError(t, err)
}

func TestDeleteTags(t *testing.T) {
	ctx := context.Background()
	tagsrepo := new(MockTagsRepo)
	apptagrepo := new(MockAppTagsRepo)
	hgtagrepo := new(MockHostgroupTagsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTagsUsecase(
		tagsrepo,
		nil,
		apptagrepo,
		hgtagrepo,
		txm,
	)

	// Test case: Validation fails
	tags := []uint32{}
	err := usecase.DeleteTags(ctx, tags)
	assert.Error(t, err)

	err = usecase.DeleteTags(ctx, nil)
	assert.Error(t, err)

	// Test case: failed on app-tag need check fail
	tags = []uint32{1, 2}
	apptagCall := apptagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(1), nil)
	hgtagrepoCall := hgtagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(0), nil)
	tagsCall := tagsrepo.On("DeleteTags", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteTags(ctx, tags)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	apptagCall.Unset()
	hgtagrepoCall.Unset()
	tagsCall.Unset()

	// Test case: failed on hostgroup-tag need check fail
	tags = []uint32{1, 2}
	apptagCall = apptagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(0), nil)
	hgtagrepoCall = hgtagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(1), nil)
	tagsCall = tagsrepo.On("DeleteTags", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteTags(ctx, tags)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	apptagCall.Unset()
	hgtagrepoCall.Unset()
	tagsCall.Unset()

	// Test case: failed on tags repo delete
	tags = []uint32{1, 2}
	apptagCall = apptagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(0), nil)
	hgtagrepoCall = hgtagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(0), nil)
	tagsCall = tagsrepo.On("DeleteTags", ctx, mock.Anything, mock.Anything).Return(errors.New("repo failed"))
	err = usecase.DeleteTags(ctx, tags)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	apptagCall.Unset()
	hgtagrepoCall.Unset()
	tagsCall.Unset()

	// Test case: success
	tags = []uint32{1, 2}

	apptagCall = apptagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(0), nil)
	hgtagrepoCall = hgtagrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTag, tags).Return(int64(0), nil)
	tagsCall = tagsrepo.On("DeleteTags", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteTags(ctx, tags)
	assert.NoError(t, err)

	apptagCall.Unset()
	hgtagrepoCall.Unset()
	tagsCall.Unset()
}

func TestGetTags(t *testing.T) {

	ctx := context.Background()
	tagsrepo := new(MockTagsRepo)
	apptagrepo := new(MockAppTagsRepo)
	hgtagrepo := new(MockHostgroupTagsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTagsUsecase(
		tagsrepo,
		nil,
		apptagrepo,
		hgtagrepo,
		txm,
	)
	// id == 0
	tag_id := uint32(0)
	_, err := usecase.GetTags(ctx, tag_id)
	t.Logf("err. %s", err)
	assert.Error(t, err)

	// repo error
	call := tagsrepo.On("GetTags", ctx, tag_id).Return(nil, errors.New("mock repo error"))
	_, err = usecase.GetTags(ctx, tag_id)
	t.Logf("err. %s", err)
	assert.Error(t, err)
	call.Unset()

	// success
	tag_id = uint32(1)
	db_tag := repo.Tag{
		ID:    tag_id,
		Key:   "tag1",
		Value: "code1",
	}
	biz_tag := &biz.Tag{
		Id:    tag_id,
		Key:   "tag1",
		Value: "code1",
	}
	call = tagsrepo.On("GetTags", ctx, tag_id).Return(&db_tag, nil)
	tag, err := usecase.GetTags(ctx, tag_id)
	t.Logf("tag. %+v", tag)
	assert.NoError(t, err)
	assert.Equal(t, biz_tag, tag)
	call.Unset()
}

func TestListTags(t *testing.T) {
	ctx := context.Background()
	tagsrepo := new(MockTagsRepo)
	apptagrepo := new(MockAppTagsRepo)
	hgtagrepo := new(MockHostgroupTagsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTagsUsecase(
		tagsrepo,
		nil,
		apptagrepo,
		hgtagrepo,
		txm,
	)

	// filter keys exceeds
	filter := biz.ListTagsFilter{
		Keys: []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7", "key8", "key9", "key10", "key11", "key12", "key13", "key14", "key15", "key16", "key17", "key18", "key19", "key20", "key21", "key22", "key23", "key24", "key25", "key26", "key27", "key28", "key29"},
	}
	_, err := usecase.ListTags(ctx, &filter)
	assert.Equal(t, biz.ErrFilterValuesExceedMax, err)
	// filter ids exceeds
	filter = biz.ListTagsFilter{
		Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
	}
	_, err = usecase.ListTags(ctx, &filter)
	assert.Equal(t, biz.ErrFilterValuesExceedMax, err)
	// filter kvs exceeds
	filter = biz.ListTagsFilter{
		Kvs: []string{"key1=value1", "key2=value2", "key3=value3", "key4=value4", "key5=value5", "key6=value6", "key7=value7", "key8=value8", "key9=value9", "key10=value1", "key11=value1", "key12=value1", "key13=value1", "key14=value1", "key15=value1", "key16=value1", "key17=value1", "key18=value1", "key19=value1", "key20=value1", "key21=value1", "key22=value1", "key23=value1"},
	}
	_, err = usecase.ListTags(ctx, &filter)
	assert.Equal(t, biz.ErrFilterValuesExceedMax, err)
	// filter malform kvs
	filter = biz.ListTagsFilter{
		Kvs: []string{"key1=value1", "key2=value2", "key3=value3", "key4=value4", "key5=value5", "key6=value6", "key7=value7", "key8=value8", "key9=value9", "key10=value1"},
	}
	_, err = usecase.ListTags(ctx, &filter)
	assert.Equal(t, biz.ErrFilterKVInvalid, err)
	// filter pagesize ==0
	filter = biz.ListTagsFilter{
		Page:     1,
		PageSize: 0,
	}
	_, err = usecase.ListTags(ctx, &filter)
	assert.Equal(t, biz.ErrFilterInvalidPagesize, err)
	// filter pagesize exceeds max
	filter = biz.ListTagsFilter{
		Page:     1,
		PageSize: 201,
	}
	_, err = usecase.ListTags(ctx, &filter)
	assert.Equal(t, biz.ErrFilterInvalidPagesize, err)
	// repo return error
	filter = biz.ListTagsFilter{
		Page:     1,
		PageSize: 50,
	}
	call := tagsrepo.On("ListTags", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("repo error")).Once()
	_, err = usecase.ListTags(ctx, &filter)
	assert.Equal(t, errors.New("repo error"), err)
	call.Unset()
	// success
	db_tags := []*repo.Tag{
		{
			ID:    1,
			Key:   "key1",
			Value: "kvs1",
		},
		{
			ID:    2,
			Key:   "key2",
			Value: "kvs2",
		},
	}
	biz_tags := []*biz.Tag{
		{
			Id:    1,
			Key:   "key1",
			Value: "kvs1",
		},
		{
			Id:    2,
			Key:   "key2",
			Value: "kvs2",
		},
	}
	tagsrepo.On("ListTags", mock.Anything, mock.Anything, mock.Anything).Return(db_tags, nil)
	tags, _ := usecase.ListTags(context.Background(), &filter)
	assert.Equal(t, biz_tags, tags)

}
