package common

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// UUID 自定义UUID类型
type UUID struct {
	uuid.UUID
}

// NewUUID 生成新的UUID
func NewUUID() UUID {
	return UUID{uuid.New()}
}

// ParseUUID 解析UUID字符串
func ParseUUID(s string) (UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}
	return UUID{u}, nil
}

// MustParseUUID 解析UUID字符串，失败时panic
func MustParseUUID(s string) UUID {
	u, err := ParseUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

// String 返回UUID字符串表示
func (u UUID) String() string {
	return u.UUID.String()
}

// IsZero 检查是否为零值
func (u UUID) IsZero() bool {
	return u.UUID == uuid.Nil
}

// MarshalJSON 实现JSON序列化
func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON 实现JSON反序列化
func (u *UUID) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		*u = UUID{}
		return nil
	}

	parsed, err := uuid.Parse(str)
	if err != nil {
		return err
	}

	*u = UUID{parsed}
	return nil
}

// Scan 实现sql.Scanner接口，用于从数据库读取
func (u *UUID) Scan(value interface{}) error {
	if value == nil {
		*u = UUID{}
		return nil
	}

	switch v := value.(type) {
	case string:
		parsed, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		*u = UUID{parsed}
	case []byte:
		parsed, err := uuid.Parse(string(v))
		if err != nil {
			return err
		}
		*u = UUID{parsed}
	default:
		return fmt.Errorf("cannot scan %T into UUID", value)
	}

	return nil
}

// Value 实现driver.Valuer接口，用于写入数据库
func (u UUID) Value() (driver.Value, error) {
	if u.IsZero() {
		return nil, nil
	}
	return u.String(), nil
}

// GormDataType 实现GORM的DataType接口
func (UUID) GormDataType() string {
	return "char(36)"
}

// GormDBDataType 实现GORM的DBDataType接口
func (UUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql":
		return "char(36)"
	case "postgres":
		return "uuid"
	case "sqlite":
		return "text"
	default:
		return "char(36)"
	}
}

// BeforeCreate GORM钩子，在创建前自动生成UUID
func (u *UUID) BeforeCreate(tx *gorm.DB) error {
	if u.IsZero() {
		*u = NewUUID()
	}
	return nil
}

// ValidateUUID 验证UUID字符串格式
func ValidateUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

// UUIDFromString 从字符串创建UUID（兼容性函数）
func UUIDFromString(s string) (UUID, error) {
	return ParseUUID(s)
}
