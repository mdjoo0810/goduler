package goduler

import (
	"errors"
	"time"
)

type SimpleSchedule struct {
	interval    time.Duration
	repeatCount int
}

type SimpleScheduleBuilder interface {
	RepeatSecondlyForever(int) SimpleScheduleBuilder
	RepeatMinutelyForever(int) SimpleScheduleBuilder
	RepeatHourlyForever(int) SimpleScheduleBuilder
	RepeatSecondlyForTotalCount(int, int) SimpleScheduleBuilder
	RepeatMinutelyForTotalCount(int, int) SimpleScheduleBuilder
	RepeatHourlyForTotalCount(int, int) SimpleScheduleBuilder
	Build() (*SimpleSchedule, error)
}

type simpleSchedule struct {
	interval    time.Duration
	repeatCount int
}

func (ss *simpleSchedule) RepeatSecondlyForever(interval int) SimpleScheduleBuilder {
	ss.interval = time.Duration(interval) * time.Second
	ss.repeatCount = RepeatIndefinitely
	return ss
}

func (ss *simpleSchedule) RepeatMinutelyForever(interval int) SimpleScheduleBuilder {
	ss.interval = time.Duration(interval) * time.Minute
	ss.repeatCount = RepeatIndefinitely
	return ss
}

func (ss *simpleSchedule) RepeatHourlyForever(interval int) SimpleScheduleBuilder {
	ss.interval = time.Duration(interval) * time.Hour
	ss.repeatCount = RepeatIndefinitely
	return ss
}

func (ss *simpleSchedule) RepeatSecondlyForTotalCount(interval int, count int) SimpleScheduleBuilder {
	ss.interval = time.Duration(interval) * time.Second
	ss.repeatCount = count
	if ss.repeatCount == RepeatIndefinitely {
		ss.repeatCount--
	}
	return ss
}

func (ss *simpleSchedule) RepeatMinutelyForTotalCount(interval int, count int) SimpleScheduleBuilder {
	ss.interval = time.Duration(interval) * time.Minute
	ss.repeatCount = count
	if ss.repeatCount == RepeatIndefinitely {
		ss.repeatCount--
	}
	return ss
}

func (ss *simpleSchedule) RepeatHourlyForTotalCount(interval int, count int) SimpleScheduleBuilder {
	ss.interval = time.Duration(interval) * time.Hour
	ss.repeatCount = count
	if ss.repeatCount == RepeatIndefinitely {
		ss.repeatCount--
	}
	return ss
}

func (ss *simpleSchedule) Build() (*SimpleSchedule, error) {
	if ss.interval <= 0 {
		return nil, errors.New(ErrorMsgInvalidIntervalGreaterThanZero)
	}

	if ss.repeatCount != RepeatIndefinitely && ss.repeatCount <= 0 {
		return nil, errors.New(ErrorMsgInvalidCountGreaterThanZero)
	}

	return &SimpleSchedule{
		interval:    ss.interval,
		repeatCount: ss.repeatCount,
	}, nil
}

func NewSimpleScheduleBuilder() SimpleScheduleBuilder {
	return &simpleSchedule{}
}
