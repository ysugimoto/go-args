package args_test

import (
	"github.com/ysugimoto/go-args"
	"testing"
)

var c = args.NewContext(
	[]string{"foo", "bar", "baz", "100"},
	map[string]interface{}{
		"string":  "LoremIpsum",
		"integer": "10",
		"boolean": true,
	},
	map[string]interface{}{},
)

func TestString(t *testing.T) {
	v := c.String("string")
	if v != "LoremIpsum" {
		t.Errorf("Context.String() assertion failed: expect LoremIpsum, actual %s", v)
	}
}

func TestStringReturnsEmpty(t *testing.T) {
	v := c.String("notFound")
	if v != "" {
		t.Errorf("Context.String() assertion failed: expect empty string, actual %s", v)
	}
}

func TestInt(t *testing.T) {
	v := c.Int("integer")
	if v != 10 {
		t.Errorf("Context.Int() assertion failed: expect 10, actual %d", v)
	}
}

func TestIntReturnsZero(t *testing.T) {
	v := c.Int("notFound")
	if v != 0 {
		t.Errorf("Context.Int() assertion failed: expect 0, actual %v", v)
	}
}

func TestBool(t *testing.T) {
	v := c.Bool("boolean")
	if v != true {
		t.Errorf("Context.Bool() assertion failed: expect true, actual %v", v)
	}
}

func TestBoolReturnsFalse(t *testing.T) {
	v := c.Bool("notFound")
	if v != false {
		t.Errorf("Context.Bool() assertion failed: expect false, actual %v", v)
	}
}

func TestHas(t *testing.T) {
	v := c.Has("string")
	if v != true {
		t.Errorf("Context.Has() assertion failed: expect true, actual %v", v)
	}
}

func TestHasReturnsFalse(t *testing.T) {
	v := c.Has("notFound")
	if v != false {
		t.Errorf("Context.Has() assertion failed: expect false, actual %v", v)
	}
}

func TestStringAt(t *testing.T) {
	v := c.StringAt(0)
	if v != "foo" {
		t.Errorf("Context.StringAt() assertion failed: expect foo, actual %s", v)
	}
}

func TestIntAt(t *testing.T) {
	v := c.IntAt(3)
	if v != 100 {
		t.Errorf("Context.IntAt() assertion failed: expect 100, actual %v", v)
	}
}
