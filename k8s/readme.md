# cheatsheet

## dash board
minikube dashboard

## create service
kubectl create -f mattermost.yml
kubectl create -f k8s/sv.yml

kubectl describe svc

## create deployment
kubectl create -f k8s/bl.yml 

kubectl create -f k8s/gr.yml 

kubectl describe deployment

## get pods
kubectl get pods

## get URL on minikube
minikube service mattermostsv --url
minikube service gobotdemo --url
