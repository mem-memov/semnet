package fact

type Fact interface {
	GetStory() (Story, error)

	AddRemark(
		storyReferences []Story,
		factReferences []Fact,
		remarkReferences []Remark,
		object string,
		property string,
	) (Remark, error)
}
