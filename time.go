package utils

import "time"

func SecondsInDay() (s int) {
	tm := time.Now()
	return tm.Second() + tm.Minute()*60 + tm.Hour()*60*60
}
