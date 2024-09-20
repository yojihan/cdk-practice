package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewDynamoDB(scope constructs.Construct) {
	awsdynamodb.NewCfnTable(scope, jsii.String("OrderDynamoDB"), &awsdynamodb.CfnTableProps{
		TableName: jsii.String("Order"),
		ProvisionedThroughput: &awsdynamodb.CfnTable_ProvisionedThroughputProperty{
			ReadCapacityUnits:  jsii.Number(2),
			WriteCapacityUnits: jsii.Number(2),
		},
		AttributeDefinitions: []interface{}{
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("OrderId"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("OrderNo"),
				AttributeType: jsii.String("S"),
			},
		},
		KeySchema: []interface{}{
			&awsdynamodb.CfnTable_KeySchemaProperty{
				AttributeName: jsii.String("OrderId"),
				KeyType:       jsii.String("HASH"),
			},
			&awsdynamodb.CfnTable_KeySchemaProperty{
				AttributeName: jsii.String("OrderNo"),
				KeyType:       jsii.String("RANGE"),
			},
		},
	})

}
