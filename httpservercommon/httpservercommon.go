package httpservercommon

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
TODO: graceful shutdown on signal
// https://gist.github.com/rgl/0351b6d9362abb32d6b55f86bd17ab65
*/
func Serve(listenPort int, mux *http.ServeMux) {
	listenAddress := fmt.Sprintf(":%d", listenPort)

	httpServer := http.Server{
		Addr:    listenAddress,
		Handler: mux}

	go func() {
		log.Printf("Listening at http://%s", listenAddress)
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe Error: %v", err)
		}
	}()

	time.Sleep(25 * time.Millisecond)
	log.Printf("Press <Return> to shutdown server")

	waitForEnter()

	log.Printf("Initiating HTTP Server shutdown...")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}

	log.Printf("HTTP Server Shutdown complete")
}

func waitForEnter() {
	var input string
	fmt.Scanln(&input)
}
