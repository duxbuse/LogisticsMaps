provider "google" {
  project     = "diceroller-220503"
  credentials = "${file("diceroller-220503-8497483a16e9.json")}"
  region      = "us-central1"
  zone        = "us-central1-c"
}

provider "kubernetes" {
  load_config_file = false

  host = "${google_container_cluster.primary.endpoint}"
  
  cluster_ca_certificate = "${base64decode(google_container_cluster.primary.master_auth.0.cluster_ca_certificate)}"
  config_context_cluster = "${google_container_cluster.primary.name}"
}
