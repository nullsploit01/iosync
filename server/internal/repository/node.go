package repository

import (
	"context"

	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/ent/node"
	"github.com/nullsploit01/iosync/internal/util"
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

func (n NodeRepository) CreateNode(ctx context.Context, name, description string) (*ent.Node, error) {
	return n.db.Node.Create().
		SetName(name).
		SetDescription(description).
		Save(ctx)
}

func (n NodeRepository) AddNodeValue(ctx context.Context, nodeId int, value string) (*ent.NodeValues, error) {
	return n.db.NodeValues.Create().
		SetNodeID(nodeId).
		SetValue(value).
		Save(ctx)
}

func (n NodeRepository) GenerateNodeAPIKey(ctx context.Context, nodeId int, apiKeyDescription string) (*ent.NodeApiKey, error) {
	return n.db.NodeApiKey.Create().
		SetNodeID(nodeId).
		SetAPIKey(util.GenerateUUID()).
		SetDescription(apiKeyDescription).
		Save(ctx)
}
