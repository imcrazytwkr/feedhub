package timeutil

import "time"

func FormatAnyOfTwo(format string, timeA time.Time, timeB time.Time) string {
	timeToFormat := AnyOfTwo(timeA, timeB)
	if timeToFormat.IsZero() {
		return ""
	}

	return timeToFormat.Format(format)
}

func FormatMaxOfTwo(format string, timeA time.Time, timeB time.Time) string {
	timeToFormat := MaxOfTwo(timeA, timeB)
	if timeToFormat.IsZero() {
		return ""
	}

	return timeToFormat.Format(format)
}
