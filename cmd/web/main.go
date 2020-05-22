package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ctfrancia/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
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
	// cfg := new(Config)
	// flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	// flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	// flag.StringVar(&cfg.DSN, "dsn", "web:Ressca000@/snippetbox?parseTime=true", "MySQL data source name")
	adr := flag.String("adr", ":4000", "HTTP Network address")
	dsn := flag.String("dsn", "web:Ressca000@/snippetbox?parseTime=true", "MySql data source name")
	secret := flag.String("secret", "Sfd%sfg*-sadfgFg@-ADg^$asdf+934GFASd", "Secret session Key")
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

	// template cache
	templateCache, err := newTemplateCache("./ui/html")
	if nil != err {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *adr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *adr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
