package handler_test

import (
	"GO-RAG/api/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestPostNews(t *testing.T){
	testCases := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "Not Implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost,"/",nil)
			handler.PostNews()(w,r)

			if w.Result().StatusCode != tc.expectedStatus{
				t.Errorf("expected :%d, go :%d",tc.expectedStatus,w.Result().StatusCode)
			}
		})
	}

}