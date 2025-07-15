package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckPerfectNumbers(t *testing.T) {
	e := echo.New()

	tests := []struct {
		name         string
		requestBody  map[string]interface{}
		expectedCode int
		expectedKeys []string
	}{
		{
			name: "valid range with perfect numbers",
			requestBody: map[string]interface{}{
				"start": 1,
				"end":   10000,
			},
			expectedCode: http.StatusOK,
			expectedKeys: []string{"perfect_numbers"},
		},
		{
			name: "range with no perfect numbers",
			requestBody: map[string]interface{}{
				"start": 10001,
				"end":   20000,
			},
			expectedCode: http.StatusOK,
			expectedKeys: []string{"perfect_numbers"},
		},
		{
			name: "invalid request body missing fields",
			requestBody: map[string]interface{}{
				"start": 1,
			},
			expectedCode: http.StatusBadRequest,
			expectedKeys: []string{"error"},
		},
		{
			name: "start greater than end",
			requestBody: map[string]interface{}{
				"start": 100,
				"end":   10,
			},
			expectedCode: http.StatusBadRequest,
			expectedKeys: []string{"error"},
		},
		{
			name: "negative range",
			requestBody: map[string]interface{}{
				"start": -1000,
				"end":   -1,
			},
			expectedCode: http.StatusBadRequest,
			expectedKeys: []string{"error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/check-perfect-numbers", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := CheckPerfectNumbers(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, rec.Code)

			var response map[string]interface{}
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			for _, key := range tt.expectedKeys {
				_, exists := response[key]
				assert.True(t, exists)
			}
		})
	}
}
