terraform {
  required_providers {
    lambdalabs = {
      source = "WillBeebe/lambdalabs"
    }
  }
}

provider "lambdalabs" {
  # via LAMBDALABS_API_KEY
  # api_key = "your_api_key_here"
}

resource "lambdalabs_ssh_key" "main" {
  name = "mr-keyyy"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCkCVfPukKm1F2xXMNdaCi+rT3mmWWTWaJMUP8Qrq1BE7OndVWYbqxn8mQrz0vXX8eKZFCxJZHmpH/IG1B0nKJ9aU/Qzn2uFpqWZO9zNd5XBZaT4F8MoD/Pv7Xw2Wwh7Fs8/1vCaQu6XdNAB2nfe42la1uMt5g6GUehc7XKCpmxcseDP/zlmvhA6qrBUOh1/4u/kC5i3If/6s12j8qU50FzfvjjSZS095/2dfoZHCtj8FMidIaNihkyp8FNzQ071assAMRi3q3BClSICxXPZwC8DnSXWn4S9OJ4m55biMBYm79t3xi+2Yknj1Nz90YpNwbdmqaaTIJY09hJ003tiJDT mr-keyyy"
}

resource "lambdalabs_instance" "main2" {
  name               = "test-instance2"
  region_name        = "us-west-1"
  instance_type_name = "gpu_1x_a10"
  ssh_key_names      = [lambdalabs_ssh_key.main.name]
}

data "lambdalabs_file_systems" "main" {
}

output "file_systems" {
  value = data.lambdalabs_file_systems.main.file_systems
}

data "lambdalabs_instance_types" "main" {
}

output "instance_types" {
  value = data.lambdalabs_instance_types.main
}

