# lambda-EFS-sqlite example
Using sqlite+EFS as the persistent layer in lambda application to acheive zero-maintance high avaialble web service

## Quick start
1. following the [guide](https://aws.amazon.com/blogs/compute/using-amazon-efs-for-aws-lambda-in-your-serverless-applications/) to create a lambda function and EFS, and mount EFS path to lambda

2. build and deploy binary
```bash
./deploy.sh
```

3. test the binary using lambda web console