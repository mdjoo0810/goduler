package goduler

type Job struct {
	jobKey      string
	description string
	jobDataMap  map[string]string
	executeFunc func(params map[string]string)
}

func (j *Job) Execute() {
	j.executeFunc(j.jobDataMap)
}

func NewJob(key, desc string, dataMap map[string]string, exeFunc func(params map[string]string)) *Job {
	return &Job{
		jobKey:      key,
		description: desc,
		jobDataMap:  dataMap,
		executeFunc: exeFunc,
	}
}
