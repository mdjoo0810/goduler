package goduler

import "time"

type SimpleSchedule struct {
	interval    time.Duration
	repeatCount int
}

type SimpleScheduleBuilder interface {
	RepeatSecondlyForever(int) SimpleScheduleBuilder
	Build() *SimpleSchedule
}

type simpleSchedule struct {
	interval    time.Duration
	repeatCount int
}

func (ss *simpleSchedule) RepeatSecondlyForever(i int) SimpleScheduleBuilder {
	ss.interval = time.Duration(i) * time.Second
	ss.repeatCount = RepeatIndefinitely
	return ss
}

func (ss *simpleSchedule) Build() *SimpleSchedule {
	if ss.interval < 0 {
		ss.interval = ss.interval * -1
	}
	return &SimpleSchedule{
		interval:    ss.interval,
		repeatCount: ss.repeatCount,
	}
}

func NewSimpleScheduleBuilder() SimpleScheduleBuilder {
	return &simpleSchedule{}
}
