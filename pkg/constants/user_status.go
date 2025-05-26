package constants

// 用户状态
type UserStatus int

const (
	UserStatusInactive UserStatus = 0 // 未激活
	UserStatusActive   UserStatus = 1 // 正常
	UserStatusDisabled UserStatus = 2 // 禁用
	UserStatusDeleted  UserStatus = 3 // 已删除
)

// String 返回用户状态的字符串表示
func (s UserStatus) String() string {
	switch s {
	case UserStatusInactive:
		return "未激活"
	case UserStatusActive:
		return "正常"
	case UserStatusDisabled:
		return "禁用"
	case UserStatusDeleted:
		return "已删除"
	default:
		return "未知"
	}
}

// IsValid 检查用户状态是否有效
func (s UserStatus) IsValid() bool {
	return s >= UserStatusInactive && s <= UserStatusDeleted
}

// IsActive 检查用户是否处于活跃状态
func (s UserStatus) IsActive() bool {
	return s == UserStatusActive
}

// 用户角色
type UserRole int

const (
	UserRoleUser  UserRole = 1 // 普通用户
	UserRoleAdmin UserRole = 2 // 管理员
	UserRoleSuper UserRole = 3 // 超级管理员
)

// String 返回用户角色的字符串表示
func (r UserRole) String() string {
	switch r {
	case UserRoleUser:
		return "普通用户"
	case UserRoleAdmin:
		return "管理员"
	case UserRoleSuper:
		return "超级管理员"
	default:
		return "未知"
	}
}

// IsValid 检查用户角色是否有效
func (r UserRole) IsValid() bool {
	return r >= UserRoleUser && r <= UserRoleSuper
}

// HasPermission 检查角色是否有指定权限
func (r UserRole) HasPermission(requiredRole UserRole) bool {
	return r >= requiredRole
}

// 用户性别
type UserGender int

const (
	UserGenderUnknown UserGender = 0 // 未知
	UserGenderMale    UserGender = 1 // 男
	UserGenderFemale  UserGender = 2 // 女
)

// String 返回用户性别的字符串表示
func (g UserGender) String() string {
	switch g {
	case UserGenderMale:
		return "男"
	case UserGenderFemale:
		return "女"
	default:
		return "未知"
	}
}

// IsValid 检查用户性别是否有效
func (g UserGender) IsValid() bool {
	return g >= UserGenderUnknown && g <= UserGenderFemale
}
