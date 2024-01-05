package dao

import (
	"github.com/derailed/popeye/pkg/config"
	"github.com/derailed/popeye/pkg/model"
)

var InspectionService = &iInspectionService{}

type iInspectionService struct {
}

func (l *iInspectionService) InsertAddonInstance(m model.CdkmInspectionRecord) (uint, error) {
	db := config.PgDB
	tx := db.Model(&model.CdkmInspectionRecord{}).Create(&m)
	return m.Id, tx.Error
}
