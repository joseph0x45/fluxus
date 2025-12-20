package db

import (
	"fluxus/logger"
	"fluxus/models"
	"fmt"
)

func (c *Conn) InsertTag(tag *models.Tag) (int, error) {
	const query = `
		insert or ignore into tags (
			id, name, owner
		)
		values (
			:id, :name, :owner
		)
	`
	res, err := c.db.NamedExec(query, tag)
	if err != nil {
		return -1, fmt.Errorf("Failed to insert tag: %w", err)
	}
	inserted, err := res.RowsAffected()
	if err != nil {
		logger.Err("Error getting affected rows")
		return 0, nil
	}
	return int(inserted), nil
}

func (c *Conn) GetUserTags(owner string) ([]models.Tag, error) {
	tags := make([]models.Tag, 0)
	const query = "select * from tags where owner=?"
	err := c.db.Select(&tags, query, owner)
	if err != nil {
		return nil, fmt.Errorf("Failed to get tags for user %s: %w", owner, err)
	}
	return tags, nil
}

func (c *Conn) DeleteTag(tagID, owner string) error {
	const query = "delete from tags where id=? and owner=?"
	_, err := c.db.Exec(query, tagID, owner)
	if err != nil {
		return fmt.Errorf("Failed to delete tag with id %s", tagID)
	}
	return nil
}
