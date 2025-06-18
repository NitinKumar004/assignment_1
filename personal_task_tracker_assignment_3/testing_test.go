package main

import (
	"testing"
)

func TestTasks(t *testing.T) {
	tt := taskTracker{tasks: make([]task, 0)}
	getNextID := idGenerator()
	tt.addTask("Buy groceries", "Pending", getNextID)
	tt.addTask("Go Gym", "Pending", getNextID)
	tt.addTask("Buy groceries", "Pending", getNextID)
	tt.markingComplete(1)
	pending, complete := tt.test()
	if pending != 2 {
		t.Errorf("Expected 2 pending tasks, got %d", pending)
	}
	if complete != 1 {
		t.Errorf("Expected 1 completed task, got %d", complete)
	}

}
