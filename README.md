# File Uploader Via APIGateway and Lambda-S3
This is small exercise to upload files to s3 using APIGateway and lambda function.
This supports for small file uploads. 
Size upload supported by aws max-
API Gateway - 10MB 
lambda function - 6MB.

This works for binary file uploads.

The gateway API is configured to upload file and content type set to application/octet-stream.
For more help read API Gateway configureation for binaray data
https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-payload-encodings-configure-with-console.html

For more help refer below link
https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html


#How to build a package for lambda function

~GOOS=linux go build main.go

~zip file-handler.zip main

If need any help? shoot me a mail.
