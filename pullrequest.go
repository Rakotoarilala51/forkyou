package forkyou

import "github.com/spf13/cobra"

type PullRequestPayload struct{
	Title string `json:"title"`
	Message string `json:"message"`
	SourceBranch string `json:"head"`
	DestBranch string `json:"base"`
	Modify string `json:"maintener_can_modify"`
}
type PullRequestResponse struct{
	URL string `json:"html_url"`
}
var PullRequestCmd =&cobra.Command{
	Use: "pullrequest",
	Short: "Create a Pull Request",
	Run: func(cmd *cobra.Command, args []string) {

	},
}