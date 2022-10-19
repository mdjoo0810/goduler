package goduler

import (
	"errors"
	"time"
)

type SimpleSchedule struct {
	interval    time.Duration
	repeatCount int
}

type SimpleScheduleFactory interface {
	RepeatSecondlyForever(int) (*SimpleSchedule, error)
	RepeatMinutelyForever(int) (*SimpleSchedule, error)
	RepeatHourlyForever(int) (*SimpleSchedule, error)
	RepeatSecondlyForTotalCount(int, int) (*SimpleSchedule, error)
	RepeatMinutelyForTotalCount(int, int) (*SimpleSchedule, error)
	RepeatHourlyForTotalCount(int, int) (*SimpleSchedule, error)
}

type simpleSchedule struct{}

func (ss *simpleSchedule) RepeatSecondlyForever(interval int) (*SimpleSchedule, error) {
	return makeSimpleSchedule(interval, RepeatIndefinitely, true, time.Second)
}

func (ss *simpleSchedule) RepeatMinutelyForever(interval int) (*SimpleSchedule, error) {
	return makeSimpleSchedule(interval, RepeatIndefinitely, true, time.Minute)
}

func (ss *simpleSchedule) RepeatHourlyForever(interval int) (*SimpleSchedule, error) {
	return makeSimpleSchedule(interval, RepeatIndefinitely, true, time.Hour)
}

func (ss *simpleSchedule) RepeatSecondlyForTotalCount(interval int, count int) (*SimpleSchedule, error) {
	return makeSimpleSchedule(interval, count, false, time.Second)
}

func (ss *simpleSchedule) RepeatMinutelyForTotalCount(interval int, count int) (*SimpleSchedule, error) {
	return makeSimpleSchedule(interval, count, false, time.Minute)
}

func (ss *simpleSchedule) RepeatHourlyForTotalCount(interval int, count int) (*SimpleSchedule, error) {
	return makeSimpleSchedule(interval, count, false, time.Hour)
}

func makeSimpleSchedule(interval int, count int, isForever bool, unit time.Duration) (*SimpleSchedule, error) {
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	if !isForever && count <= 0 {
		return nil, errors.New(ErrorMsgInvalidCountGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * unit,
		repeatCount: count,
	}, nil
}

func NewSimpleScheduleFactory() SimpleScheduleFactory {
	return &simpleSchedule{}
}
