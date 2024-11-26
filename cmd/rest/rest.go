package rest

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"switchcraft/core"
	"switchcraft/types"
	"sync"
	"time"
)

const rwTimeout = 15 * time.Second

type handler func(http.ResponseWriter, *http.Request)

func Start(logger *types.Logger, core *core.Core, port string) *http.Server {
	mux := http.NewServeMux()

	addRoutes(logger, core, mux)

	server := &http.Server{
		Handler:           trace(logger, mux),
		Addr:              net.JoinHostPort("localhost", port),
		ReadHeaderTimeout: rwTimeout,
		ReadTimeout:       rwTimeout,
		WriteTimeout:      rwTimeout,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			os.Stderr.WriteString(err.Error())
			wg.Done()
		}
	}()
	fmt.Println("Server listening on", port)
	wg.Wait()

	return server
}

func decodeBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
