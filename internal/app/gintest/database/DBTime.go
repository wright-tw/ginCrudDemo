package database

import (
	"database/sql/driver"
	"fmt"
	"gintest/internal/app/gintest/constants"
	"strconv"
	"time"
)

func NewDBTime() DBTime {
	return DBTime{time.Now()}
}

type DBTime struct {
	time.Time
}

func (t DBTime) MarshalJSON() ([]byte, error) {
	str := t.Format(constants.DateTimeFormat)
	return []byte(`"` + str + `"`), nil
}

func (t DBTime) MarshalText() ([]byte, error) {
	str := t.Format(constants.DateTimeFormat)
	return []byte(str), nil
}

func (t DBTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *DBTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DBTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *DBTime) SubHours(h int) DBTime {
	hs := strconv.Itoa(h)
	d, _ := time.ParseDuration("-" + hs + "h")
	return DBTime{t.Add(d)}
}

func (t *DBTime) AddHours(h int) DBTime {
	hs := strconv.Itoa(h)
	d, _ := time.ParseDuration(hs + "h")
	return DBTime{t.Add(d)}
}
