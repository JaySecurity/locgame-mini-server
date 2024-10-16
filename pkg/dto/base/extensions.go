package base

import "time"

func (x *Timestamp) ToTime() time.Time {
	if x.IsValid() {
		return time.Unix(x.Seconds, 0).UTC()
	}
	return time.Unix(0, 0)
}

func (x *Timestamp) IsValid() bool {
	return x != nil && x.Seconds > 0
}
