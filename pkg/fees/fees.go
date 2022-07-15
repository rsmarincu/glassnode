package fees

import "time"

type Fee struct {
	Timestamp time.Time
	Value     float32
}
