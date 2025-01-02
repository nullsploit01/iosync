package repository

import (
	"context"

	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/ent/node"
	"github.com/nullsploit01/iosync/ent/nodeapikey"
	"github.com/nullsploit01/iosync/ent/nodevalues"
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
	return n.db.Node.Query().
		WithAPIKeys(func(nakq *ent.NodeApiKeyQuery) {
			nakq.Where(nodeapikey.IsRevokedEQ(false))
		}).All(ctx)
}

func (n NodeRepository) GetNode(ctx context.Context, id int) (*ent.Node, error) {
	return n.db.Node.Query().Where(node.IDEQ(id)).
		WithAPIKeys(func(nakq *ent.NodeApiKeyQuery) {
			nakq.Where(nodeapikey.IsRevokedEQ(false))
		}).Only(ctx)
}

func (n NodeRepository) GetNodeValues(ctx context.Context, id int) ([]*ent.NodeValues, error) {
	return n.db.NodeValues.Query().Where(nodevalues.HasNodeWith(node.IDEQ(id))).All(ctx)
}

func (n NodeRepository) CreateNode(ctx context.Context, name, description string) (*ent.Node, error) {
	return n.db.Node.Create().
		SetName(name).
		SetDescription(description).
		Save(ctx)
}

func (n NodeRepository) AddNodeValue(ctx context.Context, node *ent.Node, value string) (*ent.NodeValues, error) {
	return n.db.NodeValues.Create().
		SetNode(node).
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

func (n NodeRepository) GetNodeByAPIKey(ctx context.Context, apiKey string) (*ent.Node, error) {
	return n.db.Node.Query().Where(node.HasAPIKeysWith(nodeapikey.APIKey(apiKey))).Only(ctx)
}
