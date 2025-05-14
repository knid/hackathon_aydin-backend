package views

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/knid/timezen/internal/requests"
	"github.com/knid/timezen/internal/utils"
)

func (vw *Views) GetProjects(w http.ResponseWriter, r *http.Request) {
	user, err := vw.Controller.GetUserFromRequest(r)
	if err != nil {
		utils.JSONResponse(w, 401, err)
		return 
	}

	projects := vw.Controller.GetProjectsByUser(user)

	utils.JSONResponse(w, 200, projects)
}


func (vw *Views) CreateProject(w http.ResponseWriter, r *http.Request) {
	user, err := vw.Controller.GetUserFromRequest(r)
	if err != nil {
		utils.JSONResponse(w, 401, err)
		return 
	}

	var projectCreateRequest requests.ProjectCreateRequest
	json.NewDecoder(r.Body).Decode(&projectCreateRequest)

	project := vw.Controller.CreateProject(user, projectCreateRequest.Title, projectCreateRequest.Description)

	utils.JSONResponse(w, 201, project)
}


func (vw *Views) GetProject(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId") 
	project, err := vw.Controller.GetProject(projectId)

	if err != nil {
		utils.JSONResponse(w, 400, err)
		return
	}

	utils.JSONResponse(w, 200, project)
}

func (vw *Views) GetProjectTasks(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId") 
	tasks, err := vw.Controller.GetProjectTasks(projectId)
	if err != nil {
		utils.JSONResponse(w, 400, err)
		return
	}

	utils.JSONResponse(w, 200, tasks)

}
