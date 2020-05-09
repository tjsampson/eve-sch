package service

import (
	"context"

	"gitlab.unanet.io/devops/eve/pkg/errors"
	"gitlab.unanet.io/devops/eve/pkg/eve"
	"gitlab.unanet.io/devops/eve/pkg/queue"
	"go.uber.org/zap"
)

const (
	CommandUpdateDeployment string = "api-update-deployment"
	GroupUpdateDeployment   string = "api-update-deployment"
)

func (s *Scheduler) handleMessage(ctx context.Context, m *queue.M) error {
	switch m.Command {
	case CommandDeployNamespace:
		return s.deployNamespace(ctx, m)
	default:
		return errors.Wrapf("unrecognized command: %s", m.Command)
	}
}

func (s *Scheduler) deployNamespace(ctx context.Context, m *queue.M) error {
	plan, err := eve.UnMarshalNSDeploymentFromS3LocationBody(ctx, s.downloader, m.Body)
	if err != nil {
		return errors.Wrap(err)
	}

	for _, x := range plan.Services {
		var vaultPaths []string
		x.Metadata, vaultPaths, err = ParseServiceMetadata(x.Metadata, x, plan)
		if err != nil {
			s.Logger(ctx).Error("unable to parse metadata", zap.Error(err))
		}
		if x.ArtifactoryFeedType == eve.ArtifactoryFeedTypeDocker {
			s.deployDockerService(ctx, x, plan, vaultPaths)
		}
		if len(x.ArtifactFnPtr) > 0 {
			s.triggerFunction(ctx, x.DeployArtifact, plan, vaultPaths)
		}
	}

	for _, x := range plan.Migrations {
		var vaultPaths []string
		x.Metadata, vaultPaths, err = ParseMigrationMetadata(x.Metadata, x, plan)
		if err != nil {
			s.Logger(ctx).Error("unable to parse metadata", zap.Error(err))
		}

		if x.ArtifactoryFeedType == eve.ArtifactoryFeedTypeDocker {
			s.runDockerMigrationJob(ctx, x, plan, vaultPaths)
		}
		if len(x.ArtifactFnPtr) > 0 {
			s.triggerFunction(ctx, x.DeployArtifact, plan, vaultPaths)
		}
	}

	err = s.worker.DeleteMessage(ctx, m)
	if err != nil {
		return errors.Wrap(err)
	}

	if plan.Failed() {
		plan.Status = eve.DeploymentPlanStatusErrors
	} else {
		plan.Status = eve.DeploymentPlanStatusComplete
	}

	mBody, err := eve.MarshalNSDeploymentPlanToS3LocationBody(ctx, s.uploader, plan)
	if err != nil {
		return errors.Wrap(err)
	}

	err = s.worker.Message(ctx, s.apiQUrl, &queue.M{
		ID:      m.ID,
		ReqID:   queue.GetReqID(ctx),
		GroupID: GroupUpdateDeployment,
		Body:    mBody,
		Command: CommandUpdateDeployment,
	})
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
