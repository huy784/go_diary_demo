// 数据模型层
package models

import "time"

// Diary 日记数据模型
// 对应数据库中的 diaries 表
type Diary struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"` // 主键，自增
	Title     string    `gorm:"size:255;not null"`       // 日记标题，最大255字符
	Content   string    `gorm:"type:text;not null"`       // 日记内容，文本类型
	CreatedAt time.Time `gorm:"autoCreateTime"`           // 创建时间，自动设置
	UpdatedAt time.Time `gorm:"autoUpdateTime"`           // 更新时间，自动设置
}
