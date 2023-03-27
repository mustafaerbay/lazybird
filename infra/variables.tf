

variable "tag" {
  type        = string
  description = "prefix for AWS resources"
  default     = "lazybird"
}

variable "environment" {
  type        = string
  description = "environment information"
  default     = "prod"
}

variable "region" {
  type        = string
  description = "The AWS region to create resources in."
  default     = "ap-southeast-1"
}

variable "bucketname" {
  type        = string
  description = "The AWS region to create resources in."
  default     = "lazybird1133"
}