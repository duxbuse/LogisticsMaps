variable "project" {
  default = "diceroller-220503"
}

variable "region" {
  default = "us-central1"
}

variable "zone" {
  default = "us-central1-c"
}

variable "repository" {
  type        = "string"
  description = "Name of mirror repository on GCP"
  default     = "diceroller"
}

variable "triggers" {
  type = "list"

  default = [{
    branch = "master"
  }]

  description = "Options of trigger"
}

variable "source-repo" {
  type        = "string"
  description = "Name of source repository on github or bitbucket"
  default     = "https://github.com/duxbuse/LogisticsMaps.git"
}
