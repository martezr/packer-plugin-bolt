packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
    //    bolt = {
    //      source  = "github.com/martezr/bolt"
    //      version = ">=1.0.0"
    //    }
  }
}

source "docker" "ubuntu" {
  image  = "ubuntu:xenial"
  commit = true
}

build {
  sources = [
    "source.docker.ubuntu"
  ]
  provisioner "bolt" {
    backend   = "ssh"
    user      = "root"
    bolt_task = "facts"
  }
}
