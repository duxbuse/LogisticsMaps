data "google_project" "project" {}

data "local_file" "cluster-name"{
  filename = "./../terraform-data/cluster-name.tfdata"
}
data "local_file" "deployment-name"{
  filename = "./../terraform-data/deployment-name.tfdata"
}

resource "google_sourcerepo_repository" "new_git_repository" {
  name = "${var.repository}"
}

resource "google_cloudbuild_trigger" "new_git_build_trigger" {
  count       = "${length(var.triggers)}"
  description = "Trigger Git repository ${var.repository} / ${lookup(var.triggers[count.index], "branch", "master")}"

  trigger_template {
    project     = "${data.google_project.project.project_id}"
    branch_name = "${lookup(var.triggers[count.index], "branch", "master")}"
    repo_name   = "${var.repository}"
  }

  substitutions {
    _TAG             = "${lookup(var.triggers[count.index], "branch", "master")}"
    _COMPUTE_ZONE = "${var.zone}"
    _CLUSTER = "${data.local_file.cluster-name.content}"
  }

  filename = "cloudbuild.yaml"

  # Google Git repository has been created.
  depends_on = [
    "google_sourcerepo_repository.new_git_repository",
  ]
}
