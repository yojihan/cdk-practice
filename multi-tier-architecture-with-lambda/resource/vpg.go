package resource

import (
	"multi-tier-architecture-with-lambda/enum"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type vpcSpec struct {
	Name string `field:"required"`
	CIDR string `field:"required"`
}

var vpcSpecs = map[string]vpcSpec{
	enum.VPC_ID.String(): {
		Name: enum.VPC_NANE.String(),
		CIDR: "10.0.0.0/24",
	},
}

func NewVPC(scope constructs.Construct) map[string]*awsec2.CfnVPC {
	m := map[string]*awsec2.CfnVPC{}

	for id, spec := range vpcSpecs {
		vpc := awsec2.NewCfnVPC(scope, jsii.String(id), &awsec2.CfnVPCProps{
			CidrBlock: jsii.String(spec.CIDR),
			Tags: &[]*awscdk.CfnTag{
				{Key: jsii.String("Name"), Value: jsii.String(spec.Name)},
			},
		})

		m[id] = &vpc
	}

	return m
}
