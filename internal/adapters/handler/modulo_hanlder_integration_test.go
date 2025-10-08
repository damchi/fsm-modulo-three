package handler

import (
	"bytes"
	"encoding/json"
	"fsm-modulo-three/internal/core/domain"
	"fsm-modulo-three/internal/fsm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	service := &fsm.ModuloService{}
	h := NewModuloHandler(service)

	r.POST("/check", h.Check)
	return r
}

func TestModuloHandler_Check(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name              string
		payload           domain.BinaryRequest
		expectedStatus    int
		expectedRemainder int
		expectError       bool
	}{
		{
			name: "mod 3, binary 1010",
			payload: domain.BinaryRequest{
				Binary: "1010",
				Mod:    3,
			},
			expectedStatus:    http.StatusOK,
			expectedRemainder: 1,
			expectError:       false,
		},
		{
			name: "mod 7, binary 1101",
			payload: domain.BinaryRequest{
				Binary: "1101",
				Mod:    7,
			},
			expectedStatus:    http.StatusOK,
			expectedRemainder: 6,
			expectError:       false,
		},
		{
			name: "invalid mod 0",
			payload: domain.BinaryRequest{
				Binary: "1010",
				Mod:    0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "empty binary",
			payload: domain.BinaryRequest{
				Binary: "",
				Mod:    3,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "very large binary input",
			payload: domain.BinaryRequest{
				Binary: strings.Repeat("1", 1000),
				Mod:    7,
			},
			expectedStatus:    http.StatusOK,
			expectError:       false,
			expectedRemainder: 1,
		},
		{
			name: "non-binary input",
			payload: domain.BinaryRequest{
				Binary: "10201",
				Mod:    3,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/check", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectError {
				var resp domain.BinaryResponse
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRemainder, resp.Remainder)
				assert.Equal(t, tt.payload.Binary, resp.Binary)
				assert.Equal(t, tt.expectedRemainder == 0, resp.IsDivisible)
			} else {
				var resp map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Contains(t, resp, "error")
			}
		})
	}
}
