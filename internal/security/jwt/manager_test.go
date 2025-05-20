package jwt_test

import (
	"testing"
	"time"

	"final/internal/security/jwt"
	"github.com/stretchr/testify/require"
)

func TestGenerateAndDecodeAccessToken(t *testing.T) {
	// Инициализируем JWT
	jwt.Init("test-secret", time.Hour)

	// Тестовые данные
	userID := "123"
	role := "user"

	// Генерация токена
	token, err := jwt.GenerateAccessToken(userID, role)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Декодируем токен
	claims, err := jwt.DecodeAccessToken(token)
	require.NoError(t, err)
	require.Equal(t, userID, claims.UserID)
	require.Equal(t, role, claims.Role)
}
func TestDecodeAccessToken_InvalidToken(t *testing.T) {
	jwt.Init("test-secret", time.Hour) // на всякий случай

	badToken := "this.is.not.valid.jwt"

	_, err := jwt.DecodeAccessToken(badToken)
	require.Error(t, err)
}
func TestDecodeAccessToken_ExpiredToken(t *testing.T) {
	// Устанавливаем очень короткий срок жизни токена
	jwt.Init("test-secret", time.Millisecond*10)

	token, err := jwt.GenerateAccessToken("456", "user")
	require.NoError(t, err)

	// Ждём, чтобы токен успел истечь
	time.Sleep(time.Millisecond * 20)

	_, err = jwt.DecodeAccessToken(token)
	require.Error(t, err)
	require.Contains(t, err.Error(), "token is expired")
}
