package test

import (
	"dev11/internal/cache"
	"testing"
)

func TestNewDateInfo(t *testing.T) {
	date, err := cache.NewDateInfo("2022-08-22")
	if err != nil {
		t.Fail()
	}

	t.Log(date)
}

func TestDateInfoToTime(t *testing.T) {
	date1 := cache.DateInfo{
		Year: "2022",
		Month: "02",
		Day: "01",
	}

	date2 := cache.DateInfo{
		Year: "2022",
		Month: "11",
		Day: "28",
	}

	time, err := cache.DateInfoToTime(date1)
	if err != nil {
		t.Fail()
	}
	t.Log(time, date1, err)

	time, err = cache.DateInfoToTime(date2)
	if err != nil {
		t.Fail()
	}
	t.Log(time, date2, err)
}