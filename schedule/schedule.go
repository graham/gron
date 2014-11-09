package schedule

import (
	"time"
)

type Schedule struct {
	first_run time.Time
	last_run  time.Time
	recurs    bool
	query     string
}

func (s *Schedule) CalculateNext() (nextrun time.Time, err error) {
	return time.Now(), nil
}

func (s *Schedule) Run(arguments []string) (result string, err error) {
	return "", nil
}
