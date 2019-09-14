# Dig as a Service!

## Manual Deployment on CloudRun
```bash
export GCP_PROJECT=gdg-devfest-auburn

# Build Docker image with cloudbuild
gcloud builds submit --tag gcr.io/${GCP_PROJECT}/dig

# Deployment to CloudRun
gcloud beta run deploy dig --image gcr.io/${GCP_PROJECT}/dig --platform managed --allow-unauthenticated
```
