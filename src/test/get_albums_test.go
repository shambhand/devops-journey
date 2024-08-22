package test

import (
	"encoding/json"
	"github.com/devops-journey/src/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	router := app.SetUpRouter()

	t.Run("should get the list of albums", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/albums", nil)

		router.ServeHTTP(response, request)

		expectedAlbums := []app.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		}

		var actualAlbums []string
		json.Unmarshal([]byte(response.Body.String()), &actualAlbums)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedAlbums, actualAlbums)
	})
}
