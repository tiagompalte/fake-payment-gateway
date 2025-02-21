package job

import (
	"github.com/tiagompalte/fake-payment-gateway/application"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
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

func (j jobGroup) GetJob(jobName string) (protocols.Job, bool) {
	job, ok := j.job[jobName]
	return job, ok
}
