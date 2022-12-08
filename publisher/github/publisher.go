package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ezeoleaf/larry/config"
	"github.com/ezeoleaf/larry/domain"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

type client interface {
	GetContents(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error)
	UpdateFile(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
}

type repositoryData struct {
	Owner    string
	Name     string
	FileName string
}

// Publisher represents the publisher client
type Publisher struct {
	GithubClient   client
	Config         config.Config
	RepositoryData repositoryData
}

// NewPublisher returns a new publisher client
func NewPublisher(apiKey string, cfg config.Config, repoOwner, repoName, repoFileName string) Publisher {
	log.Print("New Github Publisher")
	p := Publisher{Config: cfg, RepositoryData: repositoryData{Owner: repoOwner, Name: repoName, FileName: repoFileName}}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	p.GithubClient = github.NewClient(tc).Repositories

	return p
}

func (p Publisher) getReadmeContent(ctx context.Context) (*github.RepositoryContent, error) {

	fc, _, _, err := p.GithubClient.GetContents(ctx, p.RepositoryData.Owner, p.RepositoryData.Name, p.RepositoryData.FileName, nil)

	if err != nil {
		log.Print(fmt.Errorf("could not fetch repository, got %v", err))
		return nil, err
	}

	if fc == nil || fc.Content == nil {
		log.Print("content of README is empty")
		return nil, errors.New("no content")
	}

	return fc, nil
}

func decodeBase64(c string) (string, error) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return "", err
	}

	return string(rawDecodedText), nil
}

// PublishContent receives a content to publish and try to publish in README file
func (p Publisher) PublishContent(content *domain.Content) (bool, error) {
	if content == nil || (content.Title == nil || content.URL == nil) {
		return false, errors.New("no content to publish")
	}

	contentToAdd := fmt.Sprintf("[%s](%s)", *content.Title, *content.URL)
	if content.Subtitle != nil {
		contentToAdd += ": " + *content.Subtitle
	}

	if p.Config.SafeMode {
		log.Print("Running in Safe Mode")
		log.Print(contentToAdd)
		return true, nil
	}
	ctx := context.Background()

	repositoryContent, err := p.getReadmeContent(ctx)

	if err != nil {
		return false, err
	}

	readmeContent, err := decodeBase64(*repositoryContent.Content)
	if err != nil {
		return false, err
	}

	if strings.Contains(readmeContent, *content.Title) {
		log.Printf("repository %s already exists", *content.Title)
		return false, nil
	}

	readmeContent += fmt.Sprintf("  \n%s", contentToAdd)

	m := "Adding new repo"

	fo := github.RepositoryContentFileOptions{
		Content: []byte(readmeContent),
		SHA:     repositoryContent.SHA,
		Message: &m,
	}

	_, _, err = p.GithubClient.UpdateFile(ctx, p.RepositoryData.Owner, p.RepositoryData.Name, p.RepositoryData.FileName, &fo)

	if err != nil {
		return false, err
	}

	return true, nil
}
