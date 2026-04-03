package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/repositories"
)

type JobService struct {
	jobRepo *repositories.JobRespository
}

type JobsListResponse struct {
	Jobs  []generated.Job
	Count int64
	Page  int32
	Limit int32
}

type GetJobsInput struct {
	UserID pgtype.UUID
	Page   int32
	Limit  int32
}

func NewJobSerive(jobRepos *repositories.JobRespository) *JobService {
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

func (jobSrv *JobService) CreateJob(
	ctx context.Context,
	arg generated.CreateJobParams,
) (generated.Job, error) {
	if arg.Status == "" {
		arg.Status = StatusApplied
	}
	return jobSrv.jobRepo.CreateJob(ctx, arg)
}

func (jobSrv *JobService) GetJobs(
	ctx context.Context,
	arg GetJobsInput,
) (JobsListResponse, error) {
	const MaxLimit = 50
	if arg.Page < 1 {
		arg.Page = 1
	}
	if arg.Limit <= 0 {
		arg.Limit = 10 // default
	}
	if arg.Limit > MaxLimit {
		arg.Limit = MaxLimit
	}
	jobsList, err := jobSrv.jobRepo.GetJobsByUser(ctx, generated.GetJobsParams{
		UserID: arg.UserID,
		Limit:  arg.Limit,
		Offset: (arg.Page - 1) * arg.Limit,
	})
	if err != nil {
		return JobsListResponse{}, err
	}
	totalCount, err := jobSrv.jobRepo.GetJobsCount(ctx, arg.UserID)
	if err != nil {
		return JobsListResponse{}, err
	}

	return JobsListResponse{Jobs: jobsList, Count: totalCount, Page: arg.Page, Limit: arg.Limit}, nil
}

func (jobSrv *JobService) GetJob(
	ctx context.Context,
	arg generated.GetJobByIDParams,
) (generated.Job, error) {
	return jobSrv.jobRepo.GetJobById(ctx, arg)
}

func (jobSrv *JobService) UpdateJobStatus(
	ctx context.Context,
	arg generated.UpdateJobStatusParams,
) (generated.Job, error) {
	oldJob, err := jobSrv.jobRepo.GetJobById(
		ctx,
		generated.GetJobByIDParams{ID: arg.ID, UserID: arg.UserID},
	)
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
		jobSrv.jobRepo.InsertStatusHistory(
			ctx,
			generated.InsertJobStatusHistoryParams{
				JobID:     arg.ID,
				OldStatus: oldStatus,
				NewStatus: newStatus,
			},
		)
	}

	return newData, nil
}

func (jobSrv *JobService) GetJobApplicationStatuses(
	ctx context.Context,
) ([]generated.JobApplicationStatuse, error) {
	statuses, err := jobSrv.jobRepo.GetJobApplicationStatuses(ctx)
	if err != nil {
		return []generated.JobApplicationStatuse{}, errors.New("failed to get statuses")
	}
	return statuses, nil
}

func (s *JobService) GetJobStatusHistory(
	ctx context.Context,
	userId pgtype.UUID, jobId pgtype.UUID,
) ([]generated.JobStatusHistory, error) {
	_, err := s.jobRepo.GetJobById(ctx, generated.GetJobByIDParams{UserID: userId, ID: jobId})
	if err != nil {
		return nil, err
	}

	statusHistory, err := s.jobRepo.GetStatusHistory(ctx, jobId)
	if err != nil {
		return nil, errors.New("something went wrong while fetching")
	}
	return statusHistory, nil
}
