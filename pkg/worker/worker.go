package worker

import (
	"context"

	"github.com/go-redis/redis/v9"
	"go.temporal.io/sdk/workflow"

	"github.com/instill-ai/model-backend/pkg/acl"
	"github.com/instill-ai/model-backend/pkg/ray"
	"github.com/instill-ai/model-backend/pkg/repository"
)

// TaskQueue is the Temporal task queue name for model-backend
const TaskQueue = "model-backend"

// Worker interface
type Worker interface {
	TriggerModelWorkflow(ctx workflow.Context, param *TriggerModelWorkflowRequest) (*TriggerModelWorkflowResponse, error)
	TriggerModelActivity(ctx context.Context, param *TriggerModelActivityRequest) (*TriggerModelActivityResponse, error)
}

// worker represents resources required to run Temporal workflow and activity
type worker struct {
	redisClient *redis.Client
	repository  repository.Repository
	ray         ray.Ray
	aclClient   *acl.ACLClient
}

// NewWorker initiates a temporal worker for workflow and activity definition
func NewWorker(r repository.Repository, rc *redis.Client, ra ray.Ray, a *acl.ACLClient) Worker {

	return &worker{
		repository:  r,
		redisClient: rc,
		ray:         ra,
		aclClient:   a,
	}
}
