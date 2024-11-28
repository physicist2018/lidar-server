package db

import (
	"errors"
	"log"
	"os"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type DB struct {
	Session *r.Session
}

func NewConnection(host string, user string, pass string) (*DB, error) {
	session, err := r.Connect(
		r.ConnectOpts{
			Address:  host,
			Username: user,
			Password: pass,
			MaxOpen:  10,
			Database: "Lidar",
		},
	)
	if err != nil {
		return nil, err
	}

	r.Log.Out = os.Stderr
	return &DB{
		Session: session,
	}, nil
}

func (db *DB) Ping() (string, error) {
	res, err := r.Expr("Hello world").Run(db.Session)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return response, nil
}

func (db *DB) Close() error {
	if db.Session == nil {
		return errors.New("возможно, соединение уже закрыто")
	}

	return db.Session.Close()
}
