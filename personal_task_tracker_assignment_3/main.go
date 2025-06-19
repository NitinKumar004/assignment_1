package main

//
//import (
//	"fmt"
//	"time"
//)
//
//type task struct {
//	taskid    int
//	taskname  string
//	status    string
//	createdat string
//}
//
//// unique id generator with closure function to remember defined data at own their places
//func idGenerator() func() int {
//	id := 0
//	return func() int {
//		id++
//		return id
//	}
//}
//
//// this is greet function this will greet person according to current time
//func greet(name string) {
//	currenthour := time.Now().Hour()
//	if currenthour < 12 {
//		fmt.Println("Good Morning ", name)
//	} else if currenthour < 17 {
//		fmt.Println("Good Afternoon", name)
//	} else {
//		fmt.Println("Good Evening", name)
//
//	}
//
//}
//
//// here is addtask function help us prevent to add duplicate task and add new task
//func addTask(taskdata *[]task, description string, status string, getNextID func() int) {
//	//checking duplicate if found we have to prevent in it
//	for _, t := range *taskdata {
//		if t.taskname == description && t.status == status {
//			fmt.Println("Task  have  been  already is in your list  you are trying to adding duplicate task : ", t.taskname)
//			return
//
//		}
//	}
//	// if duplicate  not found we use  pointer to changes data directly in memory
//	*taskdata = append(*taskdata, task{
//		taskid:   getNextID(),
//		taskname: description,
//		status:   status,
//		//adding current time when new task add
//		createdat: time.Now().Format("2006-01-02 15:04:05"),
//	})
//	fmt.Println("Added task : ", description, "with id", getNextID())
//
//}
//
//// marking task as completed with their given task id
//func markingComplete(taskdata *[]task, taskid int) {
//	for i, t := range *taskdata {
//		if t.taskid == taskid {
//			(*taskdata)[i].status = "Completed"
//			fmt.Println("task id :", t.taskid, "has been Completed")
//			return
//		}
//
//	}
//	fmt.Println("there is no task with id ", taskid)
//}
//
//// whatever we have task we are listing that task here
//func Listalltask(taskdata *[]task) {
//	for _, t := range *taskdata {
//		fmt.Println("Task id: ", t.taskid, " Description : ", t.taskname, " Status : ", t.status, " created at : ", t.createdat)
//
//	}
//}
//
//// showing all pending tasks filter with Pending tasks
//func Pendingtask(taskdata []task) {
//	fmt.Println("Pending task Details : ")
//	for _, t := range taskdata {
//		if t.status == "Pending" {
//			fmt.Println("Task id: ", t.taskid, " Description : ", t.taskname, " Status : ", t.status)
//		}
//	}
//}
//func main() {
//	greet(" Sir/madam")
//
//	taskdata := []task{}
//	//assigning clousre here in main function to remember data where it defined
//	getNextID := idGenerator()
//	addTask(&taskdata, "Buy groceries", "Pending", getNextID)
//	addTask(&taskdata, "Go Gym", "Pending", getNextID)
//	addTask(&taskdata, "Buy groceries", "Pending", getNextID)
//
//	Listalltask(&taskdata)
//	markingComplete(&taskdata, 1)
//	Pendingtask(taskdata)
//
//	var process int
//	fmt.Println("Do you want to add a task if yes : press 1 otherwise press 0 Thanks :")
//	fmt.Scan(&process)
//
//	if process == 1 {
//		var taskname string
//		fmt.Println("Enter your task name: ")
//		fmt.Scan(&taskname)
//		if taskname != "" {
//			addTask(&taskdata, taskname, "Pending", getNextID)
//
//		}
//	}
//
//}
