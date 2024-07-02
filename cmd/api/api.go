package api

import (
	"flag"
    "fmt"
    "log/slog"
    "net/http"
    "os"
    "time"
    "encoding/json"
)

type apiFunc func(http.ResponseWriter, *http.Request) error
/*
makeHTTPHandleFunc adds error handling to route handlers
Golang http handlers do not return an error,
therefore this workaround is need for error handling
*/
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			ToJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}

const version = "1.0.0"

type config struct {
    port int
    env  string
}

type application struct {
    config config
    logger *slog.Logger
}


func newAPIServer(cfg config, mux *http.ServeMux) *http.Server{
	return &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      mux,
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
}


func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) error{
    fmt.Fprintln(w, "status: available")
    fmt.Fprintf(w, "environment: %s\n", app.config.env)
    _, err := fmt.Fprintf(w, "version: %s\n", version)
    if err != nil{
        return err
    }
    return nil
}

// func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) error{
//     logger.Info("Login Route Hit...")
//     fmt.Fprintln(w, "login route")
// }

// func (app *application) registerHandler(w http.ResponseWriter, r *http.Request) error{
//     logger.Info("Login Route Hit...")
//     fmt.Fprintln(w, "login route")
// }

// func (app *application) accountHandler(w http.ResponseWriter, r *http.Request) error{
//     logger.Info("Login Route Hit...")
//     fmt.Fprintln(w, "login route")
// }

// func (app *application) reportHandler(w http.ResponseWriter, r *http.Request) error{
//     logger.Info("Login Route Hit...")
//     fmt.Fprintln(w, "login route")
// }

func ToJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
}

func InitServer() *http.Server {
    var cfg config

    flag.IntVar(&cfg.port, "port", 8080, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    app := &application{
        config: cfg,
        logger: logger,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", makeHTTPHandleFunc(app.healthcheckHandler))
    // mux.HandleFunc("/v1/register", app.registernHandler)
    // mux.HandleFunc("/v1/login", app.loginHandler)
    // mux.HandleFunc("/v1/account", app.healthcheckHandler)
    
    
    // mux.HandleFunc("/v1/watchlist", app.healthcheckHandler)
    // mux.HandleFunc("/v1/report", app.healthcheckHandler)a


	
    srv := newAPIServer(cfg, mux)

    // Start the HTTP server.
    // logger.Info("starting server...", "addr", srv.Addr, "env", cfg.env)
    
    return srv
}


/*
/account authentication JWT
/watchlist watchlist database
/backtest python backtester optimization engine
/report run reporting

sqlite3

react frontend?

when you register no token
when you login given token
need token for

/report grabs report

/backtest
where to hold results
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
