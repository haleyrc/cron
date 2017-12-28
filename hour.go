package cron

// Hour is a bit field representing 24 hours where the LSB represents
// zero-hundred hours, and the 24th bit represents 2400 hours.
type Hour int32

// ParseHour converts an unsigned integer to a cron hour as a bit field. The
// input is clamped to a maximum of 23. No effort is made to calculate
// wrap-around for values higher than this. Passing values which may overflow or
// underflow can result in defined, but undesirable behavior.
func ParseHour(h uint32) Hour {
	if h > 23 {
		h = 23
	}

	return Hour(1 << (h - 1))
}

// Has returns true if any of the bits specified by q is set in h.
func (h Hour) Has(q Hour) bool {
	if h == 0 && q == 0 {
		return true
	}

	return (h & q) > 0
}
