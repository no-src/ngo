package _time

import (
	"errors"
	"time"
)

// NTime 支持自定义时间格式的时间类型
type NTime time.Time

var ntFormat = time.RFC3339

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (nt NTime) MarshalJSON() ([]byte, error) {
	format := nt.GetFormat()
	t := time.Time(nt)
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(format)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, format)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (nt *NTime) UnmarshalJSON(data []byte) error {
	format := nt.GetFormat()
	t := time.Time(*nt)
	pt := &t
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	*pt, err = time.Parse(`"`+format+`"`, string(data))
	return err
}

// GetFormat 获取全局NTime时间格式
func (nt NTime) GetFormat() string {
	return ntFormat
}

// String 根据当前设置的时间格式获取时间字符串
func (nt NTime) String() string {
	format := nt.GetFormat()
	return time.Time(nt).Format(format)
}

// SetNTimeFormat 设置全局NTime时间格式
func SetNTimeFormat(format string) {
	ntFormat = format
}
