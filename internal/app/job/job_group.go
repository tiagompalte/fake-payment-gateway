package job

import (
	"context"

	"github.com/tiagompalte/fake-payment-gateway/application"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
)

type jobGroup struct {
	job map[string]protocols.Job
}

func NewJobGroup(app application.App) jobGroup {
	return jobGroup{
		map[string]protocols.Job{
			CreateAccountJobName: NewCreateAccountJobImpl(app.UseCase().CreateAccountUseCase()),
		},
	}
}

func (j jobGroup) getJob(jobName string) (protocols.Job, bool) {
	job, ok := j.job[jobName]
	return job, ok
}

func (j jobGroup) Execute(ctx context.Context, jobName string, args ...any) error {
	job, ok := j.getJob(jobName)
	if !ok {
		return errors.Wrap(ErrJobNotExists)
	}

	err := job.Execute(ctx, args...)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
