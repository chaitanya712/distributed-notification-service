package data

import "time"

type User struct {
	ID   string
	Name string
}

type Post struct {
	ID     string
	UserID string
	Body   string
	Time   time.Time
}

type Notification struct {
	ID        string
	UserID    string
	PostID    string
	Message   string
	Timestamp time.Time
}

var Users = []User{
	{ID: "u1", Name: "Alice"},
	{ID: "u2", Name: "Bob"},
	{ID: "u3", Name: "Charlie"},
	{ID: "u4", Name: "Diana"},
	{ID: "u5", Name: "Eve"},
}

var Followers = map[string][]string{
	"u1": {"u2", "u3"},
	"u2": {"u1", "u3", "u4"},
	"u3": {"u1"},
	"u4": {"u1", "u2", "u5"},
	"u5": {"u3"},
}

var Posts = map[string][]Post{
	"u1": {{ID: "p1", UserID: "u1", Body: "Hello from Alice", Time: time.Now()}},
	"u2": {{ID: "p2", UserID: "u2", Body: "Bob's first post", Time: time.Now()}},
	"u3": {{ID: "p3", UserID: "u3", Body: "Charlie updates you!", Time: time.Now()}},
	"u4": {{ID: "p4", UserID: "u4", Body: "Diana checking in", Time: time.Now()}},
	"u5": {{ID: "p5", UserID: "u5", Body: "Eve’s adventure", Time: time.Now()}},
}

var NotificationStore = map[string][]Notification{}
