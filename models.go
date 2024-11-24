package main

type User struct {
	ID       int
	Username string
	Password string
}

type Post struct {
	ID      int
	UserID  int
	Title   string
	Content string
}
