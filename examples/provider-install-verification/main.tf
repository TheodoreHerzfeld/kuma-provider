terraform {
  required_providers {
    uptime-kuma = {
      source = "local/theodoreherzfeld/uptime-kuma"
    }
    local = {
      source = "local/local"
    }
  }
}

provider "uptime-kuma" {
  username = "admin"
  password = "admin"
  host = "http://192.168.1.163:8000"
}
# 
data "uptime-kuma_user" "testdata" {
  username = "admin"
}

data "uptime-kuma_monitor" "testmonitordata" {
  id = 16
}

resource "local_sensitive_file" "testmonitordata" {
  filename = "./monitor.yml"
  content = yamlencode(data.uptime-kuma_monitor.testmonitordata)
}

output "testoutput" {
  value = data.uptime-kuma_user.testdata
}

# resource "uptime-kuma_monitor" "testresource" {
#   name = "testmon"
#   url = "google.com"
#   type = "http"
# }
# 
# output "test" {
#   value = data.uptime-kuma_user.testdata
# }