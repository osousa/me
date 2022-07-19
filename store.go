package main

import (
	"log"
)

type DBStore struct {
	DB Database
}

func NewDBStore(db Database) *DBStore {
	return &DBStore{DB: db}
}

func (store *DBStore) CreateTables() error {
	log.Println("Checking db...")
	tx, err := store.DB.GetDB().Begin()
	if err != nil {
		return err
	}

	queries := []string{
		`
			CREATE TABLE if not exists Experience (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  desc text NOT NULL,
			  position varchar(100) NOT NULL,
			  company varchar(100) DEFAULT NULL,
			  fk int(11) DEFAULT NULL,
			  PRIMARY KEY (id),
			  KEY Experience_User_FK (fk),
			  CONSTRAINT Experience_User_FK FOREIGN KEY (fk) REFERENCES User (id) ON DELETE CASCADE
			) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;
		`,
		`
			CREATE TABLE if not exists Post (
			  title varchar(100) CHARACTER SET ascii NOT NULL,
			  id int(11) NOT NULL AUTO_INCREMENT,
			  body text NOT NULL,
			  date datetime NOT NULL,
			  url varchar(100) CHARACTER SET utf8 NOT NULL,
			  abstract text DEFAULT NULL,
			  fk int(11) NOT NULL,
			  PRIMARY KEY (id),
			  KEY Post_User_FK (fk),
			  CONSTRAINT Post_User_FK FOREIGN KEY (fk) REFERENCES User (id)
			) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;
		`,
		`
			DROP TABLE IF EXISTS User;
			CREATE TABLE User (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  username varchar(100) NOT NULL,
			  password varchar(100) NOT NULL,
			  PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;
		`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
