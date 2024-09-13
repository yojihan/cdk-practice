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
				AttributeName: jsii.String("OrderId"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("UserId"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("UserName"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("ProductName"),
				AttributeType: jsii.String("S"),
			},
			&awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("Qty"),
				AttributeType: jsii.String("N"),
			},
		},
		KeySchema: []interface{}{
			&awsdynamodb.CfnTable_KeySchemaProperty{
				AttributeName: jsii.String("OrderId"),
				KeyType:       jsii.String("HASH"),
			},
			&awsdynamodb.CfnTable_KeySchemaProperty{
				AttributeName: jsii.String("UserId"),
				KeyType:       jsii.String("RANGE"),
			},
		},
		LocalSecondaryIndexes: []interface{}{
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_OrderId_"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("OrderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("UserId"),
						KeyType:       jsii.String("RANGE"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_OrderId_"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("OrderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("UserId"),
						KeyType:       jsii.String("RANGE"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_OrderId_UserName"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("OrderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("UserName"),
						KeyType:       jsii.String("RANGE"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_OrderId_ProductName"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("OrderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("ProductName"),
						KeyType:       jsii.String("RANGE"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
			&awsdynamodb.CfnTable_LocalSecondaryIndexProperty{
				IndexName: jsii.String("index_OrderId_Qty"),
				KeySchema: []interface{}{
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("OrderId"),
						KeyType:       jsii.String("HASH"),
					},
					&awsdynamodb.CfnTable_KeySchemaProperty{
						AttributeName: jsii.String("Qty"),
						KeyType:       jsii.String("RANGE"),
					},
				},
				Projection: awsdynamodb.ProjectionType_ALL,
			},
		},
	})

}
