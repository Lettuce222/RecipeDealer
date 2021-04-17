package recipe

import "time"

type Burn struct {
	*Procedure
	Duration time.Duration
	Heat Heat
}

type Heat string
