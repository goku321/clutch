package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	dynamodbv1 "github.com/lyft/clutch/backend/api/aws/dynamodb/v1"
)

func (c *client) DescribeTable(ctx context.Context, region string, tableName string) (*dynamodbv1.Table, error) {
	cl, err := c.getRegionalClient(region)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.DescribeTableInput{TableName: aws.String(tableName)}
	result, err := cl.dynamodb.DescribeTable(ctx, input)
	if err != nil {
		return nil, err
	}

	currentCapacity := &dynamodbv1.ProvisionedThroughput{
		WriteCapacityUnits: aws.ToInt64(result.Table.ProvisionedThroughput.WriteCapacityUnits),
		ReadCapacityUnits:  aws.ToInt64(result.Table.ProvisionedThroughput.ReadCapacityUnits),
	}

	globalSecondaryIndexes := getGlobalSecondaryIndexes(result.Table.GlobalSecondaryIndexes)

	var ret = &dynamodbv1.Table{
		Name:                   aws.ToString(result.Table.TableName),
		Region:                 region,
		GlobalSecondaryIndexes: globalSecondaryIndexes,
		ProvisionedThroughput:  currentCapacity,
	}
	return ret, nil

}

func getGlobalSecondaryIndexes(indexes []ddbtypes.GlobalSecondaryIndexDescription) []*dynamodbv1.GlobalSecondaryIndex {
	gsis := make([]*dynamodbv1.GlobalSecondaryIndex, len(indexes))
	for idx, i := range indexes {
		gsis[idx] = newProtoForGlobalSecondaryIndex(i)
	}
	return gsis
}

func newProtoForTable(table ddbtypes.TableDescription) *dynamodbv1.Table {
	currentCapacity := &dynamodbv1.ProvisionedThroughput{
		ReadCapacityUnits:  aws.ToInt64(table.ProvisionedThroughput.ReadCapacityUnits),
		WriteCapacityUnits: aws.ToInt64(table.ProvisionedThroughput.WriteCapacityUnits),
	}

	return &dynamodbv1.Table{
		Name:                  aws.ToString(table.TableName),
		ProvisionedThroughput: currentCapacity,
	}
}

func newProtoForGlobalSecondaryIndex(index ddbtypes.GlobalSecondaryIndexDescription) *dynamodbv1.GlobalSecondaryIndex {
	currentCapacity := &dynamodbv1.ProvisionedThroughput{
		ReadCapacityUnits:  aws.ToInt64(index.ProvisionedThroughput.ReadCapacityUnits),
		WriteCapacityUnits: aws.ToInt64(index.ProvisionedThroughput.WriteCapacityUnits),
	}

	return &dynamodbv1.GlobalSecondaryIndex{
		Name:                  aws.ToString(index.IndexName),
		ProvisionedThroughput: currentCapacity,
	}
}

func (c *client) UpdateTableCapacity(ctx context.Context, region string, tableName string, targetTableRcu int32, targetTableWcu int32) {
	// TO DO: implement
}

func (c *client) UpdateGSICapacity(ctx context.Context, region string, tableName string, indexName string, targetIndexRcu int32, targetIndexWcu int32) {
	// TO DO: implement
}
