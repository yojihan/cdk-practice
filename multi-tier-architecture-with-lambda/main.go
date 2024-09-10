package main

import (
	"multi-tier-architecture-with-lambda/enum"
	"multi-tier-architecture-with-lambda/resource"
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
	resource.NewVPC(stack)

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
