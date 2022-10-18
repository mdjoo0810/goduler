package goduler

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"time"
)

func Test_repeatSecondlyForever(t *testing.T) {
	cases := []struct {
		name     string
		given    int
		expected *SimpleSchedule
	}{
		{
			name:  " given 1 second",
			given: 1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Second,
				repeatCount: RepeatIndefinitely,
			},
		},
		{
			name:  "given 10 seconds",
			given: 10,
			expected: &SimpleSchedule{
				interval:    time.Duration(10) * time.Second,
				repeatCount: RepeatIndefinitely,
			},
		},
		{
			name:  "given -1 second but return 1 second",
			given: -1,
			expected: &SimpleSchedule{
				interval:    time.Duration(1) * time.Second,
				repeatCount: RepeatIndefinitely,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s := NewSimpleScheduleBuilder().RepeatSecondlyForever(tc.given).Build()
			assert.Equal(t, tc.expected, s)
		})
	}
}
