# Go CDK Tutorials
Run through of the Go Cloud Development tutorials https://gocloud.dev/tutorials/

## CLI Upload
### AWS bucket storage
Create AWS access key & S3 bucket, set local env vars:
- AWS_REGION
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

`cd upload`  
`go build`  
`upload s3://bucket-name testfile.txt`  