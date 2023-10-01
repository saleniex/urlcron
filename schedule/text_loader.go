package schedule

import "strings"

type TextLoader struct {
	str string
}

func NewTextLoader(str string) *TextLoader {
	return &TextLoader{str: str}
}

func (s TextLoader) List() ([]*Schedule, error) {
	result := make([]*Schedule, 0)
	lines := strings.Split(s.str, "\n")
	for _, line := range lines {
		schedule := ParseSchedule(line)
		result = append(result, schedule)
	}
	return result, nil
}
