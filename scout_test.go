package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func TestScout(t *testing.T) {
	parsedData := getParsedData()
	is := is.New(t)

	t.Run("instantiation", func(t *testing.T) {
		s := NewScout("d", parsedData)
		is.Equal(reflect.TypeOf(s).String(), "main.Scout")
	})

	t.Run("finds single values", func(t *testing.T) {
		s := NewScout("bar", parsedData)
		s.DoSearch()

		found := s.Found()
		is.Equal(len(found), 1)
		is.Equal(s.Found()[0], ".foo")
	})

	t.Run("finds multiple values", func(t *testing.T) {
		s := NewScout("c", parsedData)
		s.DoSearch()
		found := s.Found()

		is.Equal(len(found), 2)
		is.Equal(found[0], ".a.b.0")
		is.Equal(found[1], ".baz")

	})
}

func getParsedData() jsonData {
	in := []byte(`
	{
		"a": {"b": ["c", "x"]},
		"foo": "bar",
		"baz": "c"
	}
	`)
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	return raw
}
