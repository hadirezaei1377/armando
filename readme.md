## armando learned me CICD and K8s

minikube version
minikube start --force
minikube status

minikube kubectl -- get pods -A

docker login

docker compose up -d
docker compose down

kubectl create namespace bootcamp
kubectl get namespaces


kubectl --namespace=kube-system get pods
kubectl -n=kube-system get pods

kubectl config set-context --current --namespace=bootcamp
kubectl get pods


kubectl delete -f Deployment.yml
kubectl create -f Deployment.yml

kubectl apply -f Deployment.yml
kubectl get pods
kubectl get deployments
kubectl delete deployment armando-deployment

kubectl get svc
kubectl get services

kubectl create -f Service.yml

ping armando-service
