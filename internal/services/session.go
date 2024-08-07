package services

import (
	"context"
	"errors"
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
	activeSession, err := s.sessionRepository.GetUserActiveSession(ctx, username)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if !errors.As(err, &notFoundError) {
			return "", err
		}
	} else if activeSession != nil {
		return "", errors.New("user is already logged in")
	}

	session, err := s.sessionRepository.CreateSession(ctx, username)
	if err != nil {
		return "", errors.New("error creating session")
	}

	return session.SessionID, nil
}
