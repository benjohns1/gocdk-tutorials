provider "aws" {
    region     = "us-west-2"
}

resource "aws_instance" "example" {
    ami           = "ami-0b5913cdbba67598e"
    instance_type = "t2.micro"
}