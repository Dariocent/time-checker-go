resource "azurerm_kubernetes_cluster" "cantella-k8s" {
  location            = "francecentral"
  name                = "cantella-k8s"
  resource_group_name = azurerm_resource_group.cantella_resource_group.name
  dns_prefix          = "dns-cantella-k8s"

  identity {
    type = "SystemAssigned"
  }

  default_node_pool {
    name       = "agentpool"
    vm_size    = "Standard_B2s"
    node_count = var.node_count
  }
  linux_profile {
    admin_username = var.username

    ssh_key {
      key_data = azapi_resource_action.ssh_public_key_gen.output.publicKey
    }
  }
  network_profile {
    network_plugin    = "kubenet"
    load_balancer_sku = "standard"
  }
  
}

output "client_certificate" {
  value     = azurerm_kubernetes_cluster.cantella-k8s.kube_config[0].client_certificate
  sensitive = true
}

output "kube_config" {
  value = azurerm_kubernetes_cluster.cantella-k8s.kube_config_raw

  sensitive = true
}