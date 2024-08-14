package services

import (
	"context"
	"errors"
	"iosync/ent"
	"iosync/internal/repositories"
	"time"
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

func (s *SessionService) CreateSession(ctx context.Context, username string, timeout time.Time) (*ent.Session, error) {
	activeSession, err := s.sessionRepository.GetUserActiveSession(ctx, username)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if !errors.As(err, &notFoundError) {
			return nil, err
		}
	} else if activeSession != nil {
		return nil, errors.New("user is already logged in")
	}

	session, err := s.sessionRepository.CreateSession(ctx, username, timeout)
	if err != nil {
		return nil, errors.New("error creating session")
	}

	return session, nil
}

func (s *SessionService) VerifySession(ctx context.Context, sessionId string) (*ent.Session, error) {
	session, err := s.sessionRepository.GetSessionBySessionId(ctx, sessionId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			return nil, errors.New("invalid session id")
		}
		return nil, err
	} else if session == nil {
		return nil, errors.New("invalid session id")
	}

	return session, nil
}

func (s *SessionService) GetUserActiveSession(ctx context.Context, username string) (*ent.Session, error) {
	session, err := s.sessionRepository.GetUserActiveSession(ctx, username)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			return nil, errors.New("user has no active session")
		}

		return nil, err
	}

	return session, nil
}

func (s *SessionService) RefreshSessionExpiry(ctx context.Context, session *ent.Session) error {
	remainingTime := time.Until(session.ExpiresAt)

	if remainingTime > 30*time.Minute {
		return nil
	}

	err := s.sessionRepository.UpdateSessionExpiryDate(ctx, session.SessionID, session.ExpiresAt.Add(30*time.Minute))
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			return errors.New("invalid session id")
		}
	}

	return err
}
