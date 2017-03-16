# cheatsheet
## before that change your token on gobotblue.yml gobottoken
## dash board
minikube dashboard

## create service
kubectl create -f k8s/mattermost.yml

kubectl create -f k8s/gobotsv.yml

kubectl describe svc

## create deployment
kubectl create -f k8s/gobotblue.yml 

kubectl describe deployment

## get pods
kubectl get pods

## get URL on minikube
minikube service mattermostsv --url

minikube service gobotsv --url
