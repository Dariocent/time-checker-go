terraform { 
  cloud { 
    
    organization = "REDACTED" 

    workspaces { 
      name = "dev" 
    } 
  } 
  required_version = ">=1.0"

  required_providers {
    azapi = {
      source  = "azure/azapi"
      version = "~>2.0"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>4.7.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~>3.6.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "0.12.0"
    }
  }
}

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id

}

provider "azapi" {

  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id

}

