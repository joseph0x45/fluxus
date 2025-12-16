package db

import (
	"database/sql"
	"errors"
	"fluxus/models"
	"fmt"
)

func (c *Conn) InsertSession(session *models.Session) error {
	const query = `
		insert into sessions (
			id, session_user
		)
		values (
			:id, :session_user
		);
	`
	_, err := c.db.NamedExec(query, session)
	if err != nil {
		return fmt.Errorf("Failed to insert session: %w", err)
	}
	return nil
}

func (c *Conn) GetSession(id string) (*models.Session, error) {
	const query = "select * from sessions where id=?"
	session := &models.Session{}
	err := c.db.Get(session, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to get session %s: %w", id, err)
	}
	return session, nil
}

func (c *Conn) DeleteSession(id string) error {
	const query = "delete from sessions where id=?"
	_, err := c.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete session %s: %w", id, err)
	}
	return nil
}
