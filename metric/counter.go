package metric

import (
	"fmt"
	"strings"
)

const FailTypeHttp = "http"
const FailTypeStatus = "status"

var metricLabel = "urlcron_call"

var counter *Counter

// IncResultUrlFailTypeCounter increases fail type counter for URL
func IncResultUrlFailTypeCounter(url, failType string) {
	key := fmt.Sprintf("%s{result=\"fail\", url=\"%s\", fail_type=\"%s\"}", metricLabel, url, failType)
	getCounter().Inc(key)
}

// IncResultUrlSuccessCounter increases success counter for URL
func IncResultUrlSuccessCounter(url string) {
	key := fmt.Sprintf("%s{result=\"success\", url=\"%s\"}", metricLabel, url)
	getCounter().Inc(key)
}

// getCounter ensures initialized counter variable
func getCounter() *Counter {
	if counter == nil {
		counter = &Counter{
			items: make(map[string]int),
		}
	}
	return counter
}

func PrometheusDump() string {
	return getCounter().PrometheusDump()
}

type Counter struct {
	items map[string]int
}

func (c *Counter) Inc(key string) {
	val, exists := c.items[key]
	if exists {
		c.items[key] = val + 1
	} else {
		c.items[key] = 1
	}
}

func (c *Counter) PrometheusDump() string {
	var builder strings.Builder
	for key, val := range c.items {
		builder.WriteString(fmt.Sprintf("%s %d\n", key, val))
	}
	s := builder.String()
	builder.Reset()

	return s
}
