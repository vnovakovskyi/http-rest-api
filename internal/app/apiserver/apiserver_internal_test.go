package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusOK(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/merger", nil)
	s.HandleMerger().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
