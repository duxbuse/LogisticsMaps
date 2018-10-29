provider "google" {
  project     = "${var.project}"
  credentials = "${file("diceroller-220503-8497483a16e9.json")}"
  region      = "${var.region}"
  zone        = "${var.zone}"
}

# module "trigger-gcr" {
#   source      = "google-terraform-modules/cloudbuild-gcr/google"
#   repository  = "diceroller"
#   source-repo = "https://github.com/duxbuse/LogisticsMaps.git"

#   triggers = [
#     {
#       branch = "master"
#     },
#   ]
# }

# output "repository_url" {
#   description = "Git Repository URL"
#   value       = "${module.trigger-gcr.repository_url}"
# }
