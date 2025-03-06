package format

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

// MarshalJSON formats the LocalTime to a JSON string.
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))), nil
}

// Value implements the driver.Valuer interface for writing LocalTime to the database.
func (t *LocalTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return time.Time(*t), nil
}

// Scan implements the sql.Scanner interface for reading LocalTime from the database.
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("cannot convert %v to timestamp", v)
}

// IsZero checks if the LocalTime is the zero value.
func (t *LocalTime) IsZero() bool {
	return time.Time(*t).IsZero()
}
