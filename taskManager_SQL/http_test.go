package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestHomeHandler(t *testing.T) {
//
//
//
//
//}

func TestAddTaskHandler(t *testing.T) {

	database := Calldatabase()
	d := data{db: database}

	testTask := task{
		Taskid:   1544,
		Taskname: "Sql test",
		Status:   "Pending",
	}

	jsonTask, err := json.Marshal(testTask)
	if err != nil {
		t.Fatalf("failed to marshal task: %v", err)
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, "/task", bytes.NewBuffer(jsonTask))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.addtask)
	handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
	//}
	var dbtaskname string
	err = d.db.QueryRow("SELECT taskname FROM tasks WHERE taskname = ?", testTask.Taskname).Scan(&dbtaskname)
	if err != nil {
		t.Fatalf("failed to fetch task from DB: %v", err)
	}
	if dbtaskname != testTask.Taskname {
		t.Errorf("taskname mismatch: got %v, want %v", dbtaskname, testTask.Taskname)
	}
}

func TestGetAllTasksHandler(t *testing.T) {
	db := Calldatabase()
	d := data{db: db}
	var x int = 55666
	_, err := db.Exec("INSERT INTO tasks (taskid, taskname, status) VALUES (?, ?, ?)", x, "Test Task", "Pending")
	if err != nil {
		t.Fatalf("failed to insert test task: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.alltask)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
	}

	var tasks []task
	if err := json.Unmarshal(rr.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	found := false
	for _, t := range tasks {
		if t.Taskid == x && t.Taskname == "Test Task" && t.Status == "Pending" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected task not found in response")
	}
}

func TestGetTaskHandler(t *testing.T) {
	database := Calldatabase()
	d := data{db: database}

	testask := task{
		Taskid:   7776,
		Taskname: "Sql test",
		Status:   "Pending",
	}
	_, err := d.db.Exec("INSERT INTO tasks (taskid, taskname, status) VALUES (?, ?, ?)", testask.Taskid, testask.Taskname, testask.Status)
	if err != nil {
		t.Fatalf("failed to insert test task: %v", err)
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/task/", http.NoBody)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.gettask)

	handler.ServeHTTP(rr, req)
	var dbtaskname string
	err = d.db.QueryRow("SELECT taskname FROM tasks WHERE taskname = ?", testask.Taskname).Scan(&dbtaskname)
	if err != nil {
		t.Fatalf("failed to fetch task from DB: %v", err)
	}
	if dbtaskname != testask.Taskname {
		t.Errorf("taskname mismatch: got %v, want %v", dbtaskname, testask.Taskname)
	}

}

func TestCompleteTaskHandler(t *testing.T) {
	database := Calldatabase()
	d := data{db: database}

	testask := task{
		Taskid:   66,
		Taskname: "st",
		Status:   "Pending",
	}
	_, err := d.db.Exec("INSERT INTO tasks (taskid, taskname, status) VALUES (?, ?, ?)", testask.Taskid, testask.Taskname, testask.Status)
	if err != nil {
		t.Fatalf("failed to insert test task: %v", err)
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPatch, "/task/", http.NoBody)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "66")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.completetask)
	handler.ServeHTTP(rr, req)
	var statusc string
	err = d.db.QueryRow("SELECT status FROM tasks WHERE taskname = ?", testask.Taskname).Scan(&statusc)
	if statusc != "complete" {
		t.Errorf("task is still pending")
		return
	}

}

func TestDeleteTaskHandler(t *testing.T) {
	database := Calldatabase()
	d := data{db: database}

	testask := task{
		Taskid:   126,
		Taskname: "Sltest",
		Status:   "Pending",
	}
	_, err := d.db.Exec("INSERT INTO tasks (taskid, taskname, status) VALUES (?, ?, ?)", testask.Taskid, testask.Taskname, testask.Status)
	if err != nil {
		t.Fatalf("failed to insert test task: %v", err)
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodDelete, "/task/", http.NoBody)

	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "126")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(d.deletetask)
	handler.ServeHTTP(rr, req)
	var statusc string
	err = d.db.QueryRow("SELECT status FROM tasks WHERE taskname = ?", testask.Taskname).Scan(&statusc)
	if statusc != "" {
		t.Errorf("TASK NIT DELETE SOMWETHING ")
		return
	}

}
