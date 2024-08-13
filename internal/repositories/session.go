package repositories

import (
	"context"
	"iosync/ent"
	"iosync/ent/session"
	"iosync/pkg/utils"
	"time"
)

type SessionRepository struct {
	dbClient *ent.Client
}

func NewSessionRepository(dbClient *ent.Client) *SessionRepository {
	return &SessionRepository{dbClient: dbClient}
}

func (s *SessionRepository) CreateSession(ctx context.Context, username string, timeout time.Time) (*ent.Session, error) {
	sessionId := utils.GenerateUuid()
	return s.dbClient.Session.Create().
		SetSessionID(sessionId).
		SetUsername(username).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetExpiresAt(timeout).
		Save(ctx)
}

func (s *SessionRepository) GetUserActiveSession(ctx context.Context, username string) (*ent.Session, error) {
	return s.dbClient.Session.Query().
		Where(session.ExpiresAtGTE(time.Now()), session.Username(username)).
		First(ctx)
}

func (s *SessionRepository) GetSessionBySessionId(ctx context.Context, sessionId string) (*ent.Session, error) {
	return s.dbClient.Session.Query().
		Where(session.SessionID(sessionId), session.ExpiresAtGTE(time.Now())).
		First(ctx)
}

func (s *SessionRepository) UpdateSessionExpiryDate(ctx context.Context, sessionId string, duration time.Duration) error {
	newExpiryDate := time.Now().Add(duration)

	_, err := s.dbClient.Session.Update().
		Where(session.SessionID(sessionId), session.ExpiresAtGTE(time.Now())).
		SetExpiresAt(newExpiryDate).
		Save(ctx)

	return err
}
