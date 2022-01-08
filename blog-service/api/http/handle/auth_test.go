package handle

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"

    "blog-service/api/http/dto"
)

/**
 * @author Rancho
 * @date 2022/1/7
 */

func TestGetAuth(t *testing.T) {
    var w = httptest.NewRecorder()
    var response map[string]interface{}
    body := &dto.AuthRequest{
        AppKey:    "rancho",
        AppSecret: "blog-service",
    }
    b, err := json.Marshal(body)
    require.NoError(t, err)
    req, _ := http.NewRequest(http.MethodGet, "/auth", bytes.NewReader(b))

    NewServerRoute().ServeHTTP(w, req)

    // verify
    assert.Equal(t, http.StatusOK, w.Code)
    err = json.Unmarshal(w.Body.Bytes(), &response)
    fmt.Println(response)
    assert.Nil(t, err)
    assert.NotEmpty(t, response["token"])
}
