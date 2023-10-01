package schedule

// Loader interface should be implemented by all schedule Loaders
type Loader interface {
	// List returns loaded list of schedules
	List() ([]*Schedule, error)
}
