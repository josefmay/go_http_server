package main

import (
	"flag"
    "fmt"
    "log/slog"
    "net/http"
    "os"
    "time"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

const version = "1.0.0"

type config struct {
    port int
    env  string
}

type application struct {
    config config
    logger *slog.Logger
}

func newAPIServer(cfg config, mux http.ServeMux) *http.Server{
	return &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      mux,
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }
}


func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "status: available")
    fmt.Fprintf(w, "environment: %s\n", app.config.env)
    fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
    logger.Info("Login Route Hit...")
    fmt.Fprintln(w, "login route")
}


func main() {
    var cfg config

    flag.IntVar(&cfg.port, "port", 3000, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    app := &application{
        config: cfg,
        logger: logger,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
    // mux.HandleFunc("/v1/register", app.loginHandler)
    // mux.HandleFunc("/v1/login", app.loginHandler)
    // mux.HandleFunc("/v1/account", app.healthcheckHandler)
    // mux.HandleFunc("/v1/watchlist", app.healthcheckHandler)
    // mux.HandleFunc("/v1/report", app.healthcheckHandler)


	
    srv := newAPIServer(cfg, mux)

    // Start the HTTP server.
    logger.Info("starting server...", "addr", srv.Addr, "env", cfg.env)
    
    err := srv.ListenAndServe()
    logger.Error(err.Error())
    os.Exit(1)
}


/*
/account authentication JWT
/watchlist watchlist database
/python backtester optimization engine
/report run reporting

sqlite3

react frontend?
*/


// func (s *APIServer) Run() {
// 	http.HandleFunc("/v1/healthcheck", func (w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path)
//         fmt.Fprintf(w, "Welcome to my website!")
//     })

// 	http.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

// 	http.ListenAndServe(s.listenAddr, nil)
// }

// func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
// 	log.Println(r.URL.Path)

// 	//Switch statement later

// 	if r.Method == "GET" {
// 		return s.handleGetAccount(w, r)
// 	}

// 	if r.Method == "POST" {
// 		return s.handleCreateAccount(w, r)

// 	}

// 	if r.Method == "DELETE" {
// 		return s.handleDeleteAccount(w, r)

// 	}

// 	return fmt.Errorf("Not a valid method")
// }

// func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
// 	ToJSON(w, 200, "Handled like a boss")
// 	return nil
// }

// func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }



// func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := f(w, r)
// 		if err != nil {
// 			ToJSON(w, http.StatusBadRequest, err.Error())
// 		}
// 	}
// }
