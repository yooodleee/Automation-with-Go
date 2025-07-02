package request_form 


type WorkflowRequest struct {
	RepoUrl string `json:"repo_uri"`
	Token 	string `json:"token"`
}