# Cloud-Native DevSecOps Project — Acetlisto

A production-grade Go web application built as part of a hands-on DevSecOps course at Hochschule Heilbronn. The project demonstrates a full cloud-native software delivery lifecycle — from code quality and security scanning to containerisation and automated deployment on Google Cloud Run.

## Tech Stack

| Area | Technologies |
|---|---|
| **Application** | Go (Golang), Gorilla Mux, REST API |
| **Observability** | OpenTelemetry, Google Cloud Operations (tracing & spans) |
| **CI/CD** | GitHub Actions (multi-stage pipelines) |
| **Code Quality** | GolangCI-Lint, `go test`, test coverage reporting |
| **Security** | Dependency security scanning, GHCR image scanning |
| **Containerisation** | Docker, GitHub Container Registry (GHCR) |
| **Cloud Deployment** | Google Cloud Run |

## CI/CD Pipeline

The pipeline runs automatically on every pull request and push to `main`:

```
Pull Request
  └── Code Verification Stage
        ├── GolangCI-Lint (fast feedback)
        ├── go test (unit test suite)
        └── Code coverage report

Merge to Main
  └── Build & Security Stage
        ├── Build Docker image
        ├── Security scan (dependencies + image)
        ├── Push to GitHub Container Registry (GHCR)
        └── Deploy to Google Cloud Run
```

## Exercises Completed

- **Exercise 1** — CI pipeline with `go test`, GolangCI-Lint, and test coverage measurement
- **Exercise 2** — Security scanner integrated into pipeline; Docker image built and pushed to GHCR on merge
- **Exercise 3** — Automated deployment to Google Cloud Run using `deploy-cloudrun` GitHub Action
- **Exercise 4** — OpenTelemetry integration with trace export to Google Cloud Operations; custom spans added to API handlers

## Key Features

- Automated test suite runs on every pull request — merge blocked if tests fail
- Code coverage tracked and reported on every PR
- Security scanning on every merge to `main` (dependency analysis + container image scan)
- Fully automated build → scan → push → deploy pipeline
- Distributed tracing with OpenTelemetry for all HTTP requests

## Project Structure

```
.
├── cmd/              # Application entrypoint
├── handlers/         # HTTP request handlers (with OTel spans)
├── .github/
│   └── workflows/    # GitHub Actions CI/CD pipeline definitions
├── .devcontainer/    # Development container configuration
└── telemetry.go      # OpenTelemetry setup and configuration
```

## Running Locally

```bash
git clone https://github.com/Jayeshkhairnar-git/cloud-devops-project.git
cd cloud-devops-project
go run ./cmd
```

## Course Context

Built as part of the **DevSecOps** module at Hochschule Heilbronn (Winter Semester 2025/26), covering CI/CD pipelines, container security, cloud deployment, and observability in a hands-on project environment.


## Related Cloud Computing Projects

As part of my Cloud Computing coursework at Hochschule Heilbronn, 
I completed 5 hands-on exercises covering AWS and GCP:

| Exercise | Technologies |
|---|---|
| AWS S3 Static Website + IAM + CloudTrail | AWS, S3, IAM |
| AWS Lambda + API Gateway REST API | AWS Lambda, Python |
| GCP Serverless App + Firestore | Cloud Functions, Firestore, Python |
| GCP Data Pipeline + Kubernetes (GKE) | Pub/Sub, GKE Autopilot, Docker |
| Infrastructure as Code with Terraform | Terraform, Cloud Run, VPC, IAM |
