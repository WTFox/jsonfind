package scout

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/matryer/is"
)

func TestScout(t *testing.T) {
	parsedData := getParsedData()
	is := is.New(t)

	t.Run("instantiation", func(t *testing.T) {
		s := New("d", parsedData)
		is.Equal(reflect.TypeOf(s).String(), "scout.Scout")
	})

	t.Run("finds single values", func(t *testing.T) {
		s := New("bar", parsedData)
		s.DoSearch()

		found := s.Found()
		is.Equal(len(found), 1)
		is.Equal(s.Found()[0], ".foo")
	})

	t.Run("finds multiple values", func(t *testing.T) {
		s := New("c", parsedData)
		s.DoSearch()
		found := s.Found()
		sort.Strings(found)

		is.Equal(len(found), 2)

		is.Equal(found[0], ".a.b.0")
		is.Equal(found[1], ".baz")
	})
}

func getParsedData() map[string]interface{} {
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
