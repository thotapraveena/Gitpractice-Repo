resource "aws_instance" "linux_instance"
ami               = "ami-id"
instance_type     = "t2.micro"
key_name          = "my-key"
subnet_id         = "subnet-linux"
security_groups   = ["ssh"]

provisioner "file" {
    source      ="local_file.txt"
    destination ="/home/ec2-user/remote_file.txt"

    connection {
        type        = "ssh"
        user        = "ec2-user"
        private_key = file("~/.ssh/my-key.pem")
        host        = self.public_ip
       }
}

resource "aws_instance" "ubuntu"
ami               = "ami-id"
instance_type     = "t2.micro"
key_name          = "my-key"
subnet_id         = "subnet-linux"
security_groups   = ["ssh"]
}