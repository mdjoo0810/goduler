package goduler

import (
	"fmt"
	"testing"
)

func TestJob_Execute(t *testing.T) {
	cases := []struct {
		name            string
		givenKey        string
		givenDesc       string
		givenJobDataMap map[string]string
		executeFunc     func(map[string]string)
	}{
		{
			name:      "[Ok] print 1",
			givenKey:  "print_1",
			givenDesc: "print 1",
			givenJobDataMap: map[string]string{
				"num": "1",
			},
			executeFunc: func(dataMap map[string]string) {
				if n, ok := dataMap["num"]; ok {
					fmt.Println(n)
				}
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			j := NewJob(tc.givenKey, tc.givenDesc, tc.givenJobDataMap, tc.executeFunc)
			j.Execute()
		})
	}
}
