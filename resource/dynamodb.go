package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewDynamoDB(scope constructs.Construct) {
	awsdynamodb.NewCfnTable(scope, jsii.String("OrderDynamoDB"), &awsdynamodb.CfnTableProps{
		TableName: jsii.String("order"),
		ProvisionedThroughput: &awsdynamodb.CfnTable_ProvisionedThroughputProperty{
			ReadCapacityUnits:  jsii.Number(2),
			WriteCapacityUnits: jsii.Number(2),
		},
		AttributeDefinitions: []interface{}{
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("orderId"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("userId"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("userName"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("productName"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("qty"),
				AttributeType: jsii.String("N"),
			},
		},
		KeySchema: []interface{}{
			&awsdynamodb.CfnTable_KeySchemaProperty{
				AttributeName: jsii.String("orderId"),
				KeyType:       jsii.String("HASH"),
			},
		},
		LocalSecondaryIndexes: []interface{}{
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_orderId_userId"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("orderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("userId"),
						KeyType:       jsii.String("HASH"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_orderId_userName"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("orderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("userName"),
						KeyType:       jsii.String("HASH"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_orderId_productName"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("orderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("productName"),
						KeyType:       jsii.String("HASH"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_orderId_qty"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("orderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("qty"),
						KeyType:       jsii.String("HASH"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
		},
	})

}
