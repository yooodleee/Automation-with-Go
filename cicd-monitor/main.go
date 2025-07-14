package cicdmonitor

import (
	cicdmonitor "cicd-monitor"
	"fmt"
	"os"

	"github.com/yooodleee/Automation-with-Go/cicd-monitor"
)


func main() {
	token := os.Getenv("GITHUB_TOKEN")
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPOSITORY")
	client := cicdmonitor.NewGitHubClient(token, owner, repo)

	runs, err := client.ListWorkflowRuns()
	if err != nil {
		panic(err)
	}

	for _, run := range runs {
		fmt.Printf("✅ [%s] %s (%s by %s | %s\n)",
			run.Concolusion,
			run.CommitMsg,
			run.Branch,
			run.TriggeredBy,
			run.Duration,
		)
	}
}