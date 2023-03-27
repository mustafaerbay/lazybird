resource "aws_s3_account_public_access_block" "example" {
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}


resource "aws_s3_bucket" "b" {
  bucket = var.bucketname
  tags = {
    Name        = var.bucketname
    Environment = var.environment
  }
  block_public_acls   = aws_s3_account_public_access_block.example.arn
  block_public_policy = aws_s3_account_public_access_block.example.arn
}

# resource "aws_s3_bucket_acl" "example" {
#   bucket = aws_s3_bucket.b.id
#   acl    = "public-read"
# }






# module "s3_bucket" {
#   source = "terraform-aws-modules/s3-bucket/aws"

#   bucket = var.bucketname
#   acl    = "public-read"

#   versioning = {
#     enabled = true
#   }

# }