// 日记视图模型
package view_models

import "time"

// DiaryViewModel 日记视图模型
type DiaryViewModel struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
