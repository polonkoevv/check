#!/bin/bash

# Создаем PVC для PostgreSQL
echo "Creating PostgreSQL PVC..."
kubectl apply -f postgres-pvc.yaml

# Применяем Secret для PostgreSQL
echo "Applying PostgreSQL Secret..."
kubectl apply -f postgres-secret.yaml

# Применяем Service для PostgreSQL
echo "Applying PostgreSQL Service..."
kubectl apply -f postgres-service.yaml

# Применяем Deployment для PostgreSQL
echo "Applying PostgreSQL Deployment..."
kubectl apply -f postgres-deployment.yaml

# Ждем, пока PostgreSQL будет готов
echo "Waiting for PostgreSQL to be ready..."
kubectl wait --for=condition=Available deployment/postgres --timeout=300s

# Применяем ConfigMap и Secret для приложения (если они есть)
if [ -f "configmap.yaml" ]; then
    kubectl apply -f configmap.yaml
fi

if [ -f "secret.yaml" ]; then
    kubectl apply -f secret.yaml
fi

# Применяем Service для приложения
echo "Applying Application Service..."
kubectl apply -f service.yaml

# Ждем, пока Service будет создан
echo "Waiting for Application Service to be ready..."
kubectl wait --for=condition=Available service/library-service --timeout=60s

# Применяем Deployment для приложения
echo "Applying Application Deployment..."
kubectl apply -f deployment.yaml

# Ждем, пока все поды будут готовы
echo "Waiting for Application Deployment to be ready..."
kubectl wait --for=condition=Available deployment/library-deployment --timeout=300s

echo "All resources have been applied successfully!"

# Показываем статус
echo "\nCurrent status:"
kubectl get pods,svc,deployment,pvc -l app=library 