package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"message":"you are on the home page"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAddTaskHandler(t *testing.T) {
	d := data{}
	testTask := task{
		Taskid:   1,
		Taskname: "Test task",
		Status:   "pending",
	}

	jsonTask, err := json.Marshal(testTask)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, "/task", bytes.NewBuffer(jsonTask))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.addtask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"message":"added task successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if len(d.store) != 1 {
		t.Errorf("store should contain only one task")
	}
}

func TestGetAllTasksHandler(t *testing.T) {
	d := data{
		store: []task{
			{1, "Task 1", "pending"},
			{2, "Task 2", "complete"},
		},
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/task", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.alltask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"taskid":1,"taskname":"Task 1","status":"pending"},{"taskid":2,"taskname":"Task 2","status":"complete"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestCompleteTaskHandler(t *testing.T) {
//	d := data{
//		store: []task{
//			{1, "Task 1", "pending"},
//		},
//	}
//
//	req, err := http.NewRequest("PATCH", "/task/1", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(d.completetask)
//
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	if d.store[0].Status != complete {
//		t.Errorf("task status not updated: got %v want %v",
//			d.store[0].Status, complete)
//	}
// }
//
// func TestDeleteTaskHandler(t *testing.T) {
//	d := data{
//		store: []task{
//			{1, "Task 1", "pending"},
//		},
//	}
//
//	req, err := http.NewRequest("DELETE", "/task/1", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(d.deletetask)
//
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	if len(d.store) != 0 {
//		t.Errorf("expected empty store after deletion, got length %d", len(d.store))
//	}
// }
