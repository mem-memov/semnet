package semnet

type Detail struct {
	object   string
	property string
}

func newDetail(object string, property string) Detail {
	return Detail{
		object:   object,
		property: property,
	}
}

func (d Detail) GetObject() string {
	return d.object
}

func (d Detail) GetProperty() string {
	return d.property
}
