package api

import (
	"flag"
    "fmt"
    "log/slog"
    "net/http"
    "os"
    "time"
    "encoding/json"
    "context"
    "sync"

    "github.com/joho/godotenv"

    polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
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
            return
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

func polygonAPI(ctx context.Context, client *polygon.Client, ticker string, wg *sync.WaitGroup) error{
    defer wg.Done()
    slog.Info("Retrieving stock data...", "Ticker", ticker)
    params := &models.ListAggsParams{
		Ticker:     ticker,
		Multiplier: 1,
		Timespan:   "day",
		From:       models.Millis(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
		To:         models.Millis(time.Date(2024, 3, 9, 0, 0, 0, 0, time.UTC)),
	}

    iter := client.ListAggs(ctx, params)

	
	for iter.Next() {
		slog.Info("Iteration Item", "msg", iter.Item())
	}
	if iter.Err() != nil {
		slog.Error("Error", "err", iter.Err())
	}

    return nil
}

func (app *application) stockInfoHandler(w http.ResponseWriter, r *http.Request) error{
    slog.Info("Stock Info Handler initiated...")

    slog.Info("Creating Polygon client")
    api_key:=os.Getenv("POLY_API_KEY")
    cli := polygon.New(api_key)

    var wg sync.WaitGroup
    ctx, cancel :=context.WithCancel(context.Background())
    defer cancel()

    tickers := []string{
        "AAPL",
        "VZ",
        "XEL",
        "CRWD",
        "VICI",
        "INTC",
    }

    // concurrency parrellism
    for _, ticker := range tickers {
        wg.Add(1)
        go polygonAPI(ctx, cli, ticker, &wg)
    }

    wg.Wait()
    slog.Info("All Goroutines have finished....")
    // End concurrency
    return nil
    
}

func ToJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
}

func InitServer() *http.Server {
    var cfg config
    godotenv.Load()
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
    mux.HandleFunc("/v1/getBatchStockInfo", makeHTTPHandleFunc(app.stockInfoHandler))
    mux.HandleFunc("/v1/getStockInfo", makeHTTPHandleFunc(app.stockInfoHandler))

    srv := newAPIServer(cfg, mux)
    
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
