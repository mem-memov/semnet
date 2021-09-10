package class

import "github.com/mem-memov/semnet/internal/class/node"

type Entity struct {
	bitNode       node.Bit
	characterNode node.Character
	wordNode      node.Word
	phraseNode    node.Phrase
	detailNode    node.Detail
	remarkNode    node.Remark
	factNode      node.Fact
	storyNode     node.Story
	userNode      node.User
}

func newEntity(
	bitNode node.Bit,
	characterNode node.Character,
	wordNode node.Word,
	phraseNode node.Phrase,
	detailNode node.Detail,
	remarkNode node.Remark,
	factNode node.Fact,
	storyNode node.Story,
	userNode node.User,
) Entity {
	return Entity{
		bitNode,
		characterNode,
		wordNode,
		phraseNode,
		detailNode,
		remarkNode,
		factNode,
		storyNode,
		userNode,
	}
}

func (e Entity) CreateBit() (uint, error) {
	return e.bitNode.Create()
}

func (e Entity) IsBit(identifier uint) (bool, error) {
	return e.bitNode.Is(identifier)
}

func (e Entity) GetAllBits() ([]uint, error) {
	return e.bitNode.GetAll()
}

func (e Entity) CreateCharacter() (uint, error) {
	return e.characterNode.Create()
}

func (e Entity) IsCharacter(identifier uint) (bool, error) {
	return e.characterNode.Is(identifier)
}

func (e Entity) CreateWord() (uint, error) {
	return e.wordNode.Create()
}

func (e Entity) IsWord(identifier uint) (bool, error) {
	return e.wordNode.Is(identifier)
}

func (e Entity) CreatePhrase() (uint, error) {
	return e.phraseNode.Create()
}

func (e Entity) IsPhrase(identifier uint) (bool, error) {
	return e.phraseNode.Is(identifier)
}

func (e Entity) CreateDetail() (uint, error) {
	return e.detailNode.Create()
}

func (e Entity) IsDetail(identifier uint) (bool, error) {
	return e.detailNode.Is(identifier)
}

func (e Entity) CreateRemark() (uint, error) {
	return e.remarkNode.Create()
}

func (e Entity) IsRemark(identifier uint) (bool, error) {
	return e.remarkNode.Is(identifier)
}

func (e Entity) CreateFact() (uint, error) {
	return e.factNode.Create()
}

func (e Entity) IsFact(identifier uint) (bool, error) {
	return e.factNode.Is(identifier)
}

func (e Entity) CreateStory() (uint, error) {
	return e.storyNode.Create()
}

func (e Entity) IsStory(identifier uint) (bool, error) {
	return e.storyNode.Is(identifier)
}

func (e Entity) CreateUser() (uint, error) {
	return e.userNode.Create()
}

func (e Entity) IsUser(identifier uint) (bool, error) {
	return e.userNode.Is(identifier)
}
