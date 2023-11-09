package schedule

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"urlcron/metric"
)

var httpClient *http.Client

type Schedule struct {
	Schedule string
	Target   *Target
	request  *http.Request
}

// Exec executes the schedule
// Returns exit code and error
func (s *Schedule) Exec(ctx context.Context) (int, error) {
	req, err := s.scheduleRequest()
	if err != nil {
		return 1, err
	}

	resp, doErr := getHttpClient().Do(req)
	if doErr != nil {
		metric.IncResultUrlFailTypeCounter(s.Target.Url, metric.FailTypeHttp)
		return 0, doErr
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		metric.IncResultUrlFailTypeCounter(s.Target.Url, metric.FailTypeStatus)
		return 0, fmt.Errorf("URL %s responded with status %d", s.Target.Url, resp.StatusCode)
	}
	metric.IncResultUrlSuccessCounter(s.Target.Url)

	return 0, nil
}

func (s *Schedule) scheduleRequest() (*http.Request, error) {
	if s.request == nil {
		req, err := http.NewRequest(s.Target.Method, s.Target.Url, bytes.NewBufferString(s.Target.Payload))
		if err != nil {
			metric.IncResultUrlFailTypeCounter(s.Target.Url, metric.FailTypeHttp)
			return nil, err
		}
		for key, value := range s.Target.Headers {
			req.Header.Set(key, value)
		}
		s.request = req
	}
	return s.request, nil
}

func getHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return httpClient
}
