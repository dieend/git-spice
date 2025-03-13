package github

import (
	"net/http"

	"github.com/dieend/git-spice/internal/graphqlutil"
	"github.com/shurcooL/githubv4"
)

func newGitHubEnterpriseClient(
	url string,
	httpClient *http.Client,
) *githubv4.Client {
	httpClient.Transport = graphqlutil.WrapTransport(httpClient.Transport)
	return githubv4.NewEnterpriseClient(url, httpClient)
}
