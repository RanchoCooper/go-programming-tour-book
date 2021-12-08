package router

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/magiconair/properties/assert"

    "go-programming-tour-book/blog-service/api/http/DTO"
)

/**
 * @author Rancho
 * @date 2021/12/8
 */

func performRequest(r http.Handler, method, path string, body io.Reader, contentType string) *httptest.ResponseRecorder {
    req, _ := http.NewRequest(method, path, body)
    if contentType == "" {
        req.Header.Set("Content-Type", "application/json")
    } else {
        req.Header.Set("Content-Type", contentType)
    }
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    return w
}

func TestGetAuth(t *testing.T) {
    body := DTO.AuthRequest{
        AppKey:    "rancho",
        AppSecret: "go-programming-tour-book",
    }
    b, _ := json.Marshal(body)
    w := performRequest(NewRouter(), http.MethodPost, "/auth", bytes.NewReader(b), "multipart/form-data")
    assert.Equal(t, http.StatusOK, w.Code)
}
