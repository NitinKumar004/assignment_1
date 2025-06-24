package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type task struct {
	Taskid   int    `json:"taskid"`
	Taskname string `json:"taskname"`
	Status   string `json:"status"`
}

type data struct {
	db *sql.DB
}

const complete = "complete"

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := make(map[string]string)
	response["message"] = "you are on the home page"

	jsondata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsondata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d *data) addtask(w http.ResponseWriter, r *http.Request) {
	var t task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = d.db.Exec("INSERT INTO tasks (taskname,status) VALUES (?,?)", t.Taskname, t.Status)
	if err != nil {
		http.Error(w, "Failed to insert task into database", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "added task successfully"}

	jsondata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsondata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d *data) alltask(w http.ResponseWriter, _ *http.Request) {
	rowsdata, err := d.db.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, "failed to fetch data ", http.StatusInternalServerError)
		return
	}

	var tasks []task
	for rowsdata.Next() {
		var t task
		err := rowsdata.Scan(&t.Taskid, &t.Taskname, &t.Status)
		if err != nil {
			http.Error(w, "failed to reading  data ", http.StatusInternalServerError)
		}
		tasks = append(tasks, t)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	tasksjson, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "error to converting data", http.StatusInternalServerError)
		return
	}
	w.Write(tasksjson)
}

func (d *data) gettask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var t task
	err = d.db.QueryRow("SELECT  * FROM tasks  WHERE taskid=?", id).Scan(&t.Taskid, &t.Taskname, &t.Status)
	if err != nil {
		http.Error(w, "failed to fetch data ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	taskjson, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "error to converting data", http.StatusInternalServerError)
	}
	w.Write(taskjson)

}

func (d *data) completetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	_, err = d.db.Exec("UPDATE tasks SET status=? WHERE taskid=?", complete, id)
	if err != nil {
		http.Error(w, "failed to complete task", http.StatusInternalServerError)
	}
	response := map[string]string{"message": "completed task successfully"}
	jsondata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

}

func (d *data) deletetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	_, err = d.db.Exec("DELETE FROM tasks where taskid=?", id)

	if err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return

	}
	response := map[string]string{"message": "deleted task successfully"}
	jsondata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

}

func main() {
	db := Calldatabase()
	defer db.Close()
	d := data{db: db}

	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/", home)
	http.HandleFunc("POST /task", d.addtask)
	http.HandleFunc("GET /task", d.alltask)
	http.HandleFunc("GET /task/{id}", d.gettask)
	http.HandleFunc("PATCH /task/{id}", d.completetask)
	http.HandleFunc("DELETE /task/{id}", d.deletetask)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
