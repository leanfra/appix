package repo

const AppTagTable = "app_tags"

type AppTag struct {
	Id    uint32 `gorm:"primaryKey;autoIncrement"`
	AppID uint32 `gorm:"index:idx_app_id_tag_id,unique"`
	TagID uint32 `gorm:"index:idx_app_id_tag_id,unique"`
}

type AppTagsFilter struct {
	Ids      []uint32
	AppIds   []uint32
	TagIds   []uint32
	KVs      []string // k:v format
	Page     uint32
	PageSize uint32
}
