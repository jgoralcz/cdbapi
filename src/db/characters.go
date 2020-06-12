package db

import (
	"encoding/json"
	"log"
)

type Character struct {
	Name   string
	Series string
}

func GetRandomCharacter(limit int) string {
	rows := PoolQueryRows("SELECT name, series FROM waifu_schema.waifu_table ORDER BY RANDOM() LIMIT 10")
	if rows == nil {
		return "[]"
	}

	characters := []Character{}
	for rows.Next() {
		character := new(Character)
		err := rows.Scan(&character.Name, &character.Series)
		if err != nil {
			log.Println(err)
			break
		}
		characters = append(characters, *character)
	}

	charactersJSON, err := json.Marshal(characters)
	if err != nil {
		log.Println(err)
		return "[]"
	}

	return string(charactersJSON)
	// err := row.Scan(&character.Name, &character.Series)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil
	// }

	// if rows.Err() != nil {
	// 	return rows.Err()
	// }

	// log.Print("character name:")
	// log.Print(character.Name, character.Series)
	// return &characters
}
