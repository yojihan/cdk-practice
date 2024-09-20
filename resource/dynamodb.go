package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/yojihan/cdk-practice/enum"
)

type dynamodbSpec struct {
	TableName      string
	ReadCapacity   float64
	WriteCapacity  float64
	StreamViewType awsdynamodb.StreamViewType
	Attributes     []*awsdynamodb.CfnTable_AttributeDefinitionProperty
	KeySchema      []*awsdynamodb.CfnTable_KeySchemaProperty
}

var dynamodbSpecs = map[string]dynamodbSpec{
	enum.DYNAMODB_ID.String(): {
		TableName:      "Order",
		ReadCapacity:   2,
		WriteCapacity:  2,
		StreamViewType: awsdynamodb.StreamViewType_NEW_IMAGE,
		Attributes: []*awsdynamodb.CfnTable_AttributeDefinitionProperty{
			{
				AttributeName: jsii.String("OrderId"),
				AttributeType: jsii.String("S"),
			},
			{
				AttributeName: jsii.String("OrderNo"),
				AttributeType: jsii.String("S"),
			},
		},
		KeySchema: []*awsdynamodb.CfnTable_KeySchemaProperty{
			{
				AttributeName: jsii.String("OrderId"),
				KeyType:       jsii.String("HASH"),
			},
			{
				AttributeName: jsii.String("OrderNo"),
				KeyType:       jsii.String("RANGE"),
			},
		},
	},
}

func NewDynamoDB(scope constructs.Construct) map[string]*awsdynamodb.CfnTable {
	m := map[string]*awsdynamodb.CfnTable{}

	for id, spec := range dynamodbSpecs {
		table := awsdynamodb.NewCfnTable(scope, jsii.String(id), &awsdynamodb.CfnTableProps{
			TableName: jsii.String(spec.TableName),
			ProvisionedThroughput: &awsdynamodb.CfnTable_ProvisionedThroughputProperty{
				ReadCapacityUnits:  jsii.Number(spec.ReadCapacity),
				WriteCapacityUnits: jsii.Number(spec.WriteCapacity),
			},
			StreamSpecification: &awsdynamodb.CfnTable_StreamSpecificationProperty{
				StreamViewType: jsii.String(string(spec.StreamViewType)),
			},
			AttributeDefinitions: spec.Attributes,
			KeySchema:            spec.KeySchema,
		})

		m[id] = &table
	}

	return m
}
