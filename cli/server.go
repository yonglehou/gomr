package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nytimes/gziphandler"
	"github.com/turbobytes/gomr"
	"log"
	"net/http"
	"time"
)

func getjoblist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jobs, err := gomr.FetchAllJobs()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	b, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	router := httprouter.New()
	router.GET("/api/joblist", getjoblist)
	s := &http.Server{
		Addr:           ":8181",
		Handler:        gziphandler.GzipHandler(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}