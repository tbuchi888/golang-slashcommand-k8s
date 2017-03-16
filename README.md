# golang-slashcommand-k8s
This is BOT by gollang for slash command of mattermost and Slack. 
And run on Docker and k8s. 
But these are still being verified.

## Usage

```
go get -u github.com/favclip/ucon
go install main
./main
```

or run on Docker Container

```
docker build -t tbuchi888/gobotblue:v1 ./ --force-rm=true
docker run --name golang-bot -d -p 3000:3000 tbuchi888/gobotblue:v1
```

if you use minkube

```
eval $(minikube docker-env)
docker build -t tbuchi888/gobotblue:v1 ./ --force-rm=true

kubectl create -f k8s/mattermost.yml
kubectl create -f k8s/gobotsv.yml
kubectl create -f k8s/gobotblue.yml 

minikube service mattermostsv --url
minikube service gobotsv --url
```

## Thanks
https://docs.mattermost.com/developer/slash-commands.html
https://github.com/favclip/ucon
https://developers.eure.jp/tech/gae-slashcommands-example/
