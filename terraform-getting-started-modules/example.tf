provider "aws" {
    region = "us-east-1"
}

module "consul" {
    source = "hashicorp/consul/aws"
    version = "0.3.3"
    aws_region = "us-east-1"
    num_servers = "2"
}

output "consul_server_asg_name" {
  value = "${module.consul.asg_name_servers}"
}