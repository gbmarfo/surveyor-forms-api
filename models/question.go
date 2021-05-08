package main

import (
	"database/sql"
	"errors"
)

func (question *Question) getQuestion(db *sql.DB) error {
	return errors.New("not implemented")
}

func (question *Question) updateQuestion(db *sql.DB) error {
	return errors.New("not implemented")
}

func (question *Question) deleteQuestion(db *sql.DB) error {
	return errors.New("not implemented")
}

func (question *Question) createQuestion(db *sql.DB) error {
	return errors.New("not implemented")
}

func getQuestions(db *sql.DB, start, count int) ([]Question, error) {
	return nil, errors.New("not implemented")
}
