package scout

import (
	"fmt"
	"strings"
)

// Scout is the main type that handles the traversing of JSON paths.
type Scout struct {
	data       any
	found      []string
	lookingFor string
}

// New instantiates and returns a Scout
func New(lookingFor string, target any) Scout {
	return Scout{
		data:       target,
		lookingFor: lookingFor,
	}
}

// DoSearch searches the parsed JSON and returns an array of strings that correspond to
// found locations.
func (s *Scout) DoSearch() ([]string, error) {
	switch val := s.data.(type) {
	case []any:
		s.parseArray(val, "")
	case map[string]any:
		s.parseMap(val, "")
	default:
		return []string{}, fmt.Errorf("could not parse JSON structure")
	}
	return s.found, nil
}

func (s *Scout) parseMap(data map[string]any, path string) {
	for key, val := range data {
		keyPath := strings.Join([]string{path, key}, ".")
		switch valData := val.(type) {
		case map[string]any:
			s.parseMap(val.(map[string]any), keyPath)
		case []any:
			s.parseArray(val.([]any), keyPath)
		default:
			if s.isAMatch(fmt.Sprintf("%v", valData)) {
				s.found = append(s.found, keyPath)
			}
		}
	}
}

func (s *Scout) parseArray(anArray []any, path string) {
	for idx, val := range anArray {
		if path == "" {
			path = "."
		}
		keyPath := fmt.Sprintf("%s[%d]", path, idx)
		switch valData := val.(type) {
		case map[string]any:
			s.parseMap(val.(map[string]any), keyPath)
		case []any:
			s.parseArray(val.([]any), keyPath)
		default:
			if s.isAMatch(fmt.Sprintf("%v", valData)) {
				s.found = append(s.found, keyPath)
			}
		}
	}
}

func (s *Scout) isAMatch(found string) bool {
	return found == s.lookingFor
}
