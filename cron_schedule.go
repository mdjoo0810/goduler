package goduler

import (
	"errors"
	"fmt"
	"github.com/adhocore/gronx"
	"time"
)

type CronSchedule struct {
	exp      string
	timeZone *time.Location
}

func (cs *CronSchedule) InTimeZone(locStr string) error {
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		return err
	}
	cs.timeZone = loc
	return nil
}

type cronSchedule struct{}

type CronScheduleFactory interface {
	CronSchedule(string) (*CronSchedule, error)
	DailyAtHourAndMinute(int, int) (*CronSchedule, error)
	AtHourAndMinuteOnGivenDaysOfWeek(int, int, ...int) (*CronSchedule, error)
	WeeklyOnDayAndHourAndMinute(int, int, int) (*CronSchedule, error)
	MonthlyOnDayAndHourAndMinute(int, int, int) (*CronSchedule, error)
}

func (cs *cronSchedule) CronSchedule(expStr string) (*CronSchedule, error) {
	if !validateExp(expStr) {
		return nil, errors.New(ErrorMsgInvalidCronExpression)
	}

	return &CronSchedule{
		exp:      expStr,
		timeZone: time.UTC,
	}, nil
}

func (cs *cronSchedule) DailyAtHourAndMinute(hour int, minute int) (*CronSchedule, error) {
	exp := fmt.Sprintf("0 %d %d ? * *", minute, hour)
	return cs.CronSchedule(exp)
}

func (cs *cronSchedule) AtHourAndMinuteOnGivenDaysOfWeek(hour int, minute int, daysOfWeek ...int) (*CronSchedule, error) {
	if len(daysOfWeek) == 0 {
		return nil, errors.New(ErrorMsgEmptyDaysOfWeek)
	}
	exp := fmt.Sprintf("0 %d %d ? * %d", minute, hour, daysOfWeek[0])
	for i := 1; i < len(daysOfWeek); i++ {
		exp = fmt.Sprintf("%s,%d", exp, daysOfWeek[i])
	}
	return cs.CronSchedule(exp)
}

func (cs *cronSchedule) WeeklyOnDayAndHourAndMinute(dayOfWeek int, hour int, minute int) (*CronSchedule, error) {
	exp := fmt.Sprintf("0 %d %d ? * %d", minute, hour, dayOfWeek)
	return cs.CronSchedule(exp)
}

func (cs *cronSchedule) MonthlyOnDayAndHourAndMinute(dayOfMonth int, hour int, minute int) (*CronSchedule, error) {
	exp := fmt.Sprintf("0 %d %d %d * ?", minute, hour, dayOfMonth)
	return cs.CronSchedule(exp)
}

func validateExp(exp string) bool {
	g := gronx.New()
	return g.IsValid(exp)
}

func NewCronScheduleFactory() CronScheduleFactory {
	return &cronSchedule{}
}
