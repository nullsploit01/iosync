package service

import (
	"context"

	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/internal/repository"
)

type NodeService struct {
	repo repository.NodeRepository
}

type CreateNodeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func NewNodeService(db *ent.Client) NodeService {
	return NodeService{
		repo: repository.NewNodeRepository(db),
	}
}

func (n NodeService) GetNodes(ctx context.Context) ([]*ent.Node, error) {
	return n.repo.GetNodes(ctx)
}

func (n NodeService) GetNode(ctx context.Context, id int) (*ent.Node, error) {
	return n.repo.GetNode(ctx, id)
}

func (n NodeService) CreateNode(ctx context.Context, request CreateNodeRequest) (*ent.Node, error) {
	return n.repo.CreateNode(ctx, request.Name, request.Description)
}
