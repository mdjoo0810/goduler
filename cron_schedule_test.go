package goduler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCronSchedule_CronSchedule(t *testing.T) {
	cases := []struct {
		name        string
		given       string
		expected    *CronSchedule
		expectError error
	}{
		{
			name:  "[Ok] all",
			given: "* * * * * *",
			expected: &CronSchedule{
				exp:      "* * * * * *",
				timeZone: time.UTC,
			},
		},
		{
			name:        "[Error] invalid cron expression",
			given:       "",
			expectError: errors.New(ErrorMsgInvalidCronExpression),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			schedule, err := NewCronScheduleFactory().CronSchedule(tc.given)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, schedule)
			}
		})
	}
}

func TestCronSchedule_DailyAtHourAndMinute(t *testing.T) {
	cases := []struct {
		name        string
		givenHour   int
		givenMinute int
		expected    *CronSchedule
		expectError error
	}{
		{
			name:        "[Ok] 1hour 1 minute",
			givenHour:   1,
			givenMinute: 1,
			expected: &CronSchedule{
				exp:      "0 1 1 ? * *",
				timeZone: time.UTC,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			schedule, err := NewCronScheduleFactory().DailyAtHourAndMinute(tc.givenHour, tc.givenMinute)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, schedule)
			}
		})
	}
}

func TestCronSchedule_AtHourAndMinuteOnGivenDaysOfWeek(t *testing.T) {
	cases := []struct {
		name            string
		givenHour       int
		givenMinute     int
		givenDaysOfWeek []int
		expected        *CronSchedule
		expectError     error
	}{
		{
			name:            "[Ok] 1hour 1 minute 1 daysOfWeek",
			givenHour:       1,
			givenMinute:     1,
			givenDaysOfWeek: []int{1},
			expected: &CronSchedule{
				exp:      "0 1 1 ? * 1",
				timeZone: time.UTC,
			},
		},
		{
			name:            "[Ok] 1hour 1 minute 1,2,3 daysOfWeek",
			givenHour:       1,
			givenMinute:     1,
			givenDaysOfWeek: []int{1, 2, 3},
			expected: &CronSchedule{
				exp:      "0 1 1 ? * 1,2,3",
				timeZone: time.UTC,
			},
		},
		{
			name:            "[Error] 1hour 1 minute empty daysOfWeek",
			givenHour:       1,
			givenMinute:     1,
			givenDaysOfWeek: []int{},
			expectError:     errors.New(ErrorMsgEmptyDaysOfWeek),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			schedule, err := NewCronScheduleFactory().AtHourAndMinuteOnGivenDaysOfWeek(tc.givenHour, tc.givenMinute, tc.givenDaysOfWeek...)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, schedule)
			}
		})
	}
}

func TestCronSchedule_WeeklyOnDayAndHourAndMinute(t *testing.T) {
	cases := []struct {
		name           string
		givenHour      int
		givenMinute    int
		givenDayOfWeek int
		expected       *CronSchedule
		expectError    error
	}{
		{
			name:           "[Ok] 1hour 1 minute 1 dayOfWeek",
			givenHour:      1,
			givenMinute:    1,
			givenDayOfWeek: 1,
			expected: &CronSchedule{
				exp:      "0 1 1 ? * 1",
				timeZone: time.UTC,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			schedule, err := NewCronScheduleFactory().WeeklyOnDayAndHourAndMinute(tc.givenDayOfWeek, tc.givenHour, tc.givenMinute)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, schedule)
			}
		})
	}
}

func TestCronSchedule_MonthlyOnDayAndHourAndMinute(t *testing.T) {
	cases := []struct {
		name            string
		givenHour       int
		givenMinute     int
		givenDayOfMonth int
		expected        *CronSchedule
		expectError     error
	}{
		{
			name:            "[Ok] 1hour 1 minute 1 dayOfMonth",
			givenHour:       1,
			givenMinute:     1,
			givenDayOfMonth: 1,
			expected: &CronSchedule{
				exp:      "0 1 1 1 * ?",
				timeZone: time.UTC,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			schedule, err := NewCronScheduleFactory().MonthlyOnDayAndHourAndMinute(tc.givenDayOfMonth, tc.givenHour, tc.givenMinute)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, schedule)
			}
		})
	}
}

func TestCronSchedule_InTimeZone(t *testing.T) {
	cases := []struct {
		name        string
		given       *CronSchedule
		givenLocStc string
		expectError error
	}{
		{
			name: "[Ok] in kst",
			given: &CronSchedule{
				exp:      "* * * * * *",
				timeZone: time.UTC,
			},
			givenLocStc: "Asia/Seoul",
		},
		{
			name: "[Error] invalid loc string",
			given: &CronSchedule{
				exp:      "* * * * * *",
				timeZone: time.UTC,
			},
			givenLocStc: "Asia/Japan",
			expectError: errors.New("unknown time zone Asia/Japan"),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.given.InTimeZone(tc.givenLocStc)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.Equal(t, tc.givenLocStc, tc.given.timeZone.String())
			}
		})
	}
}
