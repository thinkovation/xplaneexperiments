package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type LocList struct {
	Locs []string
}

//"github.com/gorilla/handlers"
type myServer struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

// NewServer - this is the init function for the server process
func NewServer(port string) *myServer {

	//create server

	s := &myServer{
		Server: http.Server{
			Addr:         ":" + port,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		shutdownReq: make(chan bool),
	}

	router := mux.NewRouter()

	//register handlers
	router.HandleFunc("/", s.RootHandler)
	router.HandleFunc("/dets", s.DetsHandler)
	router.HandleFunc("/vals", s.ValsHandler)
	router.HandleFunc("/{valname}", s.ValHandler)

	s.Handler = cors.AllowAll().Handler(gziphandler.GzipHandler(router))

	return s
}
func (s *myServer) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}
	log.Printf("Stopping API server ...")
	close(done)
	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}

}

func (s *myServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi\n"))
}

func (s *myServer) DetsHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(X.FD.Vals)
	w.Write([]byte(b))
}
func (s *myServer) ValsHandler(w http.ResponseWriter, r *http.Request) {
	mm := X.GetVals()

	b, _ := json.Marshal(mm)
	w.Write([]byte(b))
}
func (s *myServer) ValHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		vars := mux.Vars(r)
		valname := string(vars["valname"])
		v, verr := X.GetValue(valname)
		msg := ""
		if verr != nil {
			msg = "Could not get value"

		} else {
			newvalstring := r.URL.Query().Get("val")
			if newvalstring != "" {
				newval, verr := strconv.ParseFloat(newvalstring, 64)
				if verr == nil {
					var cmd Command
					cmd.Message = uint32(v.MsgType)
					cmd.Data = [8]float32{-999, -999, -999, -999, -999, -999, -999, -999}
					cmd.Data[v.Idx] = float32(newval)
					X.Send(cmd)
					msg = "OK"

				} else {
					msg = "Invalid Value Sent"
				}

			} else {
				msg = strconv.FormatFloat(v.Value, 'f', 4, 64)
			}
		}
		w.Write([]byte(msg))
	}

}
func getTokenFromRequest(r *http.Request) string {
	var tmptoken string
	tmptoken = r.Header.Get("wf-tkn")
	if tmptoken == "" {
		tmptoken = r.URL.Query().Get("wf_tkn")
	}
	return tmptoken
}
