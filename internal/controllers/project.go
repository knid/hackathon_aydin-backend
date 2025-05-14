package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/knid/timezen/internal/models"
	"github.com/knid/timezen/pkg/llm/clients"
	"github.com/ollama/ollama/api"
)

func (c *Controller) GetProjectsByUser(user models.User) []models.Project {
	var projects []models.Project

	c.DB.Joins("JOIN project_users ON project_users.project_id = projects.id").
   		Where("project_users.user_id = ?", user.ID).
   		Find(&projects)

	return projects
}

func (c *Controller) CreateProject(user models.User, name, description string) models.Project {
	project := models.Project{
		Name: name,
		Description: description,
		Users: []models.User{user},
		ActiveJob: false,
	}

	c.DB.Create(&project)

	// Generate tasks
	go func() {
		client := clients.OllamaCLient{
			Addr: "http://10.8.0.10:11434",
			Model: "timezen:1",
		}

		project.ActiveJob = true
		c.DB.Save(&project)

		var lineBuffer string
		ctx := context.Background()

		msg := "{\"project_description\":\"" + description + "\"}"
		fmt.Println(msg)
		client.SendChat(ctx, msg, func(resp api.GenerateResponse) error {
			if resp.Response == "\n" {
				fmt.Print(lineBuffer)
				var task models.Task
	            if err := json.Unmarshal([]byte(lineBuffer), &task); err != nil {
	                fmt.Println("Failed to parse line:", lineBuffer)
	                fmt.Println("Error:", err)
	            }

				task.Project = project
				c.DB.Create(&task)

				lineBuffer = ""
			} else {
				lineBuffer += resp.Response
			}
			return nil
		})
		project.ActiveJob = false
		c.DB.Save(&project)
    }()

	return project
}

func (c *Controller) GetProject(projectId string) (models.Project, error) {
	var project models.Project
	result := c.DB.Model(&models.Project{}).Preload("Users").Where("id = ?", projectId).First(&project)
	if result.Error != nil {
		return models.Project{}, result.Error
	}

	return project, nil
}


func (c *Controller) GetProjectTasks(projectId string) ([]models.Task, error) {	
	var project models.Project
	result := c.DB.Model(&models.Project{}).Preload("Users").Where("id = ?", projectId).First(&project)
	if result.Error != nil {
		return []models.Task{}, result.Error
	}

	var tasks []models.Task
	result = c.DB.Where("project_id = ?", projectId).Find(&tasks)
	if result.Error != nil {
		return []models.Task{}, result.Error
	}

	return tasks, nil

}


