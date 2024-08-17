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

func (s *SessionRepository) Create(ctx context.Context, user *ent.User, timeout time.Time) (*ent.Session, error) {
	sessionId := utils.GenerateUuid()
	return s.dbClient.Session.Create().
		SetSessionID(sessionId).
		SetExpiresAt(timeout).
		SetUser(user).
		Save(ctx)
}

func (s *SessionRepository) GetBySessionId(ctx context.Context, sessionId string) (*ent.Session, error) {
	return s.dbClient.Session.Query().
		Where(session.SessionID(sessionId), session.ExpiresAtGTE(time.Now())).
		WithUser().
		First(ctx)
}

func (s *SessionRepository) Update(ctx context.Context, sessionId string, newExpiryDate time.Time) error {
	_, err := s.dbClient.Session.Update().
		Where(session.SessionID(sessionId), session.ExpiresAtGTE(time.Now())).
		SetExpiresAt(newExpiryDate).
		Save(ctx)

	return err
}

func (s *SessionRepository) Delete(ctx context.Context, sessionId string) error {
	_, err := s.dbClient.Session.Delete().
		Where(session.SessionID(sessionId)).
		Exec(ctx)

	return err
}
