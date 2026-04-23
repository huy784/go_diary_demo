// 输入模型层
// 定义请求参数结构，用于接收客户端输入
package input_models

// CreateDiaryInput 创建日记请求参数
type CreateDiaryInput struct {
	Title   string `json:"title" binding:"required"`   // 日记标题，必填
	Content string `json:"content" binding:"required"` // 日记内容，必填
}

// UpdateDiaryInput 更新日记请求参数
type UpdateDiaryInput struct {
	Title   string `json:"title"`   // 日记标题，可选
	Content string `json:"content"` // 日记内容，可选
}
