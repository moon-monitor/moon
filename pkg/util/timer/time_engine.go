package timer

import "time"

type Matcher interface {
	Match(time.Time) bool
}
