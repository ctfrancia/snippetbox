package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ctfrancia/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

// Config is our config struct at runtime.
type Config struct {
	Addr      string
	StaticDir string
	DSN       string
}

func main() {
	// command line flag, default address, short descriptor
	// addr := flag.String("addr", ":4000", "Http network address") // returns a pointer
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	// flag.StringVar(&cfg.DSN, "dsn", "web:Ressca000@/snippetbox?parseTime=true", "MySQL data source name")
	dsn := flag.String("dsn", "web:Ressca000@/snippetbox?parseTime=true", "MySql data source name")
	// Parse is used to Parse the flag
	flag.Parse()

	// creating custom error handlers, one for info other for errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
