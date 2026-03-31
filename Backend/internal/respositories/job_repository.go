package respositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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

func (jbRepo *JobRespository) CreateJob(
	ctx context.Context,
	arg generated.CreateJobParams,
) (generated.Job, error) {
	return jbRepo.q.CreateJob(ctx, arg)
}

func (jbRepo *JobRespository) GetJobsByUser(
	ctx context.Context,
	arg generated.GetJobsParams,
) ([]generated.Job, error) {
	return jbRepo.q.GetJobs(ctx, arg)
}

func (jbRepo *JobRespository) GetJobById(
	ctx context.Context,
	arg generated.GetJobByIDParams,
) (generated.Job, error) {
	return jbRepo.q.GetJobByID(ctx, arg)
}

func (jbRepo *JobRespository) UpdateJobStatus(
	ctx context.Context,
	arg generated.UpdateJobStatusParams,
) (generated.Job, error) {
	return jbRepo.q.UpdateJobStatus(ctx, arg)
}

func (jbRepo *JobRespository) InsertStatusHistory(
	ctx context.Context,
	arg generated.InsertJobStatusHistoryParams,
) error {
	return jbRepo.q.InsertJobStatusHistory(ctx, arg)
}

func (jbRepo *JobRespository) GetJobApplicationStatuses(
	ctx context.Context,
) ([]generated.JobApplicationStatuse, error) {
	return jbRepo.q.GetJobApplicationStatuses(ctx)
}

func (r *JobRespository) GetStatusHistory(
	ctx context.Context,
	jobId pgtype.UUID,
) ([]generated.JobStatusHistory, error) {

	return r.q.GetHistoryForJob(ctx, jobId)

}
