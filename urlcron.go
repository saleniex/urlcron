package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"urlcron/metric"
	"urlcron/runner"
	"urlcron/schedule"
)

func main() {
	go func() {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprint(writer, "<a href=/metrics>metrics</a>")
		})
		http.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprint(writer, metric.PrometheusDump())
		})
		log.Fatal(http.ListenAndServe(env("ADDR", ":80"), nil))
	}()

	loaders := schedule.LoaderSet{Loaders: []schedule.Loader{
		schedule.NewFileLoader("crontab"),
		schedule.NewTextLoader(env("CRONTAB", "")),
	}}
	runner.New(loaders).Run()
}

func env(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		return val
	}
}
