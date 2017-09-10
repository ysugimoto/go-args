package args_test

import (
	"github.com/ysugimoto/go-args"
	"testing"
)

func createArg() *args.Args {
	return args.New().
		Alias("string", "s", "string-value").
		Alias("int", "i", 10).
		Alias("bool", "b", false).
		Alias("flag", "f", nil)
}

func TestParseBeforeOptions(t *testing.T) {
	a := createArg()
	command := []string{
		"-s", "LoremIpsum",
		"-i", "100",
		"-b",
		"-f",
		"foo",
		"bar",
	}
	ctx := a.Parse(command)
	if ctx.String("string") != "LoremIpsum" {
		t.Error("String option error")
	}
	if ctx.Int("int") != 100 {
		t.Error("Int option error")
	}
	if ctx.Bool("bool") != true {
		t.Error("Bool option error")
	}
	if ctx.Has("flag") != true {
		t.Error("Nil option error")
	}
	if ctx.Len() != 2 {
		t.Error("Command length error")
	}
	if ctx.At(0) != "foo" {
		t.Error("Index 0 command error")
	}
	if ctx.At(1) != "bar" {
		t.Error("Index 0 command error")
	}
}

func TestParseAfterOptions(t *testing.T) {
	a := createArg()
	command := []string{
		"foo",
		"bar",
		"-s", "LoremIpsum",
		"-i", "100",
		"-b",
		"-f",
	}
	ctx := a.Parse(command)
	if ctx.String("string") != "LoremIpsum" {
		t.Error("String option error")
	}
	if ctx.Int("int") != 100 {
		t.Error("Int option error")
	}
	if ctx.Bool("bool") != true {
		t.Error("Bool option error")
	}
	if ctx.Has("flag") != true {
		t.Error("Nil option error")
	}
	if ctx.Len() != 2 {
		t.Error("Command length error")
	}
	if ctx.At(0) != "foo" {
		t.Error("Index 0 command error")
	}
	if ctx.At(1) != "bar" {
		t.Error("Index 0 command error")
	}
}
