#!/bin/bash


export PROJECT_ID=sonar-status-test
gcloud projects create $PROJECT_ID

gcloud projects list

gcloud iam service-accounts \
    create devops-catalog \
    --project $PROJECT_ID \
    --display-name devops-catalog


gcloud iam service-accounts list \
    --project $PROJECT_ID

gcloud iam service-accounts \
    keys create account.json \
    --iam-account devops-catalog@$PROJECT_ID.iam.gserviceaccount.com \
    --project $PROJECT_ID

gcloud iam service-accounts \
    keys list \
    --iam-account devops-catalog@$PROJECT_ID.iam.gserviceaccount.com \
    --project $PROJECT_ID


gcloud projects \
    add-iam-policy-binding $PROJECT_ID \
    --member serviceAccount:devops-catalog@$PROJECT_ID.iam.gserviceaccount.com \
    --role roles/owner

export TF_VAR_project_id=$PROJECT_ID


terraform init

terraform apply

# Enable billing for this project 
open https://console.cloud.google.com/storage/browser?project=sonar-status-test

# Generate unique bucket name
export TF_VAR_state_bucket=doc-$(date +%Y%m%d%H%M%S)

export BUCKET_NAME=doc-$(date +%Y%m%d%H%M%S)

# Update backend.tf with the bucket name
cat backend.tf \
  | sed -e "s@devops-catalog@$TF_VAR_state_bucket@g" \
  | tee backend.tf

terraform apply


# K8S cluster

terraform apply

gcloud container get-server-config \
    --region us-east1 \
    --project $PROJECT_ID

export K8S_VERSION=[...]

terraform apply \
    --var k8s_version=$K8S_VERSION

export KUBECONFIG=$PWD/kubeconfig

gcloud container clusters \
    get-credentials \
    $(terraform output cluster_name) \
    --project \
    $(terraform output project_id) \
    --region \
    $(terraform output region)

kubectl create clusterrolebinding \
    cluster-admin-binding \
    --clusterrole \
    cluster-admin \
    --user \
    $(gcloud config get-value account)

kubectl get nodes


