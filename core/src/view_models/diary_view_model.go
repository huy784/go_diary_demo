// 视图模型层
// 定义响应数据结构，用于返回给客户端
package view_models

import "time"

// DiaryViewModel 日记视图模型
// 用于API响应的数据结构
type DiaryViewModel struct {
	ID        int64     `json:"id"`         // 日记ID
	Title     string    `json:"title"`      // 日记标题
	Content   string    `json:"content"`    // 日记内容
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
