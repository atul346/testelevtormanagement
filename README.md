# testelevtormanagement

## Getting Started

Follow these steps to run the application locally or in a Kubernetes cluster.

### Local Development

1. Build the Docker image:
   ```bash
   docker build -t elevator-management:latest .
2.Run the Docker Compose:
   docker-compose up
3.Access the application: 

Elevators: http://localhost:8080/elevators
Hotels: http://localhost:8080/hotels

Kubernetes Deployment

1.Apply the Kubernetes deployment:
kubectl apply -f kubernetes/deployment.yaml

2.kubectl get pods -w

Access the application using the assigned service URL.

Dependencies
Gin: Web framework
GORM: ORM library for database interactions

   
