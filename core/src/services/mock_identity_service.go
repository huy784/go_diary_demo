// 模拟身份服务实现
package services

// MockIdentityService 模拟身份服务实现
type MockIdentityService struct {
    userIdentity string
}

// NewMockIdentityService 创建模拟身份服务实例
func NewMockIdentityService(userIdentity string) IdentityService {
    return &MockIdentityService{
        userIdentity: userIdentity,
    }
}

// GetCurrentUserIdentity 获取当前用户身份标识
func (s *MockIdentityService) GetCurrentUserIdentity() string {
    return s.userIdentity
}
