## AWS Lambda - Go Fiber Application with Serverless Deployment


## References
[AWS Lambda](https://aws.amazon.com/lambda/) - Run code without thinking about servers or clusters (1 million requests free per month with the AWS Free Tier)

[Go Fiber](https://gofiber.io/) as they describe is an Express-inspired web framework written in Go. It is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go.

[Serverless](https://www.serverless.com/) is a framwork to easily deploy your code as Lambda Function on different cloud providers like AWS, Google Cloud etc...

[Aws Lambda Go Api Proxy](https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/README.md) - We have used this package to proxy all request from API Gateway to Fiber App running as lambda function.


## Deploying the sample
1. Please make sure you have serverless package installed globally.
2. Go to the build.sh file, change the `servicemind-prod` with your aws profile name (which you can easily configure with `aws configure --profile name`) 
3. You can then run this file which will build go executable and then deploy it on your aws account in the given region.
