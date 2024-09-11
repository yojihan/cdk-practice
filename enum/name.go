package enum

type ResourceName int

const (
	// vpc
	VPC_NANE ResourceName = iota

	// subnets
	SUBNET_1_NAME
	SUBNET_2_NAME
	SUBNET_3_NAME

	// dynamoDB
	DYNAMODB_NAME

	// lambda
	LAMBDA_ORDER_NAME
	LAMBDA_PUSH_QUEUE_NAME
	LAMBDA_SEND_EMAIL_NAME
)

func (id ResourceName) String() string {
	switch id {
	case VPC_NANE:
		return "multi-tier-vpc"
	case SUBNET_1_NAME:
		return "multi-tier-subnet-1"
	case SUBNET_2_NAME:
		return "multi-tier-subnet-2"
	case SUBNET_3_NAME:
		return "multi-tier-subnet-3"
	case DYNAMODB_NAME:
		return "order-db"
	case LAMBDA_ORDER_NAME:
		return "order-lambda"
	case LAMBDA_PUSH_QUEUE_NAME:
		return "push-queue-lambda"
	case LAMBDA_SEND_EMAIL_NAME:
		return "send-email-lambda"
	default:
		panic("Not supported ResourceName")
	}
}
