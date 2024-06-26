package option_test

import (
	"encoding/json"
	"testing"

	"github.com/BooleanCat/option"
	"github.com/BooleanCat/option/internal/assert"
)

func TestMarshalSome(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(option.Some(4))
	assert.Nil(t, err)
	assert.Equal(t, string(data), "4")
}

func TestMarshalNone(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(option.None[int]())
	assert.Nil(t, err)
	assert.Equal(t, string(data), "null")
}

func TestMarshalSomeParsed(t *testing.T) {
	t.Parallel()

	type name struct {
		MiddleName option.Option[string] `json:"middle_name"`
	}

	data, err := json.Marshal(name{MiddleName: option.Some("Barry")})
	assert.Nil(t, err)
	assert.Equal(t, string(data), `{"middle_name":"Barry"}`)
}

func TestMarshalNoneParsed(t *testing.T) {
	t.Parallel()

	type name struct {
		MiddleName option.Option[string] `json:"middle_name"`
	}

	data, err := json.Marshal(name{MiddleName: option.None[string]()})
	assert.Nil(t, err)
	assert.Equal(t, string(data), `{"middle_name":null}`)
}

func TestUnmarshalSome(t *testing.T) {
	t.Parallel()

	var number option.Option[int]
	err := json.Unmarshal([]byte("4"), &number)
	assert.Nil(t, err)
	assert.Equal(t, number, option.Some(4))
}

func TestUnmarshalNone(t *testing.T) {
	t.Parallel()

	var number option.Option[int]
	err := json.Unmarshal([]byte("null"), &number)
	assert.Nil(t, err)
	assert.True(t, number.IsNone())
}

func TestUnmarshalEmpty(t *testing.T) {
	t.Parallel()

	type name struct {
		MiddleName option.Option[string] `json:"middle_name"`
	}

	var value name
	err := json.Unmarshal([]byte("{}"), &value)
	assert.Nil(t, err)
	assert.True(t, value.MiddleName.IsNone())
}

func TestUnmarshalError(t *testing.T) {
	t.Parallel()

	var number option.Option[int]
	err := number.UnmarshalJSON([]byte("not a number"))
	assert.NotNil(t, err)
	assert.True(t, number.IsNone())
}
