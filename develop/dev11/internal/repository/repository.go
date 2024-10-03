package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/mirhijinam/wb-l2/develop/dev11/internal/models"
)

type Calendar struct {
	lastEvent int
	events    map[int]models.Event
	mu        *sync.RWMutex
}

func New() *Calendar {
	return &Calendar{
		events:    make(map[int]models.Event),
		mu:        &sync.RWMutex{},
		lastEvent: 0,
	}
}

func (c *Calendar) CreateEvent(
	e models.Event,
) models.Event {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lastEvent++
	e.ID = c.lastEvent
	c.events[c.lastEvent] = e

	return e
}

func (c *Calendar) UpdateEvent(
	id int,
	name,
	data string,
	date time.Time,
) (models.Event, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.events[id]; !ok {
		return models.Event{}, errors.New("no such event")
	}

	e := models.Event{
		ID:   id,
		Name: name,
		Data: data,
		Date: date,
	}

	c.events[id] = e

	return c.events[id], nil
}

func (c *Calendar) DeleteEvent(
	id int,
) (models.Event, error) {
	c.mu.RLock()

	if _, ok := c.events[id]; !ok {
		c.mu.RUnlock()
		return models.Event{}, errors.New("no such event")
	}

	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	deleted := c.events[id]
	delete(c.events, id)

	return deleted, nil
}

func (c *Calendar) DailyEvents() []models.Event {
	var res []models.Event

	curYear, curMonth, curDay := time.Now().Date()

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, e := range c.events {
		y, m, d := e.Date.Date()

		if curYear == y && curMonth == m && curDay == d {
			res = append(res, e)
		}
	}

	return res
}

func (c *Calendar) WeeklyEvents() []models.Event {
	var res []models.Event

	curYear, curWeek := time.Now().ISOWeek()

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, e := range c.events {
		y, w := e.Date.ISOWeek()

		if curYear == y && curWeek == w {
			res = append(res, e)
		}
	}

	return res
}

func (c *Calendar) MonthlyEvents() []models.Event {
	var res []models.Event

	curYear, curMonth, _ := time.Now().Date()

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, e := range c.events {
		y, m, _ := e.Date.Date()

		if curYear == y && curMonth == m {
			res = append(res, e)
		}
	}

	return res
}
