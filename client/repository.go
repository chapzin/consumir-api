package client

import "fmt"

const (
	repositoresBasePath      = "repositories"
	repositoryBasePath       = "repos"
	searchRepositoryBasePath = "search/repository"
)

type Repository struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Starts      int    `json:"stargazers_count"`
	Watchers    int    `json:"subscribers_count"`
}

type RepositoryService interface {
	List(*ListOptions) ([]Repository, *Response, error)
	Get(owner, repo string) (*Repository, *Response, error)
	Seach(*SearchOptions) (*SearchResponse, *Response, error)
}

type RepositoryServiceOp struct {
	client *Client
}

type ListOptions struct {
	Since int `url:"since"`
}

type SearchOptions struct {
	Query string `url:"q"`
	Sort  string `url:"sort"`
	Order string `url:"order"`
}

type SearchResponse struct {
	Total             int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

func (s *RepositoryServiceOp) List(options *ListOptions) ([]Repository, *Response, error) {
	path, err := addOptions(repositoresBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	repos := new([]Repository)
	resp, err := s.client.Do(req, repos)
	if err != nil {
		return nil, resp, err
	}

	return *repos, resp, err

}

func (s *RepositoryServiceOp) Get(owner, repoName string) (*Repository, *Response, error) {
	path := fmt.Sprintf("%s/%s/%s", repositoryBasePath, owner, repoName)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	repo := new(Repository)
	resp, err := s.client.Do(req, repo)
	if err != nil {
		return nil, resp, err
	}

	return repo, resp, err
}

func (s *RepositoryServiceOp) Search(opt *SearchOptions) (*SearchResponse, *Response, error) {
	path, err := addOptions(searchRepositoryBasePath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	searchResponse := new(SearchResponse)
	resp, err := s.client.Do(req, searchResponse)
	if err != nil {
		return nil, resp, err
	}

	return searchResponse, resp, err

}
