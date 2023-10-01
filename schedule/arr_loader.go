package schedule

// ArrLoader provides schedule list loading from array
type ArrLoader struct {
	Schedules []*Schedule
}

func (l ArrLoader) List() ([]*Schedule, error) {
	return l.Schedules, nil
}
