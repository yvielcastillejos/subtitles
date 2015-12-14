package ssa

import (
	"testing"

	"github.com/martinlindhe/subber/caption"
	"github.com/martinlindhe/subber/testExtras"
	"github.com/stretchr/testify/assert"
)

func TestParseSsa(t *testing.T) {

	in := []byte(
		"[Events]\n" +
			"Format: Layer, Start, End, Style, Actor, MarginL, MarginR, MarginV, Effect, Text\n" +
			"Dialogue: 0,0:01:06.37,0:01:08.04,Default,,0000,0000,0000,,Honey, I'm home!\n" +
			"Dialogue: 0,0:01:09.05,0:01:10.69,Default,,0000,0000,0000,,Hi.\\n- Hi, love.\n")

	var expected = []caption.Caption{
		{
			1,
			testExtras.MakeTime(0, 1, 6, 370),
			testExtras.MakeTime(0, 1, 8, 40),
			[]string{"Honey, I'm home!"},
		},
		{
			2,
			testExtras.MakeTime(0, 1, 9, 50),
			testExtras.MakeTime(0, 1, 10, 690),
			[]string{"Hi.", "- Hi, love."},
		},
	}

	assert.Equal(t, expected, ParseSsa(in))
}

func TestParseSsaFormat(t *testing.T) {

	assert.Equal(t, -1, parseSsaFormat("xxx", "some"))

	assert.Equal(t, 9, parseSsaFormat("Format: Layer, Start, End, Style, Actor, MarginL, MarginR, MarginV, Effect, Text", "Text"))
}

func TestParseSsaTime(t *testing.T) {

	t1, _ := parseSsaTime("0:01:06.37")
	assert.Equal(t, testExtras.MakeTime(0, 1, 6, 370), t1)
}

func TestColumnCountFromSsaFormat(t *testing.T) {
	columns := columnCountFromSsaFormat("Format: Layer, Start, End, Style, Actor, MarginL, MarginR, MarginV, Effect, Text")

	assert.Equal(t, 10, columns)
}
