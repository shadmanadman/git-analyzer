package handlers

import(
	"encoding/json"
	"net/http"

	"backend/gitutils"
)

type Contributor struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Count int `json:"count"`
}

func GetContributers(w http.ResponseWriter,r *http.Request){
	repoPath := "./test-repo"

	commits,err:=gitutils.GetCommits(repoPath)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	contributors:=map[string]Contributor{}

	for _, c := range commits {
		key := c.Email
		if _, exists := contributors[key]; !exists {
			contributors[key] = Contributor{Name: c.Author, Email: c.Email, Count: 0}
		}
		temp := contributors[key]
		temp.Count++
		contributors[key] = temp
	}

	var result []Contributor
	for _,v := range contributors{
		result = append(result, v)
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(result)
}