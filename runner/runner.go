package runner

import (
	"github.com/adhocore/gronx/pkg/tasker"
	"log"
	"urlcron/metric"
	"urlcron/schedule"
)

// Verbose controls runner verbosity
var Verbose = false

// Runner runs schedule loaded from loader
type Runner struct {
	loader        schedule.Loader
	metricCounter *metric.Counter
}

// New creates new runner with given schedule.Loader
func New(loader schedule.Loader) *Runner {
	return &Runner{
		loader: loader,
	}
}

// Run loads and starts schedule
func (r Runner) Run() {
	list, err := r.loader.List()
	if err != nil {
		panic(err)
	}
	if len(list) == 0 {
		log.Println("No tasks are configured. Use either file mount or environment variable to pass task configuration.")
		return
	}

	taskr := tasker.New(tasker.Option{
		Verbose: Verbose,
	})

	for _, item := range list {
		taskr.Task(item.Schedule, item.Exec)
	}

	taskr.Run()
}
