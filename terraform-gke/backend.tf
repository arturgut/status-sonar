terraform {
  backend "gcs" {
    bucket      = "doc-20201005095538"
    prefix      = "terraform/state"
    credentials = "account.json"
  }
}
