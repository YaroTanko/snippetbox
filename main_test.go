package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HomeReturnsHelloWorld(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	home(rr, r)

	result := rr.Result()
	body, _ := io.ReadAll(result.Body)

	if string(body) != "Hello World!" {
		t.Errorf("expected 'Hello World!', got '%v'", string(body))
	}
}

func SnippetCreateReturnsMessage(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/snippet/create", nil)

	snippetCreate(rr, r)

	result := rr.Result()
	body, _ := io.ReadAll(result.Body)

	if string(body) != "Create a new snippet..." {
		t.Errorf("expected 'Create a new snippet...', got '%v'", string(body))
	}
}

func SnippetViewWithValidIDReturnsSnippet(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/snippet/view?id=1", nil)

	snippetview(rr, r)

	result := rr.Result()
	body, _ := io.ReadAll(result.Body)

	if string(body) != "Display a specific snippet with ID 1..." {
		t.Errorf("expected 'Display a specific snippet with ID 1...', got '%v'", string(body))
	}
}

func SnippetViewWithInvalidIDReturns404(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/snippet/view?id=0", nil)

	snippetview(rr, r)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %v", rr.Code)
	}
}

func SnippetViewWithNonNumericIDReturns404(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/snippet/view?id=abc", nil)

	snippetview(rr, r)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %v", rr.Code)
	}
}
