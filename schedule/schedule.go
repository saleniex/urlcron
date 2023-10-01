package schedule

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"urlcron/metric"
)

var httpClient *http.Client

type Schedule struct {
	Schedule string
	Target   *Target
}

func (s Schedule) Exec(ctx context.Context) (int, error) {
	req, err := http.NewRequest(s.Target.Method, s.Target.Url, bytes.NewBufferString(s.Target.Payload))
	if err != nil {
		metric.IncCounter(metric.HttpErrorCounter)
		return 1, err
	}

	for key, value := range s.Target.Headers {
		req.Header.Set(key, value)
	}

	resp, doErr := getHttpClient().Do(req)
	if doErr != nil {
		metric.IncCounter(metric.HttpErrorCounter)
		return 0, doErr
	}

	if resp.StatusCode != http.StatusOK {
		metric.IncCounter(metric.StatusErrorCounter)
		return 1, fmt.Errorf("URL %s responded with status %d", s.Target.Url, resp.StatusCode)
	}

	metric.IncCounter(metric.SuccessCounter)

	return 0, nil
}

func getHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return httpClient
}
