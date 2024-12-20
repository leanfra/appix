package data

import (
	"fmt"

	"gorm.io/gorm"
)

const resTeamTable = "res_teams"
const resTagTable = "res_tags"
const resProductTable = "res_products"
const resFeatureTable = "res_features"

type ResTeam struct {
	ID      uint32 `gorm:"primaryKey;autoIncrement"`
	ResType string `gorm:"index:idx_team_res_type_res_id_team_id,unique"`
	ResID   uint32 `gorm:"index:idx_team_res_type_res_id_team_id,unique"`
	TeamID  uint32 `gorm:"index:idx_team_res_type_res_id_team_id,unique"`
}

type ResTag struct {
	ID      uint32 `gorm:"primaryKey;autoIncrement"`
	ResType string `gorm:"index:idx_tag_res_type_res_id_tag_id,unique"`
	ResID   uint32 `gorm:"index:idx_tag_res_type_res_id_tag_id,unique"`
	TagID   uint32 `gorm:"index:idx_tag_res_type_res_id_tag_id,unique"`
}

type ResFeature struct {
	ID        uint32 `gorm:"primaryKey;autoIncrement"`
	ResType   string `gorm:"index:idx_feature_res_type_res_id_feature_id,unique"`
	ResID     uint32 `gorm:"index:idx_feature_res_type_res_id_feature_id,unique"`
	FeatureID uint32 `gorm:"index:idx_feature_res_type_res_id_feature_id,unique"`
}

type ResProduct struct {
	ID        uint32 `gorm:"primaryKey;autoIncrement"`
	ResType   string `gorm:"index:idx_product_res_type_res_id_product_id,unique"`
	ResID     uint32 `gorm:"index:idx_product_res_type_res_id_product_id,unique"`
	ProductID uint32 `gorm:"index:idx_product_res_type_res_id_product_id,unique"`
}

func listClassIds(tx *gorm.DB,
	model interface{}, field string, resType string, resID uint32) ([]uint32, error) {

	var ids []uint32
	if r := tx.Model(model).
		Where("res_type = ? and res_id = ?", resType, resID).Pluck(field, &ids); r.Error != nil {

		return nil, r.Error
	}
	return ids, nil

}

func updateClass(tx *gorm.DB,
	model interface{}, field string, delIDs, newIDs []uint32, resType string, resID uint32) error {

	if err := deleteClass(tx, model, field, delIDs, resType, resID); err != nil {
		return err
	}
	if err := createClass(tx, model, newIDs, resType, resID); err != nil {
		return err
	}
	return nil
}

func deleteClass(tx *gorm.DB,
	model interface{}, field string, ids []uint32, resType string, resID uint32) error {

	if len(ids) == 0 {
		return nil
	}
	str_del := fmt.Sprintf("res_type = ? and res_id = ? and %s in (?) ", field)
	r := tx.Model(model).Delete(str_del, resType, resID, ids)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func createClass(tx *gorm.DB, model interface{}, ids []uint32, resType string, resID uint32) error {
	if len(ids) == 0 {
		return nil
	}
	var newItems []interface{}
	switch model.(type) {
	case *ResTeam:
		for _, id := range ids {
			newItems = append(newItems, ResTeam{
				ResType: resType,
				ResID:   resID,
				TeamID:  id,
			})
		}
	case *ResProduct:
		for _, id := range ids {
			newItems = append(newItems, ResProduct{
				ResType:   resType,
				ResID:     resID,
				ProductID: id,
			})
		}
	case *ResTag:
		for _, id := range ids {
			newItems = append(newItems, ResTag{
				ResType: resType,
				ResID:   resID,
				TagID:   id,
			})
		}
	case *ResFeature:
		for _, id := range ids {
			newItems = append(newItems, ResFeature{
				ResType:   resType,
				ResID:     resID,
				FeatureID: id,
			})
		}
	default:
		return fmt.Errorf("unknown model type")
	}
	return tx.Create(newItems).Error
}
