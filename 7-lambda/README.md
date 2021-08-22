# Lambda

## Creating a Lambda
Deploy [lambda.yaml](lambda.yaml) using the CloudFormation console. This piece of CloudFormation uses
the [Serverless Application Model](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-specification.html) which is an extension to
CloudFormation. Besides creating the Lambda function, it will also create the required IAM roles and API
Gateway to connect to the Lambda using HTTP.
