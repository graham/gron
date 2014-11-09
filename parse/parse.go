package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/graham/gron/schedule"
)

// Since I allow the user to pass in the current time, I assume they
// will take care of any timezone issues. That way this library will
// be considerably easier, and will allow the implementer determine
// how they want to handle timezone data and modification.

func Parse(current_time time.Time, query string) (*schedule.Schedule, error) {
	fmt.Println("Query: %s", query)

	// Should be set when things have "starting next" or "after friday"
	// Should be set when things have "until friday" or "until the end of the month"

	return &schedule.Schedule{}, nil
}

func ParseISOFormat(now time.Time, chunk string) (time.Time, error) {
	/* this is just to parse a specific chunk:
	   next friday
	   every friday
	   first day of the month
	   5 days before the end of the month
	   3 hours before the end of the day
	   2 days after tomorrow at 9:00pm
	   three wednesdays from now
	   the second tuesday of next month
	*/
	var deadline string = `(^|\W)([0-9]{4})\-([0-9]{1,2})\-([0-9]{1,2})(\W|T|$)`
	re := regexp.MustCompile(deadline)

	match := re.FindAllStringSubmatch(chunk, -1)

	if len(match) == 0 {
		return time.Now(), errors.New("ISOParse: No date found.")
	}

	year, _ := strconv.Atoi(match[0][2])
	month, _ := strconv.Atoi(match[0][3])
	day, _ := strconv.Atoi(match[0][4])

	location, _ := time.LoadLocation("Local")

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location), nil
}
