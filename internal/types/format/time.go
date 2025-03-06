package format

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type LocalTime struct {
	time.Time
}

// `Scan` 处理数据库读取
func (t *LocalTime) Scan(value interface{}) error {
	if value == nil {
		*t = LocalTime{Time: time.Time{}}
		return nil
	}
	v, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot convert %v to LocalTime", value)
	}
	*t = LocalTime{Time: v}
	return nil
}

// `Value` 处理存入数据库的时间格式
func (t LocalTime) Value() (driver.Value, error) {
	return t.Format("2006-01-02 15:04:05"), nil
}

// `MarshalJSON` 控制 JSON 序列化格式
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format("2006-01-02 15:04:05"))
}
