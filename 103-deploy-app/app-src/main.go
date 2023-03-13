package main

import (
  "fmt"
  "net/http"

  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
  prometheus.CounterOpts{
    Name: "ping_request_count",
    Help: "No of request handled by Ping handler",
   },
)

func ping(w http.ResponseWriter, req *http.Request) {
  pingCounter.Inc()
  fmt.Fprintf(w, "pong")
}

var helloworldCounter = prometheus.NewCounter(
  prometheus.CounterOpts{
    Name: "helloworld_request_count",
    Help: "No of request handled by helloworld handler",
   },
)

func helloworld(w http.ResponseWriter, req *http.Request) {
  helloworldCounter.Inc()
  fmt.Fprintf(w, "hello world")
}

func main() {
  prometheus.MustRegister(pingCounter)
  prometheus.MustRegister(helloworldCounter)
  http.HandleFunc("/ping", ping)
  http.HandleFunc("/", helloworld)
  http.Handle("/metrics", promhttp.Handler())
  http.ListenAndServe(":8090", nil)
}
