variable "subscription_id" {
    description = "Azure subscription ID"
}

variable "client_id" {
    description = "Azure client ID"
}

variable "client_secret" {
    description = "Azure client secret"
    type        = string
}

variable "tenant_id" {
    description = "Azure tenant ID"
    type        = string
}

variable "service_principal_id" {
  description = "The Azure service principal ID."
  type        = string
}

variable "service_principal_key" {
  description = "The Azure service principal key."
  type        = string
}

variable "node_count" {
  description = "Number of nodes in the Kubernetes cluster"
  type        = number
  default = 2
}

variable "username" {
  description = "Admin username for the Kubernetes cluster"
  type        = string
  default     = "cantella"
}