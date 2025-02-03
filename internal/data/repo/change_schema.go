package repo

type ChangeInfo struct {
	CreatedAt int64  `gorm:"column:created_at,type:bigint;immutable"`
	UpdatedAt int64  `gorm:"column:updated_at,type:bigint;"`
	CreatedBy string `gorm:"column:created_by,type:varchar(255)"`
	UpdatedBy string `gorm:"column:updated_by,type:varchar(255)"`
}
