variable location {
    default = "Australia East"
}

variable "client_id" { 
    default = "a8ac0526-21e2-4736-ad33-db6541449a4e"
}
variable "client_secret" {}

variable "agent_count" {
    default = 2 // because free tier only allows a quota of 4 cores in total
}

variable "ssh_public_key" {
    default = "~/.ssh/id_rsa.pub"
}

variable "dns_prefix" {
    default = "k8stest"
}

variable cluster_name {
    default = "k8stest"
}

variable resource_group_name {
    default = "azure-k8stest"
}



variable log_analytics_workspace_name {
    default = "testLogAnalyticsWorkspaceName"
}

# refer https://azure.microsoft.com/global-infrastructure/services/?products=monitor for log analytics available regions
variable log_analytics_workspace_location {
    default = "australiaeast" // New South Wales, Australia
}

# refer https://azure.microsoft.com/pricing/details/monitor/ for log analytics pricing 
variable log_analytics_workspace_sku {
    default = "PerGB2018"
}

variable vm_size {
    default = "Standard_B2s" // smallest k8s compatible VM type
}