package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	sync.RWMutex
	set map[string]Event
}

func NewCache() *Cache {
	cache := Cache{}
	cache.set = make(map[string]Event)
	return &cache
}

func (c *Cache) CreateEvent(e Event) {
	c.Lock()
	c.set[e.EventID] = e
	c.Unlock()
}

func (c *Cache) UpdateEvent(e Event) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.set[e.EventID]; !ok {
		return errors.New("no event")
	}

	c.set[e.EventID] = e
	return nil
}

func (c *Cache) DeleteEvent(EID string) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.set[EID]; !ok {
		return errors.New("no event")
	}

	delete(c.set, EID)
	return nil
}

func (c *Cache) GetEventsForDay(user string, date DateInfo) ([]Event, error) {
	c.RLock()
	defer c.RUnlock()

	noEventsErr := errors.New("no events for such day")

	eventsRequsted := make([]Event, 0)
	for _, event := range c.set {
		if event.UserID == user && event.Date.Day == date.Day && event.Date.Month == date.Month && event.Date.Year == date.Year {
			eventsRequsted = append(eventsRequsted, event)
		}
	}

	if len(eventsRequsted) == 0 {
		return nil, noEventsErr
	}

	return eventsRequsted, nil
}

func (c *Cache) GetEventsForWeek(user string, date DateInfo) ([]Event, error) {
	c.RLock()
	defer c.RUnlock()

	noEventsErr := errors.New("no events for such week")

	timeParsed, timeErr := DateInfoToTime(date)
	if timeErr != nil {
		return nil, timeErr
	}

	eventsRequsted := make([]Event, 0)
	for _, event := range c.set {

		if event.UserID != user {
			continue
		}

		currentTimeParsed, timeErr := DateInfoToTime(event.Date)
		if timeErr != nil {
			return nil, timeErr
		}

		userYear, userWeek := timeParsed.ISOWeek()
		currentYear, currentWeek := currentTimeParsed.ISOWeek()

		if currentYear == userYear && currentWeek == userWeek {
			eventsRequsted = append(eventsRequsted, event)
		}
	}

	if len(eventsRequsted) == 0 {
		return nil, noEventsErr
	}

	return eventsRequsted, nil
}

func (c *Cache) GetEventsForMonth(user string, date DateInfo) ([]Event, error) {
	c.RLock()
	defer c.RUnlock()

	noEventsErr := errors.New("no events for such month")

	eventsRequsted := make([]Event, 0)
	for _, event := range c.set {
		if event.UserID == user && event.Date.Year == date.Year && event.Date.Month == date.Month {
			eventsRequsted = append(eventsRequsted, event)
		}
	}

	if len(eventsRequsted) == 0 {
		return nil, noEventsErr
	}

	return eventsRequsted, nil
}
