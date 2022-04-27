package scout

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/matryer/is"
)

func TestScoutWithArray(t *testing.T) {
	is := is.New(t)

	t.Run("instantiation", func(t *testing.T) {
		s := New("d", jsonArray())
		is.Equal(reflect.TypeOf(s).String(), "scout.Scout")
	})

	t.Run("finds single values", func(t *testing.T) {
		s := New("bar", jsonArray())
		found, err := s.DoSearch()
		is.NoErr(err)

		is.Equal(len(found), 1)
		is.Equal(found[0], ".[0].foo")
	})

	t.Run("finds multiple values", func(t *testing.T) {
		s := New("c", jsonArray())
		found, err := s.DoSearch()
		is.NoErr(err)
		sort.Strings(found)

		is.Equal(len(found), 2)
		is.Equal(found[0], ".[0].a.b[0]")
		is.Equal(found[1], ".[0].baz")
	})
}

func TestScoutWithMapping(t *testing.T) {
	is := is.New(t)

	t.Run("instantiation", func(t *testing.T) {
		s := New("d", jsonMapping())
		is.Equal(reflect.TypeOf(s).String(), "scout.Scout")
	})

	t.Run("finds single values", func(t *testing.T) {
		s := New("bar", jsonMapping())
		found, err := s.DoSearch()
		is.NoErr(err)

		is.Equal(len(found), 1)
		is.Equal(found[0], ".foo")
	})

	t.Run("finds multiple values", func(t *testing.T) {
		s := New("c", jsonMapping())
		found, err := s.DoSearch()
		is.NoErr(err)
		sort.Strings(found)

		is.Equal(len(found), 2)
		is.Equal(found[0], ".a.b[0]")
		is.Equal(found[1], ".baz")
	})
}

func jsonArray() []any {
	in := []byte(`
	[{
		"a": {"b": ["c", "x"]},
		"foo": "bar",
		"baz": "c"
	}]
	`)
	var raw []any
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	return raw
}

func jsonMapping() map[string]any {
	in := []byte(`
	{
		"a": {"b": ["c", "x"]},
		"foo": "bar",
		"baz": "c"
	}
	`)
	var raw map[string]any
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	return raw
}
