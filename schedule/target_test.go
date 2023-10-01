package schedule

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseTargetFromString(t *testing.T) {
	t0 := NewTargetFromString("https://mobilly.lv POST <Header-1:Value 1> <Header-2: Value 2> {\"id\":123}")
	assert.Equal(t, "https://mobilly.lv", t0.Url)
	assert.Equal(t, "POST", t0.Method)
	assert.Equal(t, "{\"id\":123}", t0.Payload)
	assert.Len(t, t0.Headers, 2)
	assert.Equal(t, "Value 1", t0.Headers["Header-1"])
	assert.Equal(t, "Value 2", t0.Headers["Header-2"])

	t1 := NewTargetFromString("https://mobilly.lv?query=123")
	assert.Equal(t, "https://mobilly.lv?query=123", t1.Url)
	assert.Equal(t, "GET", t1.Method)
	assert.Equal(t, "", t1.Payload)
	assert.Len(t, t1.Headers, 0)

	t2 := NewTargetFromString("https://mobilly.lv?query=123 POST")
	assert.Equal(t, "https://mobilly.lv?query=123", t2.Url)
	assert.Equal(t, "POST", t2.Method)
	assert.Equal(t, "", t2.Payload)
	assert.Len(t, t2.Headers, 0)

	t3 := NewTargetFromString("https://mobilly.lv?query=123&q=abc <Header-1 : Value 1>")
	assert.Equal(t, "https://mobilly.lv?query=123&q=abc", t3.Url)
	assert.Equal(t, "GET", t3.Method)
	assert.Equal(t, "", t3.Payload)
	assert.Len(t, t3.Headers, 1)
	assert.Equal(t, "Value 1", t3.Headers["Header-1"])

	t4 := NewTargetFromString("https://mobilly.lv?query=123&q=abc POST {\"id\":123}")
	assert.Equal(t, "https://mobilly.lv?query=123&q=abc", t4.Url)
	assert.Equal(t, "POST", t4.Method)
	assert.Equal(t, "{\"id\":123}", t4.Payload)
	assert.Len(t, t4.Headers, 0)

}

func TestSubstituteWithEnv(t *testing.T) {
	_ = os.Setenv("COLOR", "yellow")
	_ = os.Setenv("SIZE", "big")

	assert.Equal(t, substituteWithEnv("Text with no substitutes"), "Text with no substitutes")
	assert.Equal(t, substituteWithEnv("My ${SIZE} jacket is ${COLOR}. I like ${SIZE} jackets."), "My big jacket is yellow. I like big jackets.")
	assert.Equal(t, substituteWithEnv("Text with unknown ${UNKNOWN} variable"), "Text with unknown  variable")

}
