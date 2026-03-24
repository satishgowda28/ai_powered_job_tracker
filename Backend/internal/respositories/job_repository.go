package respositories

import (
	"context"

	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/database"
)

type JobRespository struct {
	q *generated.Queries
}

func NewJobrepository() *JobRespository {
	return &JobRespository{
		q: generated.New(database.DB),
	}
}

func (jbRepo *JobRespository) CreateJob(ctx context.Context, arg generated.CreateJobParams) (generated.Job, error) {
	return jbRepo.q.CreateJob(ctx, arg)
}

func (jbRepo *JobRespository) GetJobsByUser(ctx context.Context, arg generated.GetJobsParams) ([]generated.Job, error) {
	return jbRepo.q.GetJobs(ctx, arg)
}

func (jbRepo *JobRespository) GetJobById(ctx context.Context, arg generated.GetJobByIDParams) (generated.Job, error) {
	return jbRepo.q.GetJobByID(ctx, arg)
}

func (jbRepo *JobRespository) UpdateJobStatus(ctx context.Context, arg generated.UpdateJobStatusParams) (generated.Job, error) {
	return jbRepo.q.UpdateJobStatus(ctx, arg)
}
