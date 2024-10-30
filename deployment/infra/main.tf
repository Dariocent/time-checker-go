resource "azurerm_resource_group" "cantella_resource_group" {
    location = "westeurope"
    name     = "DCA-RG"
    tags     = {
        "Owner" = "Dario CANTELLA"
    }
}

# Azure Container Registry
resource "azurerm_container_registry" "cantelladevregistry" {
  name                = "cantelladevregistry"
  resource_group_name = azurerm_resource_group.cantella_resource_group.name
  location            = azurerm_resource_group.cantella_resource_group.location
  sku                 = "Standard"
  admin_enabled       = true
  public_network_access_enabled = true
  anonymous_pull_enabled = true
  tags                = {}
}

resource "azurerm_storage_account" "terraformstateaccount" {
  name                     = "terraformstateeaccount"
  resource_group_name      = azurerm_resource_group.cantella_resource_group.name
  location                 = azurerm_resource_group.cantella_resource_group.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "dev"
  }
}

resource "azurerm_storage_container" "terraformstatecontainer" {
  name                  = "terraformstatecontainer"
  storage_account_name  = azurerm_storage_account.terraformstateaccount.name
  container_access_type = "private"
}