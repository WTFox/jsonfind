package scout

import (
	"fmt"
)

// Scout is the main type that handles the traversing of JSON paths.
type Scout struct {
	data       interface{}
	found      []string
	lookingFor string
}

// DoSearch searches the parsed JSON and returns an array of strings that correspond to
// found locations.
func (s *Scout) DoSearch() ([]string, error) {
	switch val := s.data.(type) {
	case []interface{}:
		s.parseArray(val, "")
	case map[string]interface{}:
		s.parseMap(val, "")
	default:
		return []string{}, fmt.Errorf("could not parse JSON structure")
	}
	return s.found, nil
}

// New instantiates and returns a Scout
func New(lookingFor string, target interface{}) Scout {
	return Scout{
		data:       target,
		lookingFor: lookingFor,
	}
}

func (s *Scout) parseMap(data map[string]interface{}, path string) {
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
		if path == "" {
			path = "."
		}
		keyPath := fmt.Sprintf("%s[%d]", path, idx)
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
