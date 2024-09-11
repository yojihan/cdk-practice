package enum

type ResourceId int

const (
	STACK_ID ResourceId = iota

	// vpc
	VPC_ID

	// Subnet
	SUBNET_1_ID
	SUBNET_2_ID
	SUBNET_3_ID

	// DynamoDB
	DYNAMODB_ID

	// lambda
	LAMBDA_ORDER_ID
	LAMBDA_PUSH_QUEUE_ID
	LAMBDA_SEND_EMAIL_ID

	// API Gateway
	API_GATEWAY_ID
)

func (id ResourceId) String() string {
	switch id {
	case STACK_ID:
		return "MyStack"
	case VPC_ID:
		return "MyVPC"
	case SUBNET_1_ID:
		return "Subnet1"
	case SUBNET_2_ID:
		return "Subnet2"
	case SUBNET_3_ID:
		return "Subnet3"
	case DYNAMODB_ID:
		return "OrderDynamoDB"
	case LAMBDA_ORDER_ID:
		return "LambdaOrder"
	case LAMBDA_PUSH_QUEUE_ID:
		return "LambdaPushQueue"
	case LAMBDA_SEND_EMAIL_ID:
		return "LambdaSendEmail"
	case API_GATEWAY_ID:
		return "MyAPIGateway"
	default:
		panic("Not supported ResourceId")
	}
}
