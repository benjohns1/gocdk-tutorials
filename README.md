# Go CDK Tutorials
Run through of the Go Cloud Development tutorials https://gocloud.dev/tutorials/

## CLI Uploader
CLI Uploader tutorial: https://gocloud.dev/tutorials/cli-uploader/  

Run it:  
`cd upload`  
`go build`  
### AWS bucket storage
Create AWS access key & S3 bucket, set local env vars:
- AWS_REGION
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
 
`./upload s3://bucket-name testfile.txt`

### GCP storage
Create GCP service account & storage bucket, set local env var to point to credential JSON file:
 - GOOGLE_APPLICATION_CREDENTIALS
  
`./upload gs://bucket-name testfile.txt`

### Azure bucket storage
Create Azure storage blob container, set local env vars:
 - AZURE_STORAGE_ACCOUNT
 - AZURE_STORAGE_KEY

`./upload azblob://bucket-name testfile.txt`

## Wire Dependency Injection
Wire tutorial: https://github.com/google/wire/blob/master/_tutorial/README.md  
1. Initial greeter without wire: [fdeb6e5](https://github.com/benjohns1/gocdk-tutorials/commit/fdeb6e5c97d3ab66dca1d07e08d7405d4bf02ed6)  
2. Adding Wire to greeter: [87c8631](https://github.com/benjohns1/gocdk-tutorials/commit/87c863180dd9d4c1007207cd34dda04861ef7edd)  
3. Grumpy greeter: [069366a](https://github.com/benjohns1/gocdk-tutorials/commit/069366a1c33fa731349139fbf314251f12243f55)
4. Pass phrase parameter: [d815cbd](https://github.com/benjohns1/gocdk-tutorials/commit/d815cbd76868b21735336efa4802df5acb2720b8)  

Run it:  
`cd wiretut`  
`go build`  
`./wiretut`