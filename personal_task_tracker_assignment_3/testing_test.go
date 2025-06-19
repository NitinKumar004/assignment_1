package main

import (
	"fmt"
	"testing"
)

func TestTasks(t *testing.T) {
	tt := taskTracker{tasks: make([]task, 0)}
	getNextID := idGenerator()
	tt.addTask("Buy groceries", "Pending", getNextID)

	if tt.tasks[len(tt.tasks)-1].taskname != "Buy groceries" {
		t.Errorf("taskTracker.addTask returned wrong task name")
	}

	tt.addTask("Go Gym", "Pending", getNextID)
	tt.addTask("Buy groceries", "Pending", getNextID)

	if len(tt.tasks) != 3 {
		t.Errorf("taskTracker.addTask returned %d tasks, want 3", len(tt.tasks))
	}
}

func TestPendingtask(t *testing.T) {
	tt := taskTracker{tasks: make([]task, 0)}

	getNextID := idGenerator()
	tt.addTask("Buy groceries", "Pending", getNextID)
	lastAdded := tt.tasks[len(tt.tasks)-1]
	if lastAdded.taskname != "Buy groceries" {
		t.Errorf("taskTracker.addTask returned wrong task name")
	}
	//lastAdded := tt.tasks[len(tt.tasks)-1]
	tt.markingComplete(tt.tasks[len(tt.tasks)-1].taskid)

	for _, tsk := range tt.tasks {
		fmt.Println(":Marking Tasks : ", tsk.taskid, " task name :", tsk.taskname, "status : ", tsk.status)
	}

	if tt.tasks[len(tt.tasks)-1].status != "Completed" {
		t.Errorf("taskTracker.addTask returned wrong status")
	}
}
