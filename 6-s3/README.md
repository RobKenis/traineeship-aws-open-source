# S3

## Creating a bucket
Deploy [bucket.yaml](bucket.yaml) using the CloudFormation console.

## Uploading objects
```shell
aws s3api put-object --bucket <bucket> --key test.json --body test.json
```

## Making objects public
Select the object in your S3 bucket, click **actions** and select **Make public** or

```shell
aws s3api put-object --bucket <bucket> --key test.json --body test.json --acl public-read
```

## Hosting static websites
Upload index.html
```shell
aws s3api put-object --bucket <bucket> --key index.html --body index.html --acl public-read --content-type text/html
```
Add some styling
```shell
aws s3api put-object --bucket <bucket> --key main.css --body main.css --acl public-read --content-type text/css
```

## Enable website hosting
The website works, but we have to type `/index.html` at the end of the URL, which is not very nice.
In the bucket under **Properties**, enable **static website hosting** and set **index.html** as index page.
After applying the changes, the console shows a **Bucket website endpoint**, use this to connect to your
website.

## Enable caching
> This CloudFormation stack needs to be deployed in **us-east-1**. 

Deploy [cloudfront.yaml](cloudfront.yaml) using the CloudFormation console.

