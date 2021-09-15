package afterclass

import (
	"errors"
)

//StringSet use map
type StringSet struct {
	elements map[string]string
}

//Add add element to set
func (s *StringSet) Add(ele string) error {
	err := s.Has(ele)
	if err != nil {
		return err
	}
	s.elements[ele] = ele
	return nil
}

//Delete delete ele from set
func (s *StringSet) Delete(ele string) error {
	err := s.Has(ele)
	if err == nil {
		return errors.New("ele not in set")
	}
	delete(s.elements, ele)
	return nil
}

//Size return the length of set
func (s StringSet) Size() int {
	return len(s.elements)
}

//Has check if ele in set
func (s *StringSet) Has(ele string) error {
	if _, ok := s.elements[ele]; ok {
		return errors.New("ele in set already")
	}
	return nil
}

//IntSet set with int elements
type IntSet struct {
	elements map[int]int
}

//Add add element to set
func (s *IntSet) Add(ele int) error {
	if s.Has(ele) {
		return errors.New("ele already in set")
	}
	s.elements[ele] = ele
	return nil
}

//Delete delete ele from set
func (s *IntSet) Delete(ele int) error {
	if !s.Has(ele) {
		return errors.New("ele not in set")
	}
	delete(s.elements, ele)
	return nil
}

//Size return the length of set
func (s IntSet) Size() int {
	return len(s.elements)
}

//Has check if ele in set
func (s *IntSet) Has(ele int) bool {
	if _, ok := s.elements[ele]; !ok {
		return false
	}
	return true
}
