package models

import (
	"crypto/sha512"
	"fmt"
	"zombiezen.com/go/sqlite"
)
import "encoding/hex"

type Session struct {
	ID        int64
	Name      string
	Value     string
	HashToken string
}

type SessionService struct {
	DB *sqlite.Conn
}

func (ss *SessionService) StoreState(name, value, Token string) (*Session, error) {
	session := &Session{}
	session.Name = name
	session.Value = value
	session.HashToken = hashSha512(Token)
	stmt, _, err := ss.DB.PrepareTransient(`INSERT OR REPLACE INTO session_stores (name, value, HashToken) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Finalize()
	stmt.BindText(1, session.Name)
	stmt.BindText(2, session.Value)
	stmt.BindText(3, session.HashToken)
	_, err = stmt.Step()
	if err != nil {
		return nil, err
	}
	session.ID = ss.DB.LastInsertRowID()
	return session, nil
}

func (ss *SessionService) GetState(token, name string) (*Session, error) {
	session := &Session{}
	session.Name = name
	session.HashToken = hashSha512(token)
	stmt, _, err := ss.DB.PrepareTransient(`SELECT id, value FROM session_stores WHERE HashToken = ? AND Name = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Finalize()
	stmt.BindText(1, session.HashToken)
	stmt.BindText(2, session.Name)
	rows, err := stmt.Step()
	if err != nil {
		return nil, err
	}
	if !rows {
		return nil, fmt.Errorf("No rows returned for name %s", name)
	}
	stmt.Step()
	session.Name = stmt.GetText("id")
	session.Value = stmt.GetText("value")
	return session, nil

	return session, nil
}

func hashSha512(value string) string {
	hash := sha512.Sum512([]byte(value))
	return hex.EncodeToString(hash[:])
}
