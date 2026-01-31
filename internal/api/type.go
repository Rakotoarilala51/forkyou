package api

type ForkResponse struct {
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

type ReadmeResponse struct {
	Content string `json:"content"`
}

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

type SearchResponse struct {
	Results []*SearchResult `json:"items"`
}
type SearchResult struct {
	FullName string `json:"full_name"`
}
