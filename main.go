package main
import (
    "net/http"
    "log"
    "github.com/rumyantseva/advent-2017/handlers"
    "github.com/rumyantseva/advent-2017/version"
    "os"
)
// How to try it: go run main.go
func main() {
    log.Print("Starting the service...")
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("Port is not set.")
    }
    r := handlers.Router()
    log.Print("The service is ready to listen and serve.")
    log.Fatal(http.ListenAndServe(":"+port, r))
    log.Printf(
        "Starting the service...\n commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release,
    )
    r := handlers.Router(version.BuildTime, version.Commit, version.Release)
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
    srv := &http.Server{
        Addr:    ":"  + port,
        Handler: r,
    }
    go func() {
        log.Fatal(srv.ListenAndServe())
    }()
    log.Pring("The service is ready to listen and serve.")
    killSignal := <-interrupt
    case os.kill:
        log.Print("Got SIGKILL...")
    case os.Interrupt:
        log.Print("Got SIGINT...")
    case syscall.SIGTERM:
        log.Print("Got SIGTERM...")
    }
    log.Print("The service is shutting down...")
    srv.Shutdown(context.Background())
    log.Print("Done")
}
