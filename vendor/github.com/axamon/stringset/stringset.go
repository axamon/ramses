// Package stringset allows the creation of sets for strings
// that are concurrecy safe.
package stringset

import (
	"sync"
)

// StringSet is a set of unique strings.
// The lock sync.RWMutex allows to solve concurrency issues
type StringSet struct {
	m    map[string]struct{}
	lock sync.RWMutex
}

// NewStringSet creates a new set for strings.
func NewStringSet(strings ...string) *StringSet {
	res := &StringSet{
		m: map[string]struct{}{},
	}
	for _, s := range strings {
		res.Add(s)
	}
	return res
}

// Add adds a string to the set.
// If string is already in the set, it has no effect.
func (s *StringSet) Add(str string) {
	s.lock.Lock()
	s.m[str] = struct{}{}
	s.lock.Unlock()
}

// Exists checks if string exists in the set.
func (s *StringSet) Exists(str string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, exists := s.m[str]
	return exists
}

// Delete removes a string from the set.
func (s *StringSet) Delete(str string) {
	s.lock.Lock()
	delete(s.m, str)
	s.lock.Unlock()
}

// Strings returns a slice of strings in the set.
func (s *StringSet) Strings() []string {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	if n == 0 {
		return nil
	}
	// for efficiency, pre-allocate the array with known, final capacity
	// this avoids re-allocating underlying array in append
	res := make([]string, 0, n)
	s.lock.Lock()
	for str := range s.m {
		res = append(res, str)
	}
	s.lock.Unlock()
	return res
}

// Contains returns true if the given set contains all elements from the other set.
func (s *StringSet) Contains(other *StringSet) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	other.lock.Lock()
	defer other.lock.Unlock()
	for str := range other.m {
		if _, exists := s.m[str]; !exists {
			return false
		}
	}
	return true
}

// Union returns a new set which contains all elements of the previous ones.
func (s *StringSet) Union(other *StringSet) (union *StringSet) {
	var slen, otherlen int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		s.lock.Lock()
		slen = len(s.m)
		s.lock.Unlock()
	}()
	go func() {
		defer wg.Done()
		other.lock.Lock()
		otherlen = len(other.m)
		other.lock.Unlock()
	}()
	wg.Wait()

	l := slen + otherlen
	ret := &StringSet{
		m: make(map[string]struct{}, l),
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		s.lock.Lock()
		for str := range s.m {
			ret.lock.Lock()
			ret.m[str] = struct{}{}
			ret.lock.Unlock()
		}
		s.lock.Unlock()
	}()
	go func() {
		defer wg.Done()
		other.lock.Lock()
		for str := range other.m {
			ret.lock.Lock()
			ret.m[str] = struct{}{}
			ret.lock.Unlock()
		}
		other.lock.Unlock()
	}()
	wg.Wait()

	return ret
}

// Len returns the number of items in the set.
// Cannot be used in for loops.
func (s *StringSet) Len() int {
	s.lock.Lock()
	n := len(s.m)
	s.lock.Unlock()
	return n
}

// Pop removes and returns an arbitrary element from the set and removes it from the
// set. If the set was empty, this returns ("", false).
func (s *StringSet) Pop() (str string, ok bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.m) != 0 {
		// deletes only one value from the set and than exits
		for str = range s.m {
			delete(s.m, str)
			return str, true
		}
	}
	return "", false
}

// Difference returns a new set with all elements from the first set and no elements from the latter.
func (s *StringSet) Difference(other *StringSet) (diff *StringSet) {
	ret := &StringSet{
		m: map[string]struct{}{},
	}
	s.lock.Lock()
	for str := range s.m {
		ret.m[str] = struct{}{}
	}
	s.lock.Unlock()
	other.lock.Lock()
	for str := range other.m {
		delete(ret.m, str)
	}
	other.lock.Unlock()
	return ret
}

// Intersect returns a new set which contains only the elemets shared by both input sets.
func (s *StringSet) Intersect(other *StringSet) (intersection *StringSet) {
	var ret *StringSet
	var wg sync.WaitGroup

	var slen, otherlen int

	createIntersect := func(smallerlen int, smaller, greater *StringSet) (ret *StringSet) {
		ret = &StringSet{
			m: make(map[string]struct{}, smallerlen),
		}
		// Copy smaller set in ret
		smaller.lock.Lock()
		for str := range smaller.m {
			ret.m[str] = struct{}{}
		}
		smaller.lock.Unlock()

		greater.lock.Lock()
		defer greater.lock.Unlock()
		for element := range ret.m {
			// If element in smaller exists also in greater moves along
			if _, exists := greater.m[element]; exists {
				continue
			}
			// otherwise deletes it also from ret
			ret.lock.Lock()
			delete(ret.m, element)
			ret.lock.Unlock()
		}
		return ret
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		s.lock.Lock()
		slen = len(s.m)
		s.lock.Unlock()
	}()
	go func() {
		defer wg.Done()
		other.lock.Lock()
		otherlen = len(other.m)
		other.lock.Unlock()
	}()
	wg.Wait()
	switch {
	case slen >= otherlen:
		ret = createIntersect(otherlen, other, s)

	case slen < otherlen:
		ret = createIntersect(slen, s, other)
	}
	return ret
}
