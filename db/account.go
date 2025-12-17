package db

import (
	"database/sql"
	"errors"
	"fluxus/models"
	"fmt"
)

func (c *Conn) InsertAccount(account models.Account) error {
	const query = `
		insert into accounts (
			id, name, balance, owner
		)
		values (
			:id, :name, :balance, :owner
		)
	`
	_, err := c.db.NamedExec(query, account)
	if err != nil {
		return fmt.Errorf("Failed to insert account: %w", err)
	}
	return nil
}

func (c *Conn) GetAccounts(owner string) ([]models.Account, error) {
	const query = "select * from accounts where owner=?"
	accounts := make([]models.Account, 0)
	err := c.db.Select(accounts, query, owner)
	if err != nil {
		return nil, fmt.Errorf("Failed to get accounts for owner %s: %w", owner, err)
	}
	return accounts, nil
}

func (c *Conn) GetAccountByName(owner, name string) (*models.Account, error) {
	const query = "select * from accounts where owner=? and name=?"
	account := &models.Account{}
	if err := c.db.Get(account, query, owner, name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to get account from ownner %s with name %s: %w", owner, name, err)
	}
	return account, nil
}

func (c *Conn) RenameAccount(id, newName string) error {
	const query = `
		update accounts set name=? where id=?
	`
	_, err := c.db.Exec(query, newName, id)
	if err != nil {
		return fmt.Errorf("Failed to update account %d name: %w", id, err)
	}
	return nil
}
