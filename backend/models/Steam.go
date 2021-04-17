package recipe

import "time"

type Steam struct {
	*Procedure
	Duration time.Duration
}
