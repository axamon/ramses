package main

import (
	"sync"
	"time"
)

// RemoveCallback is used to have a callback when an entry
// is removed from the cache.
type RemoveCallback func(key interface{}, value interface{})

// TTLMap implements a map where the values expire with TTLs
type TTLMap struct {
	entries    map[interface{}]interface{}
	eMutex     sync.Mutex
	schedule   map[interface{}]*time.Timer
	sMutex     sync.Mutex
	onRemoval  RemoveCallback
	defaultTTL time.Duration
}

// NewTTLMap creates a TTLMap with no callback function
func NewTTLMap(ttl time.Duration) *TTLMap {
	return NewTTLMapWithCallback(ttl, nil)
}

// NewTTLMapWithCallback will create a TTLMap with a callback function
func NewTTLMapWithCallback(ttl time.Duration, callback RemoveCallback) *TTLMap {
	return &TTLMap{
		make(map[interface{}]interface{}),
		sync.Mutex{},
		make(map[interface{}]*time.Timer),
		sync.Mutex{},
		callback,
		ttl,
	}
}

// clearSchedule removes expired entries from the schedule
func (ttlmap *TTLMap) clearSchedule(key interface{}) {
	ttlmap.sMutex.Lock()
	defer ttlmap.sMutex.Unlock()
	delete(ttlmap.schedule, key)
}

func (ttlmap *TTLMap) addEntry(key, value interface{}) {
	ttlmap.eMutex.Lock()
	defer ttlmap.eMutex.Unlock()
	ttlmap.entries[key] = value
}

// Add adds an entry to the map using the default TTL
func (ttlmap *TTLMap) Add(key, value interface{}) {
	ttlmap.AddWithTTL(key, value, ttlmap.defaultTTL)
}

// AddWithTTL adds an entry with a specified TTL value
func (ttlmap *TTLMap) AddWithTTL(key, value interface{}, ttl time.Duration) {
	ttlmap.sMutex.Lock()
	defer ttlmap.sMutex.Unlock()
	if ttlmap.schedule[key] != nil {
		// Reset the ttl for entries that exist already.
		ttlmap.schedule[key].Reset(ttl)
	} else {
		// create a timer and monitor for it to expire, then remove
		// the object from the entries
		ttlmap.schedule[key] = time.NewTimer(ttl)
		go func() {
			<-ttlmap.schedule[key].C
			ttlmap.Remove(key)
		}()
	}
	// Add the entry
	ttlmap.addEntry(key, value)
}

// Remove deletes an entry from the entries
func (ttlmap *TTLMap) Remove(key interface{}) {
	ttlmap.eMutex.Lock()
	defer ttlmap.eMutex.Unlock()
	delete(ttlmap.entries, key)

	// remember to clean up the schedule
	ttlmap.clearSchedule(key)
}

// Get returns a value from the map
func (ttlmap *TTLMap) Get(key interface{}) (value interface{}) {
	ttlmap.eMutex.Lock()
	defer ttlmap.eMutex.Unlock()

	if value, present := ttlmap.entries[key]; present {
		return value
	}

	return nil
}

// GetAll returns the whole map
func (ttlmap *TTLMap) GetAll() (copied map[interface{}]interface{}) {
	ttlmap.eMutex.Lock()
	defer ttlmap.eMutex.Unlock()

	copied = make(map[interface{}]interface{})
	for key, value := range ttlmap.entries {
		copied[key] = value
	}

	return copied
}