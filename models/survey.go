package main

import (
	"database/sql"
	"errors"
)

func (survey *Survey) getSurvey(db *sql.DB) error {
	return errors.New("not implemented")
}

func (survey *Survey) updateSurvey(db *sql.DB) error {
	return errors.New("not implemented")
}

func (survey *Survey) deleteSurvey(db *sql.DB) error {
	return errors.New("not implemented")
}

func (survey *Survey) createSurvey(db *sql.DB) error {
	return errors.New("not implemented")
}

func getSurveys(db *sql.DB, start, count int) ([]Survey, error) {
	return nil, errors.New("not implemented")
}
