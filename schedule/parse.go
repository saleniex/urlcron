package schedule

import "regexp"

// ParseSchedule creates Schedule struct from string description
func ParseSchedule(cronStr string) *Schedule {
	if len(cronStr) == 0 || cronStr[0:1] == "#" {
		return nil
	}

	re := regexp.MustCompile(`^\s*(\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+)\s+(.+)$`)
	matches := re.FindAllStringSubmatch(cronStr, -1)
	if len(matches) == 0 || len(matches[0]) != 3 {
		return nil
	}

	target := NewTargetFromString(matches[0][2])
	if target == nil {
		return nil
	}

	return &Schedule{
		Schedule: matches[0][1],
		Target:   target,
	}
}
