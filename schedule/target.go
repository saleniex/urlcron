package schedule

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Target struct {
	Url     string
	Method  string
	Headers map[string]string
	Payload string
}

func NewTargetFromString(s string) *Target {
	re := regexp.MustCompile(`^(http\S+)\s*(\s(POST|GET))?\s*(\s+((<[^>]+>\s*)*))?\s*(\s([^<]*))?$`)
	matches := re.FindAllStringSubmatch(s, -1)

	if matches == nil {
		log.Println(s + " <- does not match")
		return nil
	}

	url := matches[0][1]
	method := matches[0][3]
	if method == "" {
		method = http.MethodGet
	}
	headers := matches[0][5]
	payload := matches[0][8]

	return &Target{
		Url:     url,
		Method:  method,
		Headers: headerParse(headers),
		Payload: payload,
	}
}

func headerParse(headerString string) map[string]string {
	result := make(map[string]string)
	re := regexp.MustCompile(`(<(([^>]+):([^>]+))>)`)
	matches := re.FindAllStringSubmatch(headerString, -1)
	for _, match := range matches {
		name := strings.Trim(match[3], " ")
		value := substituteWithEnv(strings.Trim(match[4], " "))
		result[name] = value
	}
	return result
}

func substituteWithEnv(s string) string {
	substitutedString := s

	re := regexp.MustCompile(`\${([A-Z_]+)}`)
	matches := re.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		envVariable := match[1]
		substitutedString = strings.Replace(
			substitutedString,
			fmt.Sprintf("${%s}", envVariable),
			os.Getenv(envVariable),
			-1)
	}

	return substitutedString
}
