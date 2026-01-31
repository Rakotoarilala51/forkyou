package forkyou

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	destRepo           string
	sourceRepo         string
	pullRequestTitle   string
	pullRequestMessage string
	SourceBranch       string
	destBranch         string
)

type PullRequestPayload struct {
	Title        string `json:"title"`
	Message      string `json:"message"`
	SourceBranch string `json:"head"`
	DestBranch   string `json:"base"`
	Modify       string `json:"maintener_can_modify"`
}
type PullRequestResponse struct {
	URL string `json:"html_url"`
}

var PullRequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Short: "Create a Pull Request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreatePullRequest(); err != nil {
			log.Fatalln("Failed to create pull request: ", err)
		}
	},
}

func CreatePullRequest() error {
	sourcesValues := strings.Split(sourceRepo, ":")
	if !(len(sourcesValues) == 1 || len(sourcesValues) == 2) {
		return fmt.Errorf("source repository must in the format [owner:]branch got  %v", sourceRepo)
	}
	destBranchValues := strings.Split(destRepo, ":")
	if len(destBranchValues) != 2 {
		return fmt.Errorf("destination repository must be in format in the format owner/project:branch got %v", destRepo)
	}
	destValues := strings.Split(destBranchValues[0], "/")
	if len(destValues) != 2 {
		return fmt.Errorf("destination repository must be in the format owner/project:branch got %v", destRepo)
	}
	return nil
}

func init() {
	PullRequestCmd.Flags().StringVarP(&sourceRepo, "source", "s", "", "source repository")
	PullRequestCmd.Flags().StringVarP(&destRepo, "destination", "d", "", "destination repository")
	PullRequestCmd.Flags().StringVarP(&pullRequestTitle, "title", "t", "Basic pull request", "pull request title")
	PullRequestCmd.Flags().StringVarP(&pullRequestMessage, "message", "m", "", "pull request message")
}
