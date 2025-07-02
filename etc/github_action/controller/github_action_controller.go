package controller

import (
	"etc/github_action/controller/request_form"
	"etc/github_action/entity"
	"etc/github_action/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


// GitHubActionController 구조체
type GitHubActionController struct {
	GitHubActionService service.GitHubActionService
}

// NewGitHubActionController 생성자 함수
func NewGitHubActionController(service service.GitHubActionService) *GitHubActionController {
	return &GitHubActionController{GitHubActionService: service}
}

// GetWorkflowRuns workflow 실행 정보 가져오기
func (c *GitHubActionController) GetWorkflowRuns(ctx *fiber.Ctx) error {
	println("controller - GetWorkflowRuns()")

	// request body parsing
	var req request_form.WorkflowRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	println("controller - pass request_form")

	// check request parameter
	if req.RepoUrl == "" || req.Token == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "repo_url and token are required",
		})
	}
	println("controller - non-null request_form")

	workflowRuns, err := c.GitHubActionService.GetWorkflowRuns(req.RepoUrl, req.Token) 
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(workflowRuns)
}


// SaveWorkflowRuns - Saving workflow execution info
func (c *GitHubActionController) SaveWorkflowRuns(ctx *fiber.Ctx) error {
	var workflowRuns []entity.workflowRuns
	if err := ctx.BodyParser(&workflowRuns); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := c.GitHubActionService.SaveWorkflowRuns(workflowRuns); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(workflowRuns)
}


// GetWorkflowRunyByID - Get workflow execution info by ID
func (c *GitHubActionController) GetWorkflowRunyByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	workflowRun, err := c.GitHubActionService.GetWorkflowRunyByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "workflowRun not found"})
	}

	return ctx.JSON(workflowRun)
}


// DeleteWorkflowRun - Delete workflow execution info
func (c *GitHubActionController) DeleteWorkflowRun(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.GitHubActionService.DeleteWorkflowRun(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}