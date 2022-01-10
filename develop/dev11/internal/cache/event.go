package cache

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	UserID      string   `json:"user_id"`
	EventID     string   `json:"event_id"`
	Description string   `json:"description"`
	Date        DateInfo `json:"date"`
}

func NewEvent(UID string, EID string, description string, date string) (Event, error) {
	if UID == "" || EID == "" {
		return Event{}, errors.New("incomplete query string")
	}

	if _, err := CheckDateFormat(date); err != nil {
		return Event{}, err
	}

	event := Event{}
	event.UserID = UID
	event.EventID = EID
	event.Description = description

	var err error
	event.Date, err = NewDateInfo(date)
	if err != nil {
		return event, err
	}

	return event, err
}

func CheckUIDAndDate(UID string, date string) (string, DateInfo, error) {

	_, err := CheckDateFormat(date)
	if UID == "" || date == "" || err != nil {
		return UID, DateInfo{}, err
	}

	dateParsed, err := NewDateInfo(date)
	if err != nil {
		err := errors.New("wrong date format")

		return UID, dateParsed, err
	}

	return UID, dateParsed, nil
} 

// YYYY-MM-DD
type DateInfo struct {
	Year  string
	Month string
	Day   string
}

func NewDateInfo(currentDate string) (DateInfo, error) {
	d, err := CheckDateFormat(currentDate)
	if err != nil {
		return DateInfo{}, err
	}

	date := DateInfo{}

	date.Year = d[0]
	date.Month = d[1]
	date.Day = d[2]

	return date, err
}

func DateInfoToTime(date DateInfo) (time.Time, error) {
	layout := "2006-01-02"

	dateString := fmt.Sprintf("%s-%s-%s", date.Year, date.Month, date.Day)

	timeParsed, err := time.Parse(layout, dateString)
	return timeParsed, err
}

// [0] - year, [1] - month, [2] - day
func CheckDateFormat(date string) ([]string, error) {
	dateFormatError := errors.New("wrong date format")

	if date == "" {
		return []string{}, dateFormatError
	}

	layout := "2006-01-02"

	_, err := time.Parse(layout, date)
	if err != nil {
		return []string{}, err
	}

	d := strings.Split(date, "-")
	if len(d) > 3 {
		return []string{}, dateFormatError
	}

	return d, nil
}
