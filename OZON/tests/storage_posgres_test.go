package tests
import (
	"os"
	"testing"

	"github.com/crowyy03/OZON/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestPostgresStorage(t *testing.T) {
	dsn := os.Getenv("TEST_DSN")
	if dsn == "" {
		t.Skip("TEST_DSN not set, skipping Postgres tests")
	}

	store, err := storage.NewPostgresStorage(dsn)
	assert.NoError(t, err)

	shortURL := "test123"
	originalURL := "https://example.com"

	// Проверяем сохранение
	err = store.Save(shortURL, originalURL)
	assert.NoError(t, err)

	// Проверяем получение
	result, err := store.Get(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, originalURL, result)
}
