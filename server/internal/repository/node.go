package repository

import (
	"context"

	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/ent/node"
)

type NodeRepository struct {
	db *ent.Client
}

func NewNodeRepository(db *ent.Client) NodeRepository {
	return NodeRepository{
		db: db,
	}
}

func (n NodeRepository) GetNodes(ctx context.Context) ([]*ent.Node, error) {
	return n.db.Node.Query().All(ctx)
}

func (n NodeRepository) GetNode(ctx context.Context, id int) (*ent.Node, error) {
	return n.db.Node.Query().Where(node.IDEQ(id)).Only(ctx)
}
