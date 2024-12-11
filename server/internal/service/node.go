package service

import (
	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/internal/repository"
)

type NodeService struct {
	repo repository.NodeRepository
}

func NewNodeService(db *ent.Client) NodeService {
	return NodeService{
		repo: repository.NewNodeRepository(db),
	}
}
