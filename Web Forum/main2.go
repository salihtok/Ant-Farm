package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver; the underscore indicates it's used for initialization only.
)

/*func main() {
	// Open a SQLite database. If it doesn't exist, it will be created at the specified path.
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close() // Ensure the database connection is closed when the function returns.

	// Attempt to create the necessary database tables.
	if err := createTables1(db); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	// Parse command line arguments to determine the required operation.
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide a command")
		return
	}

	// Handle different commands
	switch args[0] {
	case "adduser":
		// Ensure the correct number of arguments for the adduser command.
		if len(args) != 3 {
			fmt.Println("Usage: adduser <email> <name>")
			return
		}
		email := args[1]
		name := args[2]
		// Add user to the database
		if err := addUser(db, email, name); err != nil {
			log.Fatal("Failed to add user:", err)
		}
		fmt.Println("User added successfully")
	case "listusers":
		// Retrieve and display all users
		users, err := listUsers(db)
		if err != nil {
			log.Fatal("Failed to list users:", err)
		}
		for _, user := range users {
			fmt.Printf("ID: %d, Email: %s, Name: %s\n", user.id, user.email, user.name)
		}
	case "addpost":
		if len(args) != 4 {
			fmt.Println("Usage: addpost <userID> <catagory> <post>")
			return
		}
		userID := args[1] // kontrol sağlamak gerekebilir.
		catagory := args[2]
		post := args[3]
		if err := AddPosts(db, userID, catagory, post); err != nil {
			log.Fatal("failed to add post", err)
		}
		fmt.Println("Post added successfully")
	case "showposts":

	default:
		fmt.Println("Unknown command")
	}
}*/

// createTables creates the required database tables if they do not already exist.
func createTables1(db *sql.DB) error {
	tableCreationQueries := []string{`CREATE TABLE IF NOT EXISTS Users (
		UserID INTEGER PRIMARY KEY AUTOINCREMENT,
		Email TEXT NOT NULL UNIQUE,
		Name TEXT NOT NULL
	);`,
		`CREATE TABLE IF NOT EXISTS Categories (
		CategoryID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL UNIQUE
	);`,
		`CREATE TABLE IF NOT EXISTS Posts (
		PostID INTEGER PRIMARY KEY AUTOINCREMENT,
		UserID INTEGER,
		CategoryID INTEGER,
		Title TEXT NOT NULL,
		Content TEXT NOT NULL,
		Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (UserID) REFERENCES Users(UserID),
		FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
	);`,
		`CREATE TABLE IF NOT EXISTS Comments (
		CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
		PostID INTEGER,
		UserID INTEGER,
		Content TEXT NOT NULL,
		Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (PostID) REFERENCES Posts(PostID),
		FOREIGN KEY (UserID) REFERENCES Users(UserID)
	);`,
		`CREATE TABLE IF NOT EXISTS Interactions (
		InteractionID INTEGER PRIMARY KEY AUTOINCREMENT,
		PostID INTEGER,
		UserID INTEGER,
		Type TEXT CHECK (Type IN ('like', 'dislike')),
		FOREIGN KEY (PostID) REFERENCES Posts(PostID),
		FOREIGN KEY (UserID) REFERENCES Users(UserID)
	);`}

	for _, query := range tableCreationQueries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil

}

// addUser inserts a new user into the database.
func addUser(db *sql.DB, email, name string) error {
	// Prepare the insert statement to avoid SQL injection
	statement, err := db.Prepare("INSERT INTO Users (Email, Name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close() // Ensure the statement is closed after it's no longer needed.
	_, err = statement.Exec(email, name)
	return err
}

func AddPosts(db *sql.DB, UserID, catagory, theText string) error {
	statement, err := db.Prepare("INSERT INTO Posts  (Title, Content) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	return err
}

// listUsers retrieves all users from the database.
// listUsers retrieves all users from the database.
func listUsers(db *sql.DB) ([]*User, error) {
	// Query all users
	rows, err := db.Query("SELECT UserID, Email, Name FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after processing.

	var users []*User
	for rows.Next() {
		var u User
		// Scan each row into a User struct
		err := rows.Scan(&u.id, &u.email, &u.name)
		if err != nil {
			// If an error occurs during scan, return immediately with the error
			return nil, err
		}
		// If no error, append the user struct to the users slice
		users = append(users, &u)
	}
	return users, nil
}

// User represents a user in the database.
type User struct {
	id    int
	email string
	name  string
}
type Post struct {
	PostID  int
	UserID  int
	Title   string
	Content string
	User
}
