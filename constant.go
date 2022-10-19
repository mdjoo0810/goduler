package goduler

const (
	RepeatIndefinitely = -1

	ErrorMsgInvalidIntervalGreaterThanZero = "invalid interval, interval is greater than zero"
	ErrorMsgInvalidCountGreaterThanZero    = "invalid count, count is greater than zero"
	ErrorMsgInvalidCronExpression          = "invalid cron expression"
	ErrorMsgEmptyDaysOfWeek                = "must specify at least one day of week"
)
