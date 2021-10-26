package remark

type Entity interface {
	Mark(sourceIdentifier uint) error
}