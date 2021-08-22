# Putting it all together

Make an application that does the following: When a photo of a famous person is uploaded to S3, trigger
a process that recognizes the person in the photo and sends an e-mail to your personal address.

## Possible solution
- Trigger a Lambda on an S3 event, documentation can be found [here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-resource-function.html) and 
[here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-s3.html).
- In the Lambda, use the [Rekognition](https://aws.amazon.com/rekognition/) API to extract the person from
the image. Documentation can be found [here](https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/rekognition.html#Rekognition.Client.recognize_celebrities).
- The **event** parameter in Lambda contains enough information about the S3 object to pass to Rekognition.
- Publish a message on an SNS topic connected to your e-mail. Documentation can be found [here](https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/sns.html#SNS.Client.publish). You can set
the **subject** to 'Image recognized of <person>' and include a downloadable image in the body.

Bonus points if you use the container approach for Lambda and push the images to ECR.

## Other possible solution
Create an ECS service that polls the contents of an S3 bucket
