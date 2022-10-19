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
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * time.Second,
		repeatCount: RepeatIndefinitely,
	}, nil
}

func (ss *simpleSchedule) RepeatMinutelyForever(interval int) (*SimpleSchedule, error) {
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * time.Minute,
		repeatCount: RepeatIndefinitely,
	}, nil
}

func (ss *simpleSchedule) RepeatHourlyForever(interval int) (*SimpleSchedule, error) {
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * time.Hour,
		repeatCount: RepeatIndefinitely,
	}, nil
}

func (ss *simpleSchedule) RepeatSecondlyForTotalCount(interval int, count int) (*SimpleSchedule, error) {
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	if count <= 0 {
		return nil, errors.New(ErrorMsgInvalidCountGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * time.Second,
		repeatCount: count,
	}, nil
}

func (ss *simpleSchedule) RepeatMinutelyForTotalCount(interval int, count int) (*SimpleSchedule, error) {
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	if count <= 0 {
		return nil, errors.New(ErrorMsgInvalidCountGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * time.Minute,
		repeatCount: count,
	}, nil
}

func (ss *simpleSchedule) RepeatHourlyForTotalCount(interval int, count int) (*SimpleSchedule, error) {
	if interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}
	if count <= 0 {
		return nil, errors.New(ErrorMsgInvalidCountGreaterThanZero)
	}
	return &SimpleSchedule{
		interval:    time.Duration(interval) * time.Hour,
		repeatCount: count,
	}, nil
}

func NewSimpleScheduleFactory() SimpleScheduleFactory {
	return &simpleSchedule{}
}
