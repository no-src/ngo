package _time

import (
	"github.com/no-src/ngo/encoding/json"
	"testing"
	"time"
)

func TestTimeToNTime(t *testing.T) {
	format := "2006-01-02 15:04:05"
	SetNTimeFormat(format)
	now := time.Now()
	timeStrExpect := now.Format(format)
	nt := NTime(now)
	timeStrResult := nt.String()
	if timeStrExpect != timeStrResult {
		t.Fail()
		t.Logf("TestTimeToNTime Failed timeStrExpect=[%s] timeStrResult=[%s]", timeStrExpect, timeStrResult)
	}
}

func TestNTime_MarshalJSON(t *testing.T) {
	format := "2006-01-02 15:04:05"
	SetNTimeFormat(format)
	now := time.Now()
	timeStrExpect := "\"" + now.Format(format) + "\""
	nt := NTime(now)
	timeStrResult := _json.SerializeObject(nt)
	if timeStrExpect != timeStrResult {
		t.Fail()
		t.Logf("TestNTime_MarshalJSON Failed timeStrExpect=[%s] timeStrResult=[%s]", timeStrExpect, timeStrResult)
	}
}

func TestNTime_Property_MarshalJSON(t *testing.T) {
	format := "2006-01-02 15:04:05"
	SetNTimeFormat(format)
	now := time.Now()
	timeStrExpect := `{"Birth":"` + now.Format(format) + `"}`
	p := &Person{
		Birth: NTime(now),
	}
	timeStrResult := _json.SerializeObject(p)
	if timeStrExpect != timeStrResult {
		t.Fail()
		t.Logf("TestNTime_Property_MarshalJSON Failed timeStrExpect=[%s] timeStrResult=[%s]", timeStrExpect, timeStrResult)
	}
}

func TestNTime_Property_UnmarshalJSON(t *testing.T) {
	format := "2006-01-02 15:04:05"
	SetNTimeFormat(format)
	now := time.Now()
	timeStrExpect := `{"Birth":"` + now.Format(format) + `"}`
	p := new(Person)
	err := _json.DeserializeObject(timeStrExpect, p)
	if err != nil || p == nil {
		t.Fail()
		t.Logf("TestNTime_Property_UnmarshalJSON Failed timeStrExpect=[%s]", timeStrExpect)
	}
}

type Person struct {
	Birth NTime
}
