package repository

import "github.com/nullsploit01/iosync/ent"

type NodeRepository struct {
	db *ent.Client
}

func NewNodeRepository(db *ent.Client) NodeRepository {
	return NodeRepository{
		db: db,
	}
}
