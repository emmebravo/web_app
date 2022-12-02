package models

type User struct {
	ID        int
	FirstName string
	Surname   string
}

var (
	users  []*User // a slice holding pointers to User objs
	nextID = 1
)

func GetUser() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++                  // increment ID
	users = append(users, &u) // append to users collection, &u grab reference user obj that came in and store it
	return u, nil
}
