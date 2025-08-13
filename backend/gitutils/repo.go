package gitutils

import(
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type CommitInfo struct{
	Hash string `json:"hash"`
	Auther string `json:"auther"`
	Email string `json:"email"`
	Message string `json:"message"`
	TimeStamp time.Time `json:"timestamp"`
}


func GetCommits(repoPath string) ([]CommitInfo,error){
	repo,err:=git.PlainOpen(repoPath)
	if err!=nil {
		return nil,err
	}

	ref,err:=repo.Head()
	if err!=nil {
		return nil,err
	}

	cIter,err:= repo.Log(&git.LogOptions{From: ref.Hash()})
	if err!=nil {
		return nil,err
	}

	var commits []CommitInfo
	err = cIter.ForEach(func(c *object.Commit) error {
		commits = append(commits, CommitInfo{
			Hash: c.Hash.String(),
			Auther: c.Author.Name,
			Email: c.Author.Email,
			TimeStamp: c.Author.When,
		})
		return nil
	})

	return commits,err
}