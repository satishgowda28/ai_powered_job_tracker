package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/utils"
)

type JobHandler struct {
	jobService *services.JobService
}

func NewJobHandler(jobService *services.JobService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

func (h *JobHandler) CreateJob(c *fiber.Ctx) error {
	// UserID         pgtype.UUID
	userId, err := utils.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "user id missing from context",
		})
	}
	var body struct {
		Company        string `json:"company"`
		Title          string `json:"title"`
		JobDescription string `json:"job_description"`
		JobLocation    string `json:"job_location"`
		Salary         int32  `json:"salary"`
		JobUrl         string `json:"job_url"`
		Status         string `json:"status"`
		Notes          string `json:"notes"`
	}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad-request",
			"message": "invalid request body",
		})
	}
	newJob, err := h.jobService.CreateJob(c.Context(), generated.CreateJobParams{
		UserID:         utils.ToPgtypeUUID(userId),
		Company:        body.Company,
		Title:          body.Title,
		JobDescription: body.JobDescription,
		JobLocation:    body.JobLocation,
		Salary:         pgtype.Int4{Valid: true, Int32: body.Salary},
		JobUrl:         body.JobUrl,
		Status:         body.Status,
		Notes:          body.Notes,
	})
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "save-failed",
			"message": "Something went work creating new job",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": newJob})
}

func (h *JobHandler) GetJobs(c *fiber.Ctx) error {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "user id missing from context",
		})
	}
	jobs, err := h.jobService.GetJobs(
		c.Context(),
		generated.GetJobsParams{UserID: utils.ToPgtypeUUID(userId)},
	)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "failed-getting-jobs",
			"message": "Something went work",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": jobs})
}

func (h *JobHandler) GetJob(c *fiber.Ctx) error {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "user id missing from context",
		})
	}
	jobIdStr := c.Params("id")
	jobId, err := uuid.Parse(jobIdStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server_error",
			"message": "user id missing from context",
		})
	}
	jobs, err := h.jobService.GetJob(
		c.Context(),
		generated.GetJobByIDParams{
			UserID: utils.ToPgtypeUUID(userId),
			ID:     utils.ToPgtypeUUID(jobId),
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "failed-getting-jobs",
			"message": "Something went work",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": jobs})
}

func (h *JobHandler) UpdateJobStatus(c *fiber.Ctx) error {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "user id missing from context",
		})
	}
	jobIDStr := c.Params("id")
	jobId, err := uuid.Parse(jobIDStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "server_error",
			"message": "job id issue",
		})
	}
	var body struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(body); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "bad-request",
			"message": "invalid request body",
		})
	}
	job, err := h.jobService.UpdateJobStatus(c.Context(),
		generated.UpdateJobStatusParams{
			ID:     utils.ToPgtypeUUID(jobId),
			Status: body.Status,
			UserID: utils.ToPgtypeUUID(userId),
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "job-update-fail",
			"message": "upadating status failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": job,
	})
}
