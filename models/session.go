package models

import (
	"crypto/sha512"
	"database/sql"
)
import "encoding/hex"

type Session struct {
	ID        int64
	Name      string
	Value     string
	HashToken string
}

type SessionService struct {
	DB *sql.DB
}

func StoreState(name, value, Token string) (*Session, error) {
	session := &Session{}
	session.Name = name
	session.Value = value
	session.HashToken = hashSha512(value)
	
	return session, nil
}

func hashSha512(value string) string {
	hash := sha512.Sum512([]byte(value))
	return hex.EncodeToString(hash[:])
}
