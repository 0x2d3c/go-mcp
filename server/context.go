package server

import (
	"context"
	"errors"
)

type sessionIDKey struct{}

func setSessionIDToCtx(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, sessionID)
}

func getSessionIDFromCtx(ctx context.Context) (string, error) {
	sessionID := ctx.Value(sessionIDKey{})
	if sessionID == nil {
		return "", errors.New("no session id found")
	}
	return sessionID.(string), nil
}

type sendChanKey struct{}

func setSendChanToCtx(ctx context.Context, sendCh chan<- []byte) context.Context {
	return context.WithValue(ctx, sendChanKey{}, sendCh)
}

func getSendChanFromCtx(ctx context.Context) (chan<- []byte, error) {
	ch := ctx.Value(sendChanKey{})
	if ch == nil {
		return nil, errors.New("no send chan found")
	}
	return ch.(chan<- []byte), nil
}
