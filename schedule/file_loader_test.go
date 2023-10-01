package schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCronStringToSchedule(t *testing.T) {
	s0 := ParseSchedule("* * * * * https://addr.net POST <Content-type: application/json> {\"token\":\"_TOKEN_\"}")
	assert.Equal(t, "* * * * *", s0.Schedule)
	assert.Equal(t, "https://addr.net", s0.Target.Url)
	assert.Equal(t, "POST", s0.Target.Method)
	assert.Equal(t, "{\"token\":\"_TOKEN_\"}", s0.Target.Payload)
	assert.Len(t, s0.Target.Headers, 1)
	assert.Equal(t, "application/json", s0.Target.Headers["Content-type"])

	s1 := ParseSchedule("*/5 2 * * *      https://addr.net")
	assert.Equal(t, "*/5 2 * * *", s1.Schedule)
	assert.Equal(t, "https://addr.net", s1.Target.Url)

	s2 := ParseSchedule("*/5 2 * * *      https://addr.net?q=1")
	assert.Equal(t, "*/5 2 * * *", s2.Schedule)
	assert.Equal(t, "https://addr.net?q=1", s2.Target.Url)

	s3 := ParseSchedule("# This is the comment")
	assert.Nil(t, s3)
}
