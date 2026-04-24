// 身份服务接口定义
package services

// IdentityService 身份服务接口
// 提供用户身份相关功能
type IdentityService interface {
    // GetCurrentUserIdentity 获取当前用户身份标识
    GetCurrentUserIdentity() string
}
