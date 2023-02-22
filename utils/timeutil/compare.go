package timeutil

import "time"

func AnyOfTwo(timeA time.Time, timeB time.Time) time.Time {
	if timeA.IsZero() {
		return timeB
	}

	return timeA
}

func MaxOfTwo(timeA time.Time, timeB time.Time) time.Time {
	if timeA.Before(timeB) {
		return timeB
	}

	return timeA
}
