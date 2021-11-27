package fcc

type User struct {
	ID   string
	Name string
}

// Users Defined type pattern
type Users []User

func NewUsers(list []User) Users {
	result := make(Users, len(list))
	for i, value := range list {
		result[i] = value
	}
	return result
}

func (u Users) Filter(condition func(u User) bool) Users {
	result := u[:0]
	for _, user := range u {
		if condition(user) {
			result = append(result, user)
		}
	}
	return result
}
