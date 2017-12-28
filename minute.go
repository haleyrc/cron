package cron

// Minute is a bit field representing 60 minutes where the LSB represents the
// top of the hour and the 59th bit represents the minute before the top of the
// hour.
type Minute int64

// ParseMinute converts an unsigned integer to a cron minute as a bit field. The
// input is clamped to a maximum of 59. No effort is made to calculate
// wrap-around for values higher than this. Passing values which may overflow or
// underflow can result in defined, but undesirable behavior.
func ParseMinute(m uint64) Minute {
	if m > 59 {
		m = 59
	}

	return Minute(1 << (m - 1))
}

// Has returns true if any of the bits specified by q is set in m.
func (m Minute) Has(q Minute) bool {
	if m == 0 && q == 0 {
		return true
	}

	return (m & q) > 0
}
