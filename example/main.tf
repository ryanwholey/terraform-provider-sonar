provider "sonar" {}

data "sonar_workflow" "workflow" {
  name = "default"
}

output "default_workflow" {
  value = data.sonar_workflow.workflow
}
