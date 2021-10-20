package core

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

type LocalTime time.Time

// 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

type LocalDate time.Time

// 写入 mysql 时调用
func (t LocalDate) Value() (driver.Value, error) {
	// 0001-01-01 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(DateFormat)), nil
}

// 检出 mysql 时调用
func (t *LocalDate) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalDate(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalDate) String() string {
	return time.Time(t).Format(DateFormat)
}

func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalDate(time.Time{})
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+DateFormat+`"`, string(data))
	*t = LocalDate(now)
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}