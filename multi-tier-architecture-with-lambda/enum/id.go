package enum

type ResourceId int

const (
	STACK_ID ResourceId = iota
	VPC_ID
)

func (id ResourceId) String() string {
	switch id {
	case STACK_ID:
		return "MyStack"
	case VPC_ID:
		return "MyVPC"
	default:
		panic("Not supported ResourceId")
	}
}
