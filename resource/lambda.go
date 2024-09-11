package resource

import (
	"os"
	"path"

	"github.com/yojihan/cdk-practice/enum"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type lambdaSpec struct {
	id      enum.ResourceId
	name    enum.ResourceName
	maxAge  float64
	retry   float64
	runtime awslambda.Runtime
	handler string
	code    string
}

var lambdaSpecs = map[string]lambdaSpec{
	enum.LAMBDA_ORDER_ID.String(): {
		id:      enum.LAMBDA_ORDER_ID,
		name:    enum.LAMBDA_ORDER_NAME,
		maxAge:  5,
		retry:   0,
		runtime: awslambda.Runtime_NODEJS_18_X(),
		handler: "index.handler",
		code:    "order",
	},
}

func NewLambda(scope constructs.Construct) map[string]*awslambda.Function {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	m := map[string]*awslambda.Function{}

	for id, spec := range lambdaSpecs {
		lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
			FunctionName:  jsii.String(spec.name.String()),
			MaxEventAge:   awscdk.Duration_Minutes(jsii.Number(spec.maxAge)),
			RetryAttempts: jsii.Number(spec.retry),
			Runtime:       spec.runtime,
			Handler:       jsii.String(spec.handler),
			Code:          awslambda.Code_FromAsset(jsii.String(path.Join(dir, "lambda", spec.code)), &awss3assets.AssetOptions{}),
		})

		m[id] = &lambda
	}

	return m
}
