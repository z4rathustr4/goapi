package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/z4rathustr4/goapi/internal/handlers"
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetReportCaller(true)
    var r *chi.Mux = chi.NewRouter()
    handlers.Handler(r)
    fmt.Println("INFO: Starting server")
    fmt.Println(`
   .aMMMMP .aMMMb         .aMMMb  dMMMMb  dMP 
  dMP"    dMP"dMP        dMP"dMP dMP.dMP amr  
 dMP MMP"dMP dMP        dMMMMMP dMMMMP" dMP   
dMP.dMP dMP.aMP        dMP dMP dMP     dMP    
VMMMP"  VMMMP"        dMP dMP dMP     dMP      
       
        `)

    err := http.ListenAndServe("localhost:8080", r)
    if err != nil {
        log.Error(err)
    }
}
