package services

import (
	"context"
	"iosync/ent"
	"iosync/internal/repositories"
)

type SessionService struct {
	sessionRepository *repositories.SessionRepository
}

func NewSessionService(dbClient *ent.Client) *SessionService {
	sessionRepository := repositories.NewSessionRepository(dbClient)

	return &SessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, username string) (string, error) {
	session, err := s.sessionRepository.CreateSession(ctx, username)
	if err != nil {
		return "", err
	}

	return session.SessionID, nil
}
