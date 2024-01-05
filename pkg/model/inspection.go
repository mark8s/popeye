package model

import "gorm.io/gorm"

type CdkmInspectionRecord struct {
	Id        uint           `json:"id" gorm:"primarykey,AUTO_INCREMENT"`   // ID 主键
	TaskName  string         `gorm:"column:task_name"`                      // 任务名称
	Results   string         `gorm:"column:results"`                        // 结果
	CreatorId int64          `gorm:"column:creator_id"`                     // 创建者Id
	UpdatedAt int64          `json:"updatedAt" gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	CreatedAt int64          `json:"createdAt" gorm:"autoCreateTime:milli"` // 使用时间戳秒数填充创建时间
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;"`   // 删除标记
}

const LoggingAddonTableName = "cdkm_inspection_record"

func (CdkmInspectionRecord) TableName() string {
	return LoggingAddonTableName
}
