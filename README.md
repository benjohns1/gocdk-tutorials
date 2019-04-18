# Go CDK Tutorials
Run through of the Go Cloud Development tutorials https://gocloud.dev/tutorials/

## CLI Upload
Tutorial: https://gocloud.dev/tutorials/cli-uploader/  

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

## Wire Tutorial
Run through of the Wire tutorial: https://github.com/google/wire/blob/master/_tutorial/README.md  
Initial greeter without wire: [fdeb6e5](commit/fdeb6e5c97d3ab66dca1d07e08d7405d4bf02ed6)  
Adding Wire to greeter: [87c8631](commit/87c863180dd9d4c1007207cd34dda04861ef7edd)  