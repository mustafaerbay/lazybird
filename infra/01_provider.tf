# locals {
#   envs = { for tuple in regexall("(.*)=(.*)", file("../.env")) : tuple[0] => sensitive(tuple[1]) }
# }


provider "aws" {
  region  = var.region
  profile = "lazybird"
}