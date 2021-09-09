package semnet

type Story struct {
	id   uint
	user User
}

func newStory(id uint, user User) Story {
	return Story{
		id:   id,
		user: user,
	}
}

func (s Story) GetId() uint {
	return s.id
}

func (s Story) GetUser() User {
	return s.user
}
