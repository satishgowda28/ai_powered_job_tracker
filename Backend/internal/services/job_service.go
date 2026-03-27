package services

import (
	"context"
	"errors"

	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/respositories"
)

type JobService struct {
	jobRepo *respositories.JobRespository
}

func NewJobSerive(jobRepos *respositories.JobRespository) *JobService {
	return &JobService{
		jobRepo: jobRepos,
	}
}

const (
	StatusApplied   = "applied"
	StatusInterview = "interview"
	StatusOffer     = "offer"
	StatusRejected  = "rejected"
)

func (jobSrv *JobService) CreateJob(ctx context.Context, arg generated.CreateJobParams) (generated.Job, error) {
	if arg.Status == "" {
		arg.Status = StatusApplied
	}
	return jobSrv.jobRepo.CreateJob(ctx, arg)
}

func (jobSrv *JobService) GetJobs(ctx context.Context, arg generated.GetJobsParams) ([]generated.Job, error) {
	return jobSrv.jobRepo.GetJobsByUser(ctx, arg)
}

func (jobSrv *JobService) GetJob(ctx context.Context, arg generated.GetJobByIDParams) (generated.Job, error) {
	return jobSrv.jobRepo.GetJobById(ctx, arg)
}

func (jobSrv *JobService) UpdateJobStatus(ctx context.Context, arg generated.UpdateJobStatusParams) (generated.Job, error) {
	oldJob, err := jobSrv.jobRepo.GetJobById(ctx, generated.GetJobByIDParams{ID: arg.ID, UserID: arg.UserID})
	if err != nil {
		return generated.Job{}, errors.New("no data found")
	}
	oldStatus := oldJob.Status
	newStatus := arg.Status
	newData, err := jobSrv.jobRepo.UpdateJobStatus(ctx, arg)
	if err != nil {
		return generated.Job{}, errors.New("something went wrong while updating")
	}
	if oldStatus != newStatus {
		jobSrv.jobRepo.InsertStatusHistory(ctx, generated.InsertJobStatusHistoryParams{JobID: arg.ID, OldStatus: oldStatus, NewStatus: newStatus})
	}

	return newData, nil
}

func (jobSrv *JobService) GetJobApplicationStatuses(ctx context.Context) ([]generated.JobApplicationStatuse, error) {
	statuses, err := jobSrv.jobRepo.GetJobApplicationStatuses(ctx)
	if err != nil {
		return []generated.JobApplicationStatuse{}, errors.New("failed to get statuses")
	}
	return statuses, nil
}
