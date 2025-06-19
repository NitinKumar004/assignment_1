package main

import (
	"fmt"
	"time"
)

type task struct {
	taskid    int
	taskname  string
	status    string
	createdat string
}

type taskTracker struct {
	tasks []task
}

// unique id generator with closure function to remember defined data at own their places
func idGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// this is greet function this will greet person according to current time
func greet(name string) {
	currenthour := time.Now().Hour()
	if currenthour < 12 {
		fmt.Println("Good Morning ", name)
	} else if currenthour < 17 {
		fmt.Println("Good Afternoon", name)
	} else {
		fmt.Println("Good Evening", name)

	}

}

// here is addtask function help us prevent to add duplicate task and add new task
func (t *taskTracker) addTask(description string, status string, getNextID func() int) {
	//checking duplicate if found we have to prevent in it
	//for i := range t {
	//	if t[i].taskname == description && t[i].status == status {
	//		fmt.Println("Task  have  been  already is in your list  you are trying to adding duplicate task : ", t.taskname)
	//		return
	//	}
	//}
	// if duplicate  not found we use  pointer to changes data directly in memory
	t.tasks = append(t.tasks, task{
		taskid:   getNextID(),
		taskname: description,
		status:   status,
		//adding current time when new task add
		createdat: time.Now().Format("2006-01-02 15:04:05"),
	})
	fmt.Println("Added task : ", description, "with id", getNextID())

}

// marking task as completed with their given task id
func (t *taskTracker) markingComplete(taskid int) {
	fmt.Println("Marking task : ", taskid)
	for i, tsk := range t.tasks {
		if tsk.taskid == taskid {
			t.tasks[i].status = "Completed"
			fmt.Println("task id :", tsk.taskid, "has been Completed")
			return
		}

	}

	fmt.Println("there is no task with id ", taskid)
}

// whatever we have task we are listing that task here
func Listalltask(taskdata []task) {
	for _, t := range taskdata {
		fmt.Println("Task id: ", t.taskid, " Description : ", t.taskname, " Status : ", t.status, " created at : ", t.createdat)

	}
}

// showing all pending tasks filter with Pending tasks
func Pendingtask(taskdata []task) {
	fmt.Println("Pending task Details : ")
	for _, t := range taskdata {

		if t.status == "Pending" {
			fmt.Println("Task id: ", t.taskid, " Description : ", t.taskname, " Status : ", t.status)
		}
	}
}

func main() {
	greet(" Sir/madam")

	tt := taskTracker{tasks: make([]task, 0)}
	//assigning clousre here in main function to remember data where it defined
	getNextID := idGenerator()
	tt.addTask("Buy groceries", "Pending", getNextID)
	tt.addTask("Go Gym", "Pending", getNextID)
	tt.addTask("Buy groceries", "Pending", getNextID)

	Listalltask(tt.tasks)
	tt.markingComplete(1)
	Pendingtask(tt.tasks)

	var process int
	fmt.Println("Do you want to add a task if yes : press 1 otherwise press 0 Thanks :")
	fmt.Scan(&process)

	for process == 1 {
		var taskname string
		fmt.Println("Enter your task name: ")
		fmt.Scan(&taskname)
		if taskname != "" {
			tt.addTask(taskname, "Pending", getNextID)
			Pendingtask(tt.tasks)
		}

		fmt.Println("want to add more?")
		fmt.Scan(&process)
	}
}
