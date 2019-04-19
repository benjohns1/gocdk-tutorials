provider "aws" {
    region     = "us-west-2"
}

resource "aws_instance" "example" {
    ami           = "ami-0b5913cdbba67598e"
    instance_type = "t2.micro"
    provisioner "local-exec" {
        command = "echo ${aws_instance.example.public_ip} > ip_address.txt"
    }
}

resource "aws_eip" "ip" {
    instance = "${aws_instance.example.id}"
}