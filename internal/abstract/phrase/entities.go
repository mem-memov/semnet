package phrase

type Entities interface {
	Create(
		classIdentifier uint,
		wordIdentifier uint,
		phraseIdentifier uint,
		detailIdentifier uint,
	) Entity

	CreateAndAddClass(
		wordIdentifier uint,
		phraseIdentifier uint,
		detailIdentifier uint,
	) (Entity, error)

	CreateWithDetail(detailIdentifier uint) (Entity, error)
}
