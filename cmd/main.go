package main

import (
	"fmt"
	http1 "invoices-manager/pkg/http"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/synthesio/zconfig"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(errors.Wrap(err, "unable to load env"))
	}

	var s Server
	if err := zconfig.Configure(&s); err != nil {
		log.Fatal(errors.Wrap(err, "unable to configure"))
	}

	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc(http1.GetAllUsers, s.GetAllUsers).Methods("GET")
	mux.HandleFunc(http1.GetUser, s.GetUser).Methods("GET")

	mux.HandleFunc(http1.GetUserInvoices, s.GetUserInvoices).Methods("GET")
	mux.HandleFunc(http1.GetAllInvoices, s.GetAllInvoices).Methods("GET")
	mux.HandleFunc(http1.CreateInvoice, s.CreateInvoice).Methods("POST")

	mux.HandleFunc(http1.ExecuteTransaction, s.ExecuteTransaction).Methods("POST")

	go func() {
		time.Sleep(time.Second)
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d%s", s.Port, http1.GetAllUsers))
		if err != nil || resp.StatusCode != http.StatusOK {
			logrus.Errorf("Healthcheck request failed with status %d and err %+v", resp.StatusCode, err)
			return
		}
		logrus.Info("Server started.")
	}()

	logrus.Info("Starting server...")
	if err := s.ListenAndServe(mux); err != nil {
		logrus.WithError(err).Fatal("unable to start server")
	}
}

type DB struct {
	User      string `key:"user" validate:"required"`
	Password  string `key:"password"  validate:"required"`
	Port      int    `key:"port"  validate:"required"`
	InnerPort int    `key:"inner_port"  validate:"required"`
	Name      string `key:"db"  validate:"required"`
	*sqlx.DB
}

type Server struct {
	DB   `key:"postgres"`
	Port int    `key:"app_port"`
	Host string `key:"app_host"`
	*http.Server
}

func (s Server) ListenAndServe(mux *mux.Router) error {
	s.Server.Addr = fmt.Sprintf("%s:%d", s.Host, s.Port)
	s.Server.Handler = alice.New(http1.All()...).Then(mux)
	return s.Server.ListenAndServe()
}

func (s *Server) Init() error {
	// Waiting for the database to be up.
	time.Sleep(time.Second * 5)
	dsn := fmt.Sprintf("host=database port=%d user=%s password=%s dbname=%s sslmode=disable", s.DB.InnerPort, s.DB.User, s.DB.Password, s.DB.Name)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.WithError(err).Fatal("unable to start db: ", err)
		return err
	}
	if err := http1.Check(db); err != nil {
		return err
	}
	s.DB.DB = db
	return nil
}
