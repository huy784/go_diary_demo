// 日记数据模型
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Diary 日记数据模型
type Diary struct {
	ID               string    `gorm:"primaryKey;type:varchar(36)"`
	UserIdentityGuid string    `gorm:"type:varchar(50);not null;index"`
	Title            string    `gorm:"size:255;not null"`
	Content          string    `gorm:"type:text;not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

// BeforeCreate GORM 钩子，创建前生成 UUID
func (d *Diary) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.New().String()
	}
	return nil
}
