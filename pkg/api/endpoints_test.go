package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCommentByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getComments)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	/*expected := `{"id":1,"comment":"good","comment_date":"04/12/2022","user_id":005}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}
