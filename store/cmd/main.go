package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/PierreLx/verre-tech/store"
	storesvc "github.com/PierreLx/verre-tech/store/implementation"
	"github.com/PierreLx/verre-tech/store/postgresql"
	"github.com/PierreLx/verre-tech/store/transport"
	httptransport "github.com/PierreLx/verre-tech/store/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", "localhost:8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "store", "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", "port=5432 host=localhost user=postgres password=JuVT75XX dbname=VerreTech sslmode=disable")
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	var svc store.Service
	{
		repository, err := postgresql.New(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		svc = storesvc.NewService(repository, logger)
	}

	var h http.Handler
	{
		endpoints := transport.MakeEndpoints(svc)
		ocTracing := kitoc.HTTPServerTrace()
		serverOptions := []kithttp.ServerOption{ocTracing}
		h = httptransport.NewService(endpoints, serverOptions, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}
