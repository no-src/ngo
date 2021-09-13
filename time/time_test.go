package _time

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	testParseTime(t, "2019-6-4")
	testParseTime(t, "2019-6-4 16:45:12")
	testParseTime(t, "2019-06-04")
	testParseTime(t, "2019-06-04 16:45:12")
	testParseTime(t, "2019/6/4")
	testParseTime(t, "2019/6/4 16:45:12")
}

func testParseTime(t *testing.T, timeStr string) {
	_, err := ParseTime(timeStr)
	if err != nil {
		t.Fail()
		t.Logf("TestParseTime Failed ->%s", timeStr)
	}
}

func TestToDay(t *testing.T) {
	today := ToDay()
	now := time.Now()
	if today.Year() != now.Year() || today.Month() != now.Month() || today.Day() != now.Day() || today.Hour() != 0 || today.Minute() != 0 || today.Second() != 0 {
		t.Fail()
		t.Logf("TestToDay Failed today=%v", today)
	}
}

func TestGetMondayAndFriday(t *testing.T) {
	timeStr := "2019-06-04"
	date, err := ParseTime(timeStr)
	if err != nil {
		t.Fail()
		t.Logf("TestGetMondayAndFriday Failed date=%s", timeStr)
	}
	monday, friday := GetMondayAndFriday(date)
	if monday.Format("2006-01-02") != "2019-06-03" || friday.Format("2006-01-02") != "2019-06-07" {
		t.Fail()
		t.Logf("TestGetMondayAndFriday Failed date=%s monday=%v friday=%v", timeStr, monday, friday)
	}
}

func TestGetTimestamp(t *testing.T) {
	timeStr := "2019/6/4 13:27:12 UTC"
	date, err := ParseTime(timeStr)
	if err != nil {
		t.Fail()
		t.Logf("TestGetTimestamp Failed date=%s", timeStr)
	}
	timestamp := GetTimestamp(date)
	var expect int64 = 1559654832000
	if timestamp != expect {
		t.Fail()
		t.Logf("TestGetMondayAndFriday Failed date=%s timestamp=%d expect=%d", timeStr, timestamp, expect)
	}
}
