const aws = require('aws-sdk');
const dynamo = new aws.DynamoDB.DocumentClient();

export const handler = async (event, context) => {
    const res = await order(event);
    return res;
}

const order = async (event) => {
    // process order

    return null;
}