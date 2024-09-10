package enum

type ResourceName int

const (
	VPC_NANE ResourceName = iota
)

func (id ResourceName) String() string {
	switch id {
	case VPC_NANE:
		return "multi-tier-vpc"
	default:
		panic("Not supported ResourceName")
	}
}
