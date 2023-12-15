variable "aws_region" {
  type        = string
  default     = "ap-northeast-1"
  description = "AWS Region"
}

variable "node_size" {
  description = "Size of total nodes"
  default     = 2
}

variable "loadtest_dir_source" {
  default = "plan/"
}

variable "locust_plan_filename" {
  default = "basic.py"
}

variable "subnet_name" {
  default     = "simesime-subnet-public1-ap-northeast-1a"
  description = "Subnet name"
}

variable "ssh_export_pem" {
  description = "Export private ssh key"
  type        = bool
  default     = false
}
