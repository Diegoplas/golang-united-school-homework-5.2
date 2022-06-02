package cache

import (
	"time"
)

var infiniteDeadline time.Time = time.Date(3000, 12, 30, 23, 59, 59, 100, time.Local)

type Cache struct {
	cachedValues map[string]string
	deadlines    map[string]time.Time
}

func NewCache() Cache {
	return Cache{
		cachedValues: map[string]string{},
		deadlines:    map[string]time.Time{},
	}
}

func (c *Cache) Get(key string) (string, bool) { // comparar fechas y no desplegar si se venció
	value, ok := c.cachedValues[key]
	keyDeadline := c.deadlines[key]
	if ok {
		now := time.Now()
		if now.Before(keyDeadline) {
			return value, true
		} else {
			delete(c.cachedValues, key)
			delete(c.deadlines, key)
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) { //agregar timeline infinito
	c.cachedValues[key] = value
	c.deadlines[key] = infiniteDeadline
}

func (c *Cache) Keys() []string { // comparar fechas y no desplegar las vencidas
	allKeys := []string{}
	for key := range c.cachedValues {
		now := time.Now()
		if now.Before(c.deadlines[key]) {
			allKeys = append(allKeys, key)
		} else {
			delete(c.cachedValues, key)
			delete(c.deadlines, key)
		}
	}
	return allKeys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	now := time.Now()
	if now.Before(deadline) {
		c.deadlines[key] = deadline
		c.cachedValues[key] = value
	}
}
