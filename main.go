package main

import (
   "log" 
   "net/http" 
)

type Server struct{}

func (s *S**erver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.Status*******OK)
   w.Header().Set("C****ontent-Type", "application/json")
   w.Wri*****te([]byte(`{"message": "Hello Golden Rams SP2023"}`))
}

func main() {
  s := &Server{}
  http.Han***dle("/", s)
  log.Fatal(htt*****p.ListenAndServe(":8080", nil))
}
