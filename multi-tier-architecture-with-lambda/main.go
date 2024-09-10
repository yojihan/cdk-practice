package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	awscdk.NewStack(app, jsii.String("MyStack"), &awscdk.StackProps{
		Env: env(),
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("AWS_DEFAULT_ACCOUNT "),
		Region:  jsii.String("AWS_DEFAULT_REGION"),
	}
}
