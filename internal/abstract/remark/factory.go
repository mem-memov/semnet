package remark

type Factory interface {

	CreateExistingEntity(
		classIdentifier uint,
		detailIdentifier uint,
		remarkIdentifier uint,
		factIdentifier uint,
	) Entity

	CreateNewEntity(
		detailIdentifier uint,
		remarkIdentifier uint,
		factIdentifier uint,
	) (Entity, error)
}
