package main

import (
	"fmt"
)

type jsonData map[string]interface{}

type Scout struct {
	data       map[string]interface{}
	found      []string
	lookingFor string
}

func (s *Scout) DoSearch() {
	var path string
	s.parseMap(s.data, path)
	return
}

func (s Scout) Found() []string {
	return s.found
}

func NewScout(lookingFor string, target jsonData) Scout {
	return Scout{
		data:       target,
		lookingFor: lookingFor,
	}
}

func (s *Scout) parseMap(data jsonData, path string) {
	for key, val := range data {
		keyPath := fmt.Sprintf("%s.%s", path, key)
		switch valData := val.(type) {
		case map[string]interface{}:
			s.parseMap(val.(map[string]interface{}), keyPath)
		case []interface{}:
			s.parseArray(val.([]interface{}), keyPath)
		default:
			if fmt.Sprintf("%v", valData) == s.lookingFor {
				s.found = append(s.found, keyPath)
			}
		}
	}
	return
}

func (s *Scout) parseArray(anArray []interface{}, path string) {
	for idx, val := range anArray {
		keyPath := fmt.Sprintf("%s.%d", path, idx)
		switch valData := val.(type) {
		case map[string]interface{}:
			s.parseMap(val.(map[string]interface{}), keyPath)
		case []interface{}:
			s.parseArray(val.([]interface{}), keyPath)
		default:
			if fmt.Sprintf("%v", valData) == s.lookingFor {
				s.found = append(s.found, keyPath)
			}
		}
	}
}
