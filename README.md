# Kubernetes App for AUI

## INSTALLATION GUIDE

  1. **helm repo add my-app http://utheman.github.io/ExampleDockerImage/**
  
  2. **helm install my-app/my-app --generate-name**
  
  3. On second terminal run command: **minikube tunnel**
  
## RUN Horizontal Pod Autoscaler

  1. Check if metrics-server is enabled **minikube addons list**
  
  2. If not, run command: **minikube addons enable metrics-server** and wait couple of minutes
