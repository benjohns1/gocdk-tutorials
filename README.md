# Go CDK Tutorials
Run through of the Go Cloud Development tutorials https://gocloud.dev/tutorials/

## CLI Upload
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