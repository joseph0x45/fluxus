package db

import (
	"database/sql"
	"errors"
	"fluxus/models"
	"fmt"
)

func (c *Conn) GetUser(by, value string) (*models.User, error) {
	const get_by_id_query = "select * from users where id=?"
	const get_by_username_query = "select * from users where username=?"
	var err error
	user := &models.User{}
	if by == "id" {
		err = c.db.Get(user, get_by_id_query, value)
	} else {
		err = c.db.Get(user, get_by_username_query, value)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to get user by %s: %w", by, err)
	}
	return user, nil
}

func (c *Conn) InsertUser(user *models.User) error {
	const query = `
		insert into users (
			id, username, password
		)
		values (
			:id, :username, :password
		);
	`
	_, err := c.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}
	return nil
}
