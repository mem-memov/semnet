package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Entity struct {
	class uint
	fact uint
	user uint
	storage abstract.Storage
}

var _ abstractStory.Entity = Entity{}

func createEntity(storage abstract.Storage) (Entity, error) {

	class, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	fact, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	user, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(class, fact)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(fact, user)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:  class,
		fact: fact,
		user: user,
		storage: storage,
	}, nil
}

func readEntityByClass(storage abstract.Storage, class uint) (Entity, error) {

	_, fact, err := storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, user, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:  class,
		fact: fact,
		user: user,
		storage: storage,
	}, nil
}

func readEntityByFact(storage abstract.Storage, fact uint) (Entity, error) {

	class, user, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:  class,
		fact: fact,
		user: user,
		storage: storage,
	}, nil
}

func readEntityByUser(storage abstract.Storage, user uint) (Entity, error) {

	fact, _, err := storage.GetReference(user)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:  class,
		fact: fact,
		user: user,
		storage: storage,
	}, nil
}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetFact() uint {

	return e.fact
}

func (e Entity) GetUser() uint {

	return e.user
}

func (e Entity) PointToClass(class abstractClass.Entity) error {

	return e.storage.Connect(e.fact, class.GetStory())
}

func (e Entity) PointToFact(fact abstractFact.Entity) error {

	return e.storage.Connect(e.fact, fact.GetStory())
}


