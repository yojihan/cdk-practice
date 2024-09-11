package resource

import (
	"multi-tier-architecture-with-lambda/enum"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type subnetSpec struct {
	VpcId enum.ResourceId
	Name  enum.ResourceName
	AZ    enum.AZName
	CIDR  string
}

var subnetSpecs = map[string]subnetSpec{
	enum.SUBNET_1_ID.String(): {
		VpcId: enum.VPC_ID,
		Name:  enum.SUBNET_1_NAME,
		AZ:    enum.AZ_AP_NORTHEAST_1_A,
		CIDR:  "10.0.0.0/28",
	},
	enum.SUBNET_2_ID.String(): {
		VpcId: enum.VPC_ID,
		Name:  enum.SUBNET_2_NAME,
		AZ:    enum.AZ_AP_NORTHEAST_1_C,
		CIDR:  "10.0.0.16/28",
	},
	enum.SUBNET_3_ID.String(): {
		VpcId: enum.VPC_ID,
		Name:  enum.SUBNET_3_NAME,
		AZ:    enum.AZ_AP_NORTHEAST_1_D,
		CIDR:  "10.0.0.32/28",
	},
}

type SubnetProps struct {
	VpcMap map[string]*awsec2.CfnVPC
}

func NewSubnet(scope constructs.Construct, props *SubnetProps) map[string]*awsec2.CfnSubnet {
	m := map[string]*awsec2.CfnSubnet{}

	for id, spec := range subnetSpecs {
		subnet := awsec2.NewCfnSubnet(scope, jsii.String(id), &awsec2.CfnSubnetProps{
			VpcId:            (*props.VpcMap[spec.VpcId.String()]).AttrVpcId(),
			AvailabilityZone: jsii.String(spec.AZ.String()),
			CidrBlock:        jsii.String(spec.CIDR),
			Tags: &[]*awscdk.CfnTag{
				{Key: jsii.String("Name"), Value: jsii.String(spec.Name.String())},
			},
		})

		m[id] = &subnet
	}

	return m
}
