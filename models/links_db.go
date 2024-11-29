package models

import (
	"errors"
	"fmt"
	"log"
	"time"
	"urlShortener/database"
)

func SaveLink(item LinkSave) error {
	db, err := database.GetDb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, saveErr := db.Exec("INSERT INTO links (url, short, users_id, expired_at) VALUES ($1, $2, $3, $4)",
		item.Url, item.Short, item.UserId, item.ExpiredAt)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func GetLink(short string) (*Link, error) {
	db, err := database.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result := db.QueryRow("SELECT * FROM links WHERE short = $1", short)
	link := Link{}
	scanErr := result.Scan(&link.ID, &link.Url, &link.Short, &link.UserId, &link.ExpiredAt, &link.Counter, &link.LastUsedAt)
	if scanErr != nil {
		log.Println(scanErr)
		return nil, scanErr
	}
	return &link, nil
}

func DeleteLink(item Link) error {
	db, err := database.GetDb()
	if err != nil {
		return err
	}
	defer db.Close()
	result, deleteErr := db.Exec("DELETE FROM links WHERE id = $1", item.ID)
	if deleteErr != nil {
		return deleteErr
	}
	affected, rowErr := result.RowsAffected()
	if rowErr != nil {
		return rowErr
	}
	if affected == 0 {
		return errors.New("no links deleted")
	}

	return nil
}

func SearchLink(search string) (*Link, error) {
	db, err := database.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result := db.QueryRow("SELECT * FROM links WHERE short = $1 OR url = $1", search)
	link := Link{}
	scanErr := result.Scan(&link.ID, &link.Url, &link.Short, &link.UserId, &link.ExpiredAt, &link.Counter, &link.LastUsedAt)
	if scanErr != nil {
		log.Println(scanErr)
		return nil, scanErr
	}
	return &link, nil
}

func UserLinks(userId int) ([]Link, error) {
	db, err := database.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, resErr := db.Query("SELECT * FROM links WHERE users_id = $1", userId)
	defer rows.Close()
	var links []Link
	if resErr != nil {
		return links, err
	}

	for rows.Next() {
		link := Link{}
		err := rows.Scan(&link.ID, &link.Url, &link.Short, &link.UserId, &link.ExpiredAt, &link.Counter, &link.LastUsedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		links = append(links, link)
	}
	return links, nil
}

func LinkAddCounter(item Link) error {
	db, err := database.GetDb()
	if err != nil {
		return err
	}
	counter := item.Counter + 1
	usedDate := time.Now().Format(time.RFC3339)
	defer db.Close()
	_, updateErr := db.Exec("UPDATE links SET counter=$2, last_used_at=$3 WHERE id = $1", item.ID, counter, usedDate)
	if updateErr != nil {
		return updateErr
	}
	return nil
}
