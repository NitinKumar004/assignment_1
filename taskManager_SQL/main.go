package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	response := map[string]string{"message": "you are on the home page"}
	jsondata, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(jsondata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (d *data) addtask(w http.ResponseWriter, r *http.Request) {
	var t task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &t); err != nil {
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
	if _, err = w.Write(jsondata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (d *data) alltask(w http.ResponseWriter, _ *http.Request) {
	rowsdata, err := d.db.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, "failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer rowsdata.Close()

	var tasks []task

	for rowsdata.Next() {
		var t task
		scanErr := rowsdata.Scan(&t.Taskid, &t.Taskname, &t.Status)
		if scanErr != nil {
			http.Error(w, "failed to read data", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	tasksjson, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "error converting data", http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(tasksjson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (d *data) gettask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var t task
	err = d.db.QueryRow("SELECT * FROM tasks WHERE taskid=?", id).Scan(&t.Taskid, &t.Taskname, &t.Status)
	if err != nil {
		http.Error(w, "failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	taskjson, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "error converting data", http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(taskjson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
		return
	}

	response := map[string]string{"message": "completed task successfully"}
	jsondata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(jsondata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (d *data) deletetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	_, err = d.db.Exec("DELETE FROM tasks WHERE taskid=?", id)
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
	if _, err = w.Write(jsondata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	db := Calldatabase()
	if db == nil {
		log.Fatal("failed to initialize database")
	}
	d := data{db: db}

	fmt.Println("Listening on port 8081...")
	http.HandleFunc("/", home)
	http.HandleFunc("POST /task", d.addtask)
	http.HandleFunc("GET /task", d.alltask)
	http.HandleFunc("GET /task/{id}", d.gettask)
	http.HandleFunc("PATCH /task/{id}", d.completetask)
	http.HandleFunc("DELETE /task/{id}", d.deletetask)

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
