# provider "kubernetes" {
#   load_config_file = false
#    host = "${google_container_cluster.primary.endpoint}"
  
#   cluster_ca_certificate = "${base64decode(google_container_cluster.primary.master_auth.0.cluster_ca_certificate)}"
#   config_context_cluster = "${google_container_cluster.primary.name}"
# }