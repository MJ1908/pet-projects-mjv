package Services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	config "petgitlab/configs"
	"petgitlab/model"
	"time"
)

func GetAllProjects() ([]model.Project, error) {

	response, err := http.Get(config.GitlabApiPath)
	if err != nil {
		return nil, fmt.Errorf("HTTP error: %s", err)
	} else {
		responseData, _ := ioutil.ReadAll(response.Body)

		var projects []model.Project
		json.Unmarshal(responseData, &projects)
		return projects, nil
	}
}

func IsInTimeWindow(target, startDate, endDate time.Time) bool {
	return !target.Before(startDate) && target.Before(endDate)
}
