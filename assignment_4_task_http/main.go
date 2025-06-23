package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type task struct {
	Taskid   int    `json:"taskid"`
	Taskname string `json:"taskname"`
	Status   string `json:"status"`
}
type data struct {
	store []task
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from main page"))
	w.WriteHeader(200)
}
func (d *data) addtask(w http.ResponseWriter, r *http.Request) {
	//d.store=append(d.store,task{1,"hr","45"})
	//fmt.Println("add task calling here")

	var t task

	//reading data for r.body using bodybytes
	//
	//bodybytes, err := io.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//
	//}
	//fmt.Println(string(bodybytes))

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {

		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	d.store = append(d.store, t)

	json.NewEncoder(w).Encode(map[string]string{"message": "added task successfully"})

}
func (d *data) alltask(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "all data fetched succesfully"})
	json.NewEncoder(w).Encode(d.store)
}
func (d *data) gettask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	fmt.Println("id we getting", id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	fmt.Println(id)
	for _, t := range d.store {
		if t.Taskid == id {
			//fmt.Println("task found here is task ")
			json.NewEncoder(w).Encode(map[string]string{"message": "data fetched successfully"})
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	//fmt.Println(y)
	json.NewEncoder(w).Encode(map[string]string{"message": "task not found"})

}
func (d *data) completetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	fmt.Println("id we getting", id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	fmt.Println(id)
	for idx, t := range d.store {
		if t.Taskid == id {
			//fmt.Println("task found", t)
			d.store[idx].Status = "completed"
			json.NewEncoder(w).Encode(map[string]string{"message": "completed task successfully"})
			json.NewEncoder(w).Encode(d.store[idx])
			return
		}
	}
	//fmt.Println(y)
	json.NewEncoder(w).Encode(map[string]string{"message": "task not found"})

}
func (d *data) deletetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	fmt.Println("id we getting", id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	fmt.Println(id)
	for idx, t := range d.store {
		if t.Taskid == id {
			//fmt.Println("task found", t)
			d.store[idx].Status = "completed"
			d.store = append(d.store[:idx], d.store[idx+1:]...)

			json.NewEncoder(w).Encode(map[string]string{"message": "deleted  task successfully"})
			//json.NewEncoder(w).Encode(d.store[idx])
			return
		}
	}
	//fmt.Println(y)
	json.NewEncoder(w).Encode(map[string]string{"message": "task not found"})

}

func main() {
	d := data{}
	d.store = append(d.store, task{1, "do hard work", "pending"})
	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/", home)
	http.HandleFunc("POST /task", d.addtask)
	http.HandleFunc("GET  /task", d.alltask)
	http.HandleFunc("GET /task/{id}", d.gettask)
	http.HandleFunc("PATCH /task/{id}", d.completetask)

	http.HandleFunc("DELETE /task/{id}", d.deletetask)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
	///ADD DELERTE FUNCTIONALIT
	//	//PU TO MAKRK TASK COMPLETE
}
