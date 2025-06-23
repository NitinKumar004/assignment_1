package main

import (
	"encoding/json"
	"fmt"
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
	store []task
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

	d.store = append(d.store, t)
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
	jsondata, err := json.Marshal(d.store)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	n, err := w.Write(jsondata)
	if err != nil {
		fmt.Printf("Write failed: wrote %d bytes, error: %v\n", n, err)
	}
}

func (d *data) gettask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	fmt.Println("id we getting", id)

	for _, t := range d.store {
		if t.Taskid != id {
			continue
		}

		response := map[string]string{"message": "data fetched successfullyyyy"}
		jsonresponse, err := json.Marshal(response)

		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		n, err := w.Write(jsonresponse)
		if err != nil {
			fmt.Printf("Write failed: wrote %d bytes, error: %v\n", n, err)
			return
		}
	}
}

func (d *data) completetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	fmt.Println("id we getting", id)

	for idx, t := range d.store {
		if t.Taskid != id {
			continue
		}

		d.store[idx].Status = complete
		response := map[string]string{"message": "completed task successfully"}

		var jsonresponse []byte
		jsonresponse, err = json.Marshal(response)

		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonresponse)

		if err != nil {
			http.Error(w, "error to sending response", http.StatusInternalServerError)
			return
		}

		return
	}

	response := map[string]string{"message": "task not found"}
	jsonresponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	_, err = w.Write(jsonresponse)
	if err != nil {
		http.Error(w, "error to sending response", http.StatusInternalServerError)
		return
	}
}

func (d *data) deletetask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	fmt.Println("id we getting", id)

	for idx, t := range d.store {
		if t.Taskid != id {
			continue
		}

		d.store = append(d.store[:idx], d.store[idx+1:]...)
		jsondata := map[string]string{"message": "deleted task successfully"}

		var jsonresponse []byte

		jsonresponse, err = json.Marshal(jsondata)
		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		_, err = w.Write(jsonresponse)
		if err != nil {
			http.Error(w, "invalid response", http.StatusBadRequest)
			return
		}
	}

	response := map[string]string{"message": "task not found"}
	jsonresponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	n, err := w.Write(jsonresponse)
	if err != nil {
		fmt.Printf("Write failed: wrote %d bytes, error: %v\n", n, err)
		return
	}
}

func main() {
	d := data{}
	d.store = append(d.store, task{1, "do hard work", "pending"})

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
