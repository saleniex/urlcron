package schedule

import (
	"bufio"
	"os"
)

// FileLoader provides schedule loading from file
type FileLoader struct {
	filePath string
}

// NewFileLoader creates new FileLoader with path to file where schedule is located
func NewFileLoader(filePath string) *FileLoader {
	return &FileLoader{filePath: filePath}
}

func (l FileLoader) List() ([]*Schedule, error) {
	file, err := os.Open(l.filePath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	result := make([]*Schedule, 0)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		schedule := ParseSchedule(line)
		if schedule != nil {
			result = append(result, schedule)
		}
	}

	return result, nil
}
