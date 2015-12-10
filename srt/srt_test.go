package srt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeTime(h int, m int, s int, ms int) time.Time {
	return time.Date(0, 1, 1, h, m, s, ms*1000*1000, time.UTC)
}

func TestParseTime(t *testing.T) {

	assert.Equal(t, makeTime(18, 40, 22, 110), parseTime("18:40:22.110"))
	assert.Equal(t, makeTime(18, 40, 22, 110), parseTime("18:40:22,110"))
	assert.Equal(t, makeTime(18, 40, 22, 0), parseTime("18:40:22"))
}

func TestParseSrt(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"<i>Go ninja!</i>\n" +
		"\n" +
		"2\n" +
		"00:00:10,000 --> 00:00:11,000\n" +
		"<i>Subtitles By MrCool</i>\n" +
		"\n" +
		"3\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"<i>No ninja!</i>\n"

	var expected []Caption
	expected = append(expected, Caption{seq: 1, text: []string{"<i>Go ninja!</i>"}, start: makeTime(0, 0, 4, 630), end: makeTime(0, 0, 6, 18)})
	expected = append(expected, Caption{seq: 2, text: []string{"<i>Subtitles By MrCool</i>"}, start: makeTime(0, 0, 10, 0), end: makeTime(0, 0, 11, 0)})
	expected = append(expected, Caption{seq: 3, text: []string{"<i>No ninja!</i>"}, start: makeTime(0, 1, 9, 630), end: makeTime(0, 1, 11, 005)})

	assert.Equal(t, expected, ParseSrt(in))
}

func TestParseSrtCrlf(t *testing.T) {

	in := "1\n" +
		"00:00:04,630 --> 00:00:06,018\r\n" +
		"<i>Go ninja!</i>\r\n" +
		"\r\n"

	var expected []Caption
	expected = append(expected, Caption{seq: 1, text: []string{"<i>Go ninja!</i>"}, start: makeTime(0, 0, 4, 630), end: makeTime(0, 0, 6, 18)})

	assert.Equal(t, expected, ParseSrt(in))
}

func TestParseSrtUtf8Bom(t *testing.T) {

	in := "\ufeff1\n" +
		"00:00:04,630 --> 00:00:06,018\r\n" +
		"<i>Go ninja!</i>\r\n" +
		"\r\n"

	var expected []Caption
	expected = append(expected, Caption{seq: 1, text: []string{"<i>Go ninja!</i>"}, start: makeTime(0, 0, 4, 630), end: makeTime(0, 0, 6, 18)})

	assert.Equal(t, expected, ParseSrt(in))
}

func TestRenderTime(t *testing.T) {

	cap := Caption{seq: 1, text: []string{"<i>Go ninja!</i>"}, start: makeTime(18, 40, 22, 110), end: makeTime(18, 41, 20, 123)}

	assert.Equal(t, "18:40:22,110 --> 18:41:20,123", cap.srtTime())
}

func TestRenderSrt(t *testing.T) {

	expected := "1\n" +
		"00:00:04,630 --> 00:00:06,018\n" +
		"<i>Go ninja!</i>\n" +
		"\n" +
		"2\n" +
		"00:01:09,630 --> 00:01:11,005\n" +
		"<i>No ninja!</i>\n\n"

	var in []Caption
	in = append(in, Caption{seq: 1, text: []string{"<i>Go ninja!</i>"}, start: makeTime(0, 0, 4, 630), end: makeTime(0, 0, 6, 18)})
	in = append(in, Caption{seq: 2, text: []string{"<i>No ninja!</i>"}, start: makeTime(0, 1, 9, 630), end: makeTime(0, 1, 11, 005)})

	assert.Equal(t, expected, RenderSrt(in))
}
