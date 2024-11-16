terraform {
  required_providers {
    uptime-kuma = {
      source = "hashicorp.com/redted/uptime-kuma"
    }
  }
}

provider "uptime-kuma" {
  username = "admin"
  password = "admin"
  host = "http://192.168.1.163:8000"
}

data "uptime-kuma_user" "testdata" {
  username = "admin"
}

resource "uptime-kuma_monitor" "testresource" {
  name = "testmon"
}

output "test" {
  value = data.uptime-kuma_user.testdata
}