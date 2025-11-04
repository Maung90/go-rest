package parser

import (
	"errors"
	"time"
)

func ParseDateString(dateString string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, errors.New("format tanggal tidak valid, gunakan YYYY-MM-DD")
	}
	return parsedTime, nil
}