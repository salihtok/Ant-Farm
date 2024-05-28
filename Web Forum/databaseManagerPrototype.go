package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver; the underscore indicates it's used for initialization only.
)

func main() {
	InitializeDB()
}

func InitializeDB() {
	db, err := sql.Open("sqlite3", "/home/salih/Documents/Go Projects/Web Forum/protoype.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	if err := TableGeneration(db); err != nil {
		log.Fatal("failed to create and doing something for tables ", err)
	}

}

func TableGeneration(db *sql.DB) error { // table creation.
	creationQueries := []string{
		`CREATE TABLE IF NOT EXISTS Users (
			UserID INTEGER PRIMARY KEY AUTOINCREMENT,
			Email TEXT NOT NULL UNIQUE,
			Name TEXT NOT NULL,
			Password TEXT NOT NULL,
			UserLevel INTEGER NOT NULL
	);`,
		`CREATE TABLE IF NOT EXISTS Posts (
		PostID INTEGER PRIMARY KEY AUTOINCREMENT,
		UserID INTEGER,
		CategoryID INTEGER,
		Title TEXT NOT NULL,
		Content TEXT NOT NULL,
		Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		Likes INTEGER,
		Dislikes INTEGER, 
		FOREIGN KEY (UserID) REFERENCES Users(UserID),
		FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
	);`,
		`CREATE TABLE IF NOT EXISTS Comments (
		CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
		PostID INTEGER,
		UserID INTEGER,
		Content TEXT NOT NULL,
		Likes INTEGER,
		Dislikes INTEGER,
		Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (PostID) REFERENCES Posts(PostID),
		FOREIGN KEY (UserID) REFERENCES Users(UserID)
	);`}

	for _, query := range creationQueries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func SingUpOperations(db *sql.DB, Email, Name, Password string, UserLevel int) error { // adding users actually.
	statement, err := db.Prepare("INSERT INTO Users (Email, Name, Password, UserLevel) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	return err
}
func AddingPosts(db *sql.DB, UserID, CatagoryID int, Title, Content string) error {
	statement, err := db.Prepare("INSERT INTO Posts (UserID, CatagoryID, Title, Content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	return err
}
func AddingComments(db *sql.DB, PostID, UserID int, Content string) error {
	statement, err := db.Prepare("INSERT INTO Comments (PostID, UserID, Content) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	return err
}

// could be a function that adding catagories for admins. for versatility.
func LikesDislikes(db *sql.DB, PostID, ldP int) { // should be versatile.

}
