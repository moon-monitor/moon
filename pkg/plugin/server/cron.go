package server

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/moon-monitor/moon/pkg/util/safety"
	"github.com/robfig/cron/v3"
)

var _ transport.Server = (*CronJobServer)(nil)

type CronSpec string

const (
	CronSpecYearly CronSpec = "@yearly"

	CronSpecAnnually CronSpec = "@annually"

	CronSpecMonthly CronSpec = "@monthly"

	CronSpecWeekly CronSpec = "@weekly"

	CronSpecDaily CronSpec = "@daily"

	CronSpecMidnight CronSpec = "@midnight"

	CronSpecHourly CronSpec = "@hourly"
)

func Every(duration time.Duration) CronSpec {
	return CronSpec("@every " + duration.String())
}

func Custom(s, m, h, d, M, y string) CronSpec {
	return CronSpec(s + " " + m + " " + h + " " + d + " " + M + " " + y)
}

type CronJob interface {
	cron.Job

	ID() cron.EntryID
	Index() string
	Sepc() CronSpec
	WithID(id cron.EntryID) CronJob
}

var defaultCronParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

type CronJobServer struct {
	cron *cron.Cron

	tasks *safety.Map[string, CronJob]
}

func NewCronJobServer() *CronJobServer {
	c := &CronJobServer{
		cron:  cron.New(cron.WithParser(defaultCronParser)),
		tasks: safety.NewMap[string, CronJob](),
	}
	return c
}

func (c *CronJobServer) AddJob(job CronJob) (cron.EntryID, error) {
	if _, ok := c.tasks.Get(job.Index()); ok {
		return 0, errors.New("job already exists")
	}
	id, err := c.cron.AddJob(string(job.Sepc()), job)
	if err != nil {
		return 0, err
	}
	job.WithID(id)
	c.tasks.Set(job.Index(), job)
	return id, nil
}

func (c *CronJobServer) Remove(job CronJob) {
	c.cron.Remove(job.ID())
	c.tasks.Delete(job.Index())
}

func (c *CronJobServer) Start(ctx context.Context) error {
	c.cron.Start()
	return nil
}

func (c *CronJobServer) Stop(ctx context.Context) error {
	c.cron.Stop()
	c.tasks.Clear()
	return nil
}
