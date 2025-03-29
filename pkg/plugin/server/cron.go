package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/robfig/cron/v3"

	"github.com/moon-monitor/moon/pkg/util/safety"
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

func CronSpecEvery(duration time.Duration) CronSpec {
	return CronSpec("@every " + duration.String())
}

func CronSpecCustom(s, m, h, d, M, y string) CronSpec {
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
	cron  *cron.Cron
	tasks *safety.Map[string, CronJob]
	help  *log.Helper
}

func NewCronJobServer(logger log.Logger, jobs ...CronJob) *CronJobServer {
	c := &CronJobServer{
		cron:  cron.New(cron.WithParser(defaultCronParser)),
		tasks: safety.NewMap[string, CronJob](),
		help:  log.NewHelper(log.With(logger, "module", "server.cron")),
	}
	for _, job := range jobs {
		c.AddJob(job)
	}
	return c
}

func (c *CronJobServer) AddJob(job CronJob) {
	if _, ok := c.tasks.Get(job.Index()); ok {
		return
	}
	id, err := c.cron.AddJob(string(job.Sepc()), job)
	if err != nil {
		c.help.Warnw("method", "add job", "err", err)
		return
	}
	job.WithID(id)
	c.tasks.Set(job.Index(), job)
}

func (c *CronJobServer) Remove(job CronJob) {
	c.cron.Remove(job.ID())
	c.tasks.Delete(job.Index())
}

func (c *CronJobServer) Start(_ context.Context) error {
	c.cron.Start()
	return nil
}

func (c *CronJobServer) Stop(_ context.Context) error {
	c.cron.Stop()
	c.tasks.Clear()
	return nil
}
