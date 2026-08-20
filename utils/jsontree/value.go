package jsontree

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Value struct {
	t Type
	o Object
	a []*Value
	s string
	n float64
}

func (v *Value) Type() Type {
	if v == nil {
		return TypeNull
	}

	return v.t
}

func (v *Value) Get(keys ...string) *Value {
	if v == nil {
		return nil
	}

	for _, key := range keys {
		switch v.t {
		case TypeObject:
			v = v.o.Get(key)
			if v == nil {
				return nil
			}
		case TypeArray:
			n, err := strconv.Atoi(key)
			if err != nil || n < 0 || n >= len(v.a) {
				return nil
			}

			v = v.a[n]
		default:
			return nil
		}
	}

	return v
}

func (v *Value) GetObject(keys ...string) *Object {
	v = v.Get(keys...)
	if v.Type() != TypeObject {
		return nil
	}

	return &v.o
}

func (v *Value) GetArray(keys ...string) []*Value {
	v = v.Get(keys...)
	if v.Type() != TypeArray {
		return nil
	}

	return v.a
}

func (v *Value) GetString(keys ...string) string {
	v = v.Get(keys...)
	if v.Type() != TypeString {
		return ""
	}

	return v.s
}

func (v *Value) GetFloat64(keys ...string) float64 {
	v = v.Get(keys...)
	if v.Type() != TypeNumber {
		return 0
	}

	return v.n
}

func (v *Value) GetBool(keys ...string) bool {
	return v.Get(keys...).Type() == TypeTrue
}

func (v *Value) MarshalTo(dst []byte) []byte {
	if v == nil {
		return append(dst, "null"...)
	}

	switch v.t {
	case TypeNull:
		// Duplicating it just for clarity
		return append(dst, "null"...)
	case TypeTrue:
		return append(dst, "true"...)
	case TypeFalse:
		return append(dst, "false"...)
	case TypeNumber:
		b, err := json.Marshal(v.n)
		if err != nil {
			panic(fmt.Errorf("could not serialize number %f to json: %w", v.n, err))
		}

		return append(dst, b...)
	case TypeString:
		b, err := json.Marshal(v.s)
		if err != nil {
			panic(fmt.Errorf("could not serialize string %q to json: %w", v.s, err))
		}

		return append(dst, b...)
	case TypeArray:
		dst = append(dst, '[')
		for i, el := range v.a {
			if i > 0 {
				dst = append(dst, ',')
			}
			dst = el.MarshalTo(dst)
		}

		return append(dst, ']')
	case TypeObject:
		dst = append(dst, '{')
		i := 0

		// Cannot think of a better way to join a map without resorting to a lot of
		// redundant allocations
		for k, el := range v.o.kvs {
			if i > 0 {
				dst = append(dst, ',')
			}

			b, err := json.Marshal(k)
			if err != nil {
				panic(fmt.Errorf("could not serialize object key %q to json: %w", k, err))
			}

			dst = append(dst, b...)
			dst = append(dst, ':')
			dst = el.MarshalTo(dst)
			i++
		}

		return append(dst, '}')
	default:
		panic(fmt.Errorf("BUG: unknown Value type: %d", v.t))
	}
}

func (v *Value) MarshalJSON() ([]byte, error) {
	return v.MarshalTo(nil), nil
}

func (v *Value) UnmarshalJSON(data []byte) error {
	*v = Value{}

	i := 0
	for i < len(data) && isJSONSpace(data[i]) {
		i++
	}

	if i >= len(data) {
		return errors.New("unexpected end of JSON input")
	}

	switch data[i] {
	case 'n':
		if hasPrefixAt(data, "null", i) {
			v.t = TypeNull
			return nil
		}

		// `JSON.parse` in V8 accepts NaN case-insensitively
		fallthrough
	case 'N':
		if hasPrefixFoldAt(data, "NaN", i) {
			v.t = TypeNumber
			v.n = math.NaN()
			return nil
		}

		return unexpectedValueAt(data, i)
	case 't':
		if hasPrefixAt(data, "true", i) {
			v.t = TypeTrue
			return nil
		}

		return unexpectedValueAt(data, i)
	case 'f':
		if hasPrefixAt(data, "false", i) {
			v.t = TypeFalse
			return nil
		}

		return unexpectedValueAt(data, i)
	case '"':
		if err := json.Unmarshal(data, &v.s); err != nil {
			return wrapUnexpectedValueAt(err, i)
		}

		v.t = TypeString
		return nil
	case '[':
		var raw []json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			return wrapUnexpectedValueAt(err, i)
		}

		v.t = TypeArray
		v.a = make([]*Value, len(raw))
		for j, item := range raw {
			el := new(Value)
			if err := el.UnmarshalJSON(item); err != nil {
				return err
			}

			v.a[j] = el
		}

		return nil
	case '{':
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			return wrapUnexpectedValueAt(err, i)
		}

		v.t = TypeObject
		v.o.kvs = make(map[string]*Value, len(raw))
		for k, item := range raw {
			el := new(Value)
			if err := el.UnmarshalJSON(item); err != nil {
				return err
			}

			v.o.kvs[k] = el
		}

		return nil
	default:
		if err := json.Unmarshal(data, &v.n); err != nil {
			return wrapUnexpectedValueAt(err, i)
		}

		v.t = TypeNumber
		return nil
	}
}

func wrapUnexpectedValueAt(err error, p int) error {
	return fmt.Errorf("unexpected value found at position %d: %w", p, err)
}

func unexpectedValueAt(b []byte, p int) error {
	var m json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return wrapUnexpectedValueAt(err, p)
	}

	return fmt.Errorf("unexpected value found at index %d: %q", p, m)
}

func isJSONSpace(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r':
		return true
	default:
		return false
	}
}
