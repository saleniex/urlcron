package metric

import (
	"fmt"
	"strings"
)

const SuccessCounter = "success_total"
const HttpErrorCounter = "error_total{type=\"http\"}"
const StatusErrorCounter = "error_total{type=\"status\"}"

var counter *Counter

// IncCounter increases
func IncCounter(key string) {
	if counter == nil {
		counter = NewCounter()
	}
	counter.Inc(key)
}

func PrometheusDump() string {
	return counter.PrometheusDump()
}

type Counter struct {
	items map[string]int
}

func NewCounter() *Counter {
	c := &Counter{
		items: make(map[string]int),
	}
	// Preset well known counters
	c.items[SuccessCounter] = 0
	c.items[HttpErrorCounter] = 0
	c.items[StatusErrorCounter] = 0

	return c
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
	return builder.String()
}
