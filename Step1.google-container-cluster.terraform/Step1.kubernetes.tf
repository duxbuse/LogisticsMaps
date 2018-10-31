# resource "kubernetes_pod" "nginx" {
#   metadata {
#     name      = "${var.deployment-name}"
#      labels {
#       App = "${var.deployment-name}"
#     }
#   }
#    spec {
#     container {
#       name  = "${var.container-name}"
#        port {
#         container_port = 9000
#       }
#     }
#   }
# }

#  resource "kubernetes_service" "nginx" {
#   metadata {
#     name = "${var.deployment-name}-service"
#   }
#    spec {
#     selector {
#       App = "${kubernetes_pod.nginx.metadata.0.labels.App}"
#     }
#      port {
#       port        = 80
#       target_port = 9000
#     }
#      type = "LoadBalancer"
#   }
# }

# # outputting the container name so that Step 2 can read it.
# resource "local_file" "loadbalancer-ip" {
  
#   content = "${kubernetes_service.nginx.load_balancer_ingress.0.ip}"
#   filename = "./../terraform-data/loadbalancer-ip.tfdata"
# }
# resource "local_file" "container-name" {
  
#   content = "${var.container-name}"
#   filename = "./../terraform-data/container-name.tfdata"
# }
# resource "local_file" "deployment-name" {
#   content = "${var.deployment-name}"
#   filename = "./../terraform-data/deployment-name.tfdata"
# }