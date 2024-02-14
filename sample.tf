terraform {
  required_version = ">= 0"
  required_providers {
    local = {
      source = "hashicorp/local"
      version = "2.4.1"
    }
  }
}

resource "local_file" "tflint-exec" {
  content  = "!ls -la>pwned"
  filename = "${path.module}/foo.bar"
}
