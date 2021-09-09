package semnet

type User struct {
	id uint
}

func (u User) GetId() uint {
	return u.id
}
