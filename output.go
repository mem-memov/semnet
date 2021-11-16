package semnet

import "bytes"

type Output struct {
	graph *Graph
}

func NewOutput(graph *Graph) *Output {

	return &Output{
		graph: graph,
	}
}

func (o *Output) GetStory(remark Remark) (string, error) {

	factRemark, err := remark.GetFirstFact()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	for {
		object, property, err := factRemark.GetObjectAndProperty()
		if err != nil {
			return "", err
		}

		buffer.WriteString(object)
		buffer.WriteString(" ")
		buffer.WriteString(property)
		buffer.WriteString(" ")
	}

	return buffer.String(), nil
}
