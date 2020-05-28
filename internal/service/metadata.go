package service

import (
	"bytes"
	"encoding/json"

	"github.com/docker/docker/utils/templates"
	"gitlab.unanet.io/devops/eve/pkg/eve"
)

type TemplateServiceData struct {
	Plan    *eve.NSDeploymentPlan
	Service *eve.DeployService
}

type TemplateMigrationData struct {
	Plan      *eve.NSDeploymentPlan
	Migration *eve.DeployMigration
}

func ParseServiceMetadata(metadata map[string]interface{}, service *eve.DeployService, plan *eve.NSDeploymentPlan) (map[string]interface{}, error) {
	metadataJson, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	temp, err := templates.Parse(string(metadataJson))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = temp.Execute(&b, TemplateServiceData{
		Plan:    plan,
		Service: service,
	})
	if err != nil {
		return nil, err
	}

	var returnMap map[string]interface{}
	err = json.Unmarshal(b.Bytes(), &returnMap)
	if err != nil {
		return nil, err
	}

	return returnMap, nil
}

func ParseMigrationMetadata(metadata map[string]interface{}, migration *eve.DeployMigration, plan *eve.NSDeploymentPlan) (map[string]interface{}, error) {
	metadataJson, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	temp, err := templates.Parse(string(metadataJson))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = temp.Execute(&b, TemplateMigrationData{
		Plan:      plan,
		Migration: migration,
	})
	if err != nil {
		return nil, err
	}

	var returnMap map[string]interface{}
	err = json.Unmarshal(b.Bytes(), &returnMap)
	if err != nil {
		return nil, err
	}

	return returnMap, nil
}
