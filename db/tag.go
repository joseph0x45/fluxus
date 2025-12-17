package db

import (
	"fluxus/models"
	"fmt"
)

func (c *Conn) InsertTag(tag *models.Tag) error {
	const query = `
		insert or ignore into tags (
			id, name
		)
		values (
			:id, :name
		)
	`
	_, err := c.db.NamedExec(query, tag)
	if err != nil {
		return fmt.Errorf("Failed to insert tag: %w", err)
	}
	return nil
}

func (c *Conn) GetAllTags() ([]models.Tag, error) {
	tags := make([]models.Tag, 0)
	const query = "select * from tags"
	err := c.db.Select(tags, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all tags: %w", err)
	}
	return tags, nil
}

func (c *Conn) DeleteTag(id string) error {
	return nil
}
