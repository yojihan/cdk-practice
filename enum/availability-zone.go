package enum

type AZName int

const (
	// vpc
	AZ_AP_NORTHEAST_1_A = iota
	AZ_AP_NORTHEAST_1_C
	AZ_AP_NORTHEAST_1_D
)

func (id AZName) String() string {
	switch id {
	case AZ_AP_NORTHEAST_1_A:
		return "ap-northeast-1a"
	case AZ_AP_NORTHEAST_1_C:
		return "ap-northeast-1c"
	case AZ_AP_NORTHEAST_1_D:
		return "ap-northeast-1d"
	default:
		panic("Not supported AZName")
	}
}
