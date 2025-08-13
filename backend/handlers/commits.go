package handlers

import(
	"encoding/json"
	"net/http"

	"golang-git-analyzor-plus-kobweb/backend/gitutils"
)

func GetCommits(w http.ResponseWriter,r *http.Request){
	repoPath := "./repo-test"

	commits,err:= gitutils.GetCommits(repoPath)

	if err!=nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","aaplication/json")
	json.NewEncoder(w).Encode(commits)
}