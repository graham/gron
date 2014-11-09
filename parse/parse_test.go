package parse

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	now := time.Now()
	fmt.Println(ParseISOFormat(now, "2014-3-22"))
}
