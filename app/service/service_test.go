package service

import (
	"testing"
	"time"

	"github.com/himanshu-sah/goss/app/utils"

	"github.com/stretchr/testify/require"
)

func createRandomSession(session string, t *testing.T) string {
	sessionId, err := CreateSession(session, time.Second*5)
	require.NoError(t, err)
	require.NotEmpty(t, sessionId)
	return sessionId
}
func TestCreateSession(t *testing.T) {
	createRandomSession(utils.RandomSession(), t)
}
func TestGetSession(t *testing.T) {
	session := utils.RandomSession()
	sessionId := createRandomSession(session, t)

	storedSession, err := GetSession(sessionId)
	require.NoError(t, err)
	require.NotEmpty(t, storedSession)
	require.Equal(t, session, storedSession)
}
func TestDeleteSession(t *testing.T) {
	session := utils.RandomSession()
	sessionId := createRandomSession(session, t)

	// reading the session
	storedSession, err := GetSession(sessionId)
	require.NoError(t, err)
	require.NotEmpty(t, storedSession)
	require.Equal(t, session, storedSession)

	// deleting the session
	err = DeleteSession(sessionId)
	require.NoError(t, err)

	// reading the session for error
	storedSession, err = GetSession(sessionId)
	require.Error(t, err)
	require.Empty(t, storedSession)
}
func TestTruncateStore(t *testing.T) {
	sessionIds := [5]string{}
	for i := 1; i < 5; i++ {
		session := utils.RandomSession()
		sessionId := createRandomSession(session, t)
		sessionIds[i] = sessionId

		// reading the session
		storedSession, err := GetSession(sessionId)
		require.NoError(t, err)
		require.NotEmpty(t, storedSession)
		require.Equal(t, session, storedSession)
	}

	err := TruncateStore()
	require.NoError(t, err)

	// reading the session for error
	for i := 1; i < 5; i++ {
		storedSession, err := GetSession(sessionIds[i])
		require.Error(t, err)
		require.Empty(t, storedSession)
	}
}
