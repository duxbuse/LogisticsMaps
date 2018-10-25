resource "google_container_cluster" "primary" {
  name                     = "diceroller-cluster"
  remove_default_node_pool = true


  master_auth {
    username = "abc"
    password = "12345678910111213"
  }

  node_pool {
    name = "default-pool"

    management {
      auto_repair  = true
      auto_upgrade = true
    }

    autoscaling {
      min_node_count = 2
      max_node_count = 5
    }
  }
}

# data "google_container_cluster" "primary-data" {
#   name = "diceroller-cluster"
# }

resource "google_container_node_pool" "primary_pool" {
  name       = "primary-pool"
  cluster    = "${google_container_cluster.primary.name}"
  node_count = "2"

  node_config {
    machine_type = "f1-micro"
    preemptible = true
  }
}

# resource "kubernetes_namespace" "diceroller-namespace-resource" {
#   metadata {
#     name = "diceroller-namespace"
#   }
# }

resource "kubernetes_pod" "nginx" {
  metadata {
    name      = "nginx-example"
    # namespace = "${kubernetes_namespace.diceroller-namespace-resource.metadata.name}"

    labels {
      App = "nginx"
    }
  }

  spec {
    container {
      image = "nginx:1.7.8"
      name  = "example"

      port {
        container_port = 80
      }
    }
  }
}

resource "kubernetes_service" "nginx" {
  metadata {
    name = "nginx-example"
  }

  spec {
    selector {
      App = "${kubernetes_pod.nginx.metadata.0.labels.App}"
    }

    port {
      port        = 80
      target_port = 80
    }

    type = "LoadBalancer"
  }
}
