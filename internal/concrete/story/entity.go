package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Entity struct {
	class   uint
	fact    uint
	position uint
	user    uint
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

	position, err := storage.Create()
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

	err = storage.SetReference(fact, position)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(position, user)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		fact:    fact,
		position: position,
		user:    user,
		storage: storage,
	}, nil
}

func readEntityByClass(storage abstract.Storage, class uint) (Entity, error) {

	_, fact, err := storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, position, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	_, user, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		fact:    fact,
		position: position,
		user:    user,
		storage: storage,
	}, nil
}

func readEntityByFact(storage abstract.Storage, fact uint) (Entity, error) {

	class, position, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	_, user, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		fact:    fact,
		position: position,
		user:    user,
		storage: storage,
	}, nil
}

func readEntityByUser(storage abstract.Storage, user uint) (Entity, error) {

	position, _, err := storage.GetReference(user)
	if err != nil {
		return Entity{}, err
	}

	fact, _, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		fact:    fact,
		position: position,
		user:    user,
		storage: storage,
	}, nil
}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetFact() uint {

	return e.fact
}

func (e Entity) GetPosition() uint {

	return e.position
}

func (e Entity) GetUser() uint {

	return e.user
}

func (e Entity) PointToClass(class abstractClass.Entity) error {

	return e.storage.Connect(e.fact, class.GetStory())
}

func (e Entity) PointToFact(fact uint) error {

	return e.storage.Connect(e.fact, fact)
}

func (e Entity) HasNextStory() (bool, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}
