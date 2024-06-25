package cmd

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	log = logrus.New()
}

func TestGetTargetBit(t *testing.T) {
	cases := map[string]struct {
		weekday time.Weekday
		expect  int
	}{
		"monday": {
			weekday: time.Monday,
			expect:  0,
		},
		"tuesday": {
			weekday: time.Tuesday,
			expect:  1,
		},
		"wednesday": {
			weekday: time.Wednesday,
			expect:  2,
		},
		"thursday": {
			weekday: time.Thursday,
			expect:  3,
		},
		"friday": {
			weekday: time.Friday,
			expect:  4,
		},
		"saturday": {
			weekday: time.Saturday,
			expect:  5,
		},
		"sunday": {
			weekday: time.Sunday,
			expect:  6,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := getTargetBit(tc.weekday)
			if actual != tc.expect {
				t.Errorf("expect: %d, actual: %d", tc.expect, actual)
			}
		})
	}

}

func TestIsWeekday(t *testing.T) {

	cases := map[string]struct {
		number  int
		weekday time.Weekday
		expect  bool
	}{
		"all1: monday": {
			number:  127,
			weekday: time.Monday,
			expect:  true,
		},
		"all1: tuesday": {
			number:  127,
			weekday: time.Tuesday,
			expect:  true,
		},
		"all1: wednesday": {
			number:  127,
			weekday: time.Wednesday,
			expect:  true,
		},
		"all1: thursday": {
			number:  127,
			weekday: time.Thursday,
			expect:  true,
		},
		"all1: friday": {
			number:  127,
			weekday: time.Friday,
			expect:  true,
		},
		"all1: satday": {
			number:  127,
			weekday: time.Saturday,
			expect:  true,
		},
		"all1: sunday": {
			number:  127,
			weekday: time.Sunday,
			expect:  true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := isWeekday(tc.number, tc.weekday)
			if actual != tc.expect {
				t.Errorf("expect: %t, actual: %t", tc.expect, actual)
			}
		})
	}
}
