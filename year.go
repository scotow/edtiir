package edtiir

import (
	"log"
	"time"
)

var (
	locale   *time.Location
	Year2018 *Year
)

func init() {
	paris, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		log.Panicln(err)
	}

	locale = paris
	Year2018 = NewYear(
		"https://docs.google.com/spreadsheets/d/1vNAr4_zLq2RCwbtAy2y0p-RoLr4lbBbbYqtP5o0O80M/edit#gid=0",
		2018, time.September, 3,
		Config{Size{8, 6}, Padding{4, 1}, 1},
	)
}

func NewYear(link string, year int, month time.Month, day int, config Config) *Year {
	return NewYearFromTime(link, time.Date(year, month, day, 0, 0, 0, 0, locale), config)
}

func NewYearFromTime(link string, start time.Time, config Config) *Year {
	return &Year{link, start, config}
}

type Year struct {
	link   string
	start  time.Time
	config Config
}

func (y *Year) CurrentWeek() *Week {
	now := time.Now().In(locale)
	weekIndex := int(now.Sub(y.start).Truncate(time.Hour*24).Hours()) / (7 * 24)

	return &Week{y.link, weekIndex, y.config}
}
