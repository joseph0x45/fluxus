package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func getDataDir() string {
	return filepath.Join(
		os.Getenv("HOME"),
		".local",
		"share",
		"fluxus",
	)
}

func GetDatabaseFilePath() string {
	return filepath.Join(
		getDataDir(),
		"fluxus.db",
	)
}

func EnsureDataDir() error {
	_, err := os.Stat(getDataDir())
	if errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(getDataDir(), 0755)
	}
	return nil
}

type Conn struct {
	db *sqlx.DB
}

const DB_SCHEMA = `
	create table if not exists users (
		id text not null primary key,
		username text not null unique,
		password text not null,
		safe_mode integer not null default 0
	);

	create table if not exists sessions (
		id text not null primary key,
		session_user text not null references users(id)
	);

	create table if not exists accounts (
		id text not null primary key,
		name text not null,
		balance integer not null,
		owner text not null references users(id),
		unique(name, owner)
	);

	create table if not exists tags (
		id text not null primary key,
		name text not null,
		owner text not null references users(id),
		unique(name, owner)
	);

`

func GetConn(resetDB bool) (*Conn, error) {
	if resetDB {
		if err := os.Remove(GetDatabaseFilePath()); err != nil {
			panic(err)
		}
	}
	if err := EnsureDataDir(); err != nil {
		return nil, fmt.Errorf("Setup failed: %w", err)
	}
	db, err := sqlx.Connect("sqlite3", GetDatabaseFilePath())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}
	_, err = db.Exec("PRAGMA foreign_keys=ON")
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to enable foreign keys: %w", err)
	}
	_, err = db.Exec(DB_SCHEMA)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to initialize database schema: %w", err)
	}

	return &Conn{db: db}, nil
}
