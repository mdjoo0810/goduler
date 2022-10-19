package goduler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"time"
)

func TestSimpleSchedule_RepeatSecondlyForever(t *testing.T) {
	cases := []struct {
		name        string
		given       int
		expected    *SimpleSchedule
		expectError error
	}{
		{
			name:  "[Ok] given 1 second",
			given: 1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Second,
				repeatCount: RepeatIndefinitely,
			},
		},
		{
			name:        "[Error] given -1 second",
			given:       -1,
			expectError: errors.New(ErrorMsgInvalidIntervalGreaterThanZero),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSimpleScheduleFactory().RepeatSecondlyForever(tc.given)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s)
			}
		})
	}
}

func TestSimpleSchedule_RepeatMinutelyForever(t *testing.T) {
	cases := []struct {
		name        string
		given       int
		expected    *SimpleSchedule
		expectError error
	}{
		{
			name:  "[Ok] given 1 minute",
			given: 1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Minute,
				repeatCount: RepeatIndefinitely,
			},
		},
		{
			name:        "[Error] given -1 minute",
			given:       -1,
			expectError: errors.New(ErrorMsgInvalidIntervalGreaterThanZero),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSimpleScheduleFactory().RepeatMinutelyForever(tc.given)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s)
			}
		})
	}
}

func TestSimpleSchedule_RepeatHourlyForever(t *testing.T) {
	cases := []struct {
		name        string
		given       int
		expected    *SimpleSchedule
		expectError error
	}{
		{
			name:  "[Ok] given 1 Hour",
			given: 1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Hour,
				repeatCount: RepeatIndefinitely,
			},
		},
		{
			name:        "[Error] given -1 Hour",
			given:       -1,
			expectError: errors.New(ErrorMsgInvalidIntervalGreaterThanZero),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSimpleScheduleFactory().RepeatHourlyForever(tc.given)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s)
			}
		})
	}
}

func TestSimpleSchedule_RepeatSecondlyForTotalCount(t *testing.T) {
	cases := []struct {
		name          string
		givenInterval int
		givenCount    int
		expected      *SimpleSchedule
		expectError   error
	}{
		{
			name:          "[Ok] given 1 second and 1 count",
			givenInterval: 1,
			givenCount:    1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Second,
				repeatCount: 1,
			},
		},
		{
			name:          "[Error] given -1 second",
			givenInterval: -1,
			givenCount:    1,
			expectError:   errors.New(ErrorMsgInvalidIntervalGreaterThanZero),
		},
		{
			name:          "[Error] given 0 count",
			givenInterval: 1,
			givenCount:    0,
			expectError:   errors.New(ErrorMsgInvalidCountGreaterThanZero),
		},
		{
			name:          "[Error] given -1 count",
			givenInterval: 1,
			givenCount:    -1,
			expectError:   errors.New(ErrorMsgInvalidCountGreaterThanZero),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSimpleScheduleFactory().RepeatSecondlyForTotalCount(tc.givenInterval, tc.givenCount)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s)
			}
		})
	}
}

func TestSimpleSchedule_RepeatMinutelyForTotalCount(t *testing.T) {
	cases := []struct {
		name          string
		givenInterval int
		givenCount    int
		expected      *SimpleSchedule
		expectError   error
	}{
		{
			name:          "[Ok] given 1 minute and 1 count",
			givenInterval: 1,
			givenCount:    1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Minute,
				repeatCount: 1,
			},
		},
		{
			name:          "[Error] given -1 minute",
			givenInterval: -1,
			givenCount:    1,
			expectError:   errors.New(ErrorMsgInvalidIntervalGreaterThanZero),
		},
		{
			name:          "[Error] given 0 count",
			givenInterval: 1,
			givenCount:    0,
			expectError:   errors.New(ErrorMsgInvalidCountGreaterThanZero),
		},
		{
			name:          "[Error] given -1 count",
			givenInterval: 1,
			givenCount:    -1,
			expectError:   errors.New(ErrorMsgInvalidCountGreaterThanZero),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSimpleScheduleFactory().RepeatMinutelyForTotalCount(tc.givenInterval, tc.givenCount)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s)
			}
		})
	}
}

func TestSimpleSchedule_RepeatHourlyForTotalCount(t *testing.T) {
	cases := []struct {
		name          string
		givenInterval int
		givenCount    int
		expected      *SimpleSchedule
		expectError   error
	}{
		{
			name:          "[Ok] given 1 hour and 1 count",
			givenInterval: 1,
			givenCount:    1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Hour,
				repeatCount: 1,
			},
		},
		{
			name:          "[Error] given -1 hour",
			givenInterval: -1,
			givenCount:    1,
			expectError:   errors.New(ErrorMsgInvalidIntervalGreaterThanZero),
		},
		{
			name:          "[Error] given 0 count",
			givenInterval: 1,
			givenCount:    0,
			expectError:   errors.New(ErrorMsgInvalidCountGreaterThanZero),
		},
		{
			name:          "[Error] given -1 count",
			givenInterval: 1,
			givenCount:    -1,
			expectError:   errors.New(ErrorMsgInvalidCountGreaterThanZero),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSimpleScheduleFactory().RepeatHourlyForTotalCount(tc.givenInterval, tc.givenCount)
			if tc.expectError != nil {
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s)
			}
		})
	}
}
