package services

import (
	"Rest_Api/AuthorDTO"
	"Rest_Api/datasources"
	"Rest_Api/errors"
)

const (
	queryInsertAuthor = "INSERT INTO authors(name, surname, email, age) VALUES (?, ?, ?, ?);"
)

func CreateAuthor(author *Authors.Author) *errors.RestErr {
	stmt, err := datasources.Client.Prepare(queryInsertAuthor)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(author.Name, author.Surname, author.Email, author.Age)

	if err != nil {
		return errors.NewInternalServerError("Can't store the data into database.")
	}

	return nil
}
