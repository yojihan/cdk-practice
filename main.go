package main

import (
	"github.com/yojihan/cdk-practice/enum"
	"github.com/yojihan/cdk-practice/resource"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stack := awscdk.NewStack(app, jsii.String(enum.STACK_ID.String()), &awscdk.StackProps{
		Env: env(),
	})

	// VPC
	vpcMap := resource.NewVPC(stack)

	// Subnets
	resource.NewSubnet(stack, &resource.SubnetProps{VpcMap: vpcMap})

	// AWS Lambda
	resource.NewLambda(stack)

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
