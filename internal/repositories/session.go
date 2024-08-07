package repositories

import (
	"context"
	"iosync/ent"
	"iosync/pkg/utils"
	"time"
)

type SessionRepository struct {
	dbClient *ent.Client
}

func NewSessionRepository(dbClient *ent.Client) *SessionRepository {
	return &SessionRepository{dbClient: dbClient}
}

func (s *SessionRepository) CreateSession(ctx context.Context, username string) (*ent.Session, error) {
	sessionId := utils.GenerateUuid()

	return s.dbClient.Session.Create().
		SetSessionID(sessionId).
		SetUsername(username).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetExpiresAt(time.Now().Add(time.Minute * 30)).
		Save(ctx)
}
