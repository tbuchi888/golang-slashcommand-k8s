# golang-slashcommand-k8s
- This BOT made by gollang for slash command of mattermost and Slack. And it run on Docker and k8s. 
- But these are still being verified.
- This is Sample program.

## before that
- You need to set up for slash command of mattermost or Slack. 
- And, When you get slash command token , you need to change or Set Env GOBOTTOKEN(Your_slashcommand_token).
  - k8s: k8s/gobotblue.yml or edit deployment,
  - Docker: docker run -e option
  - golang: export

## Usage example
### on k8s / minikube

```
# use Docker on minikube
eval $(minikube docker-env)

# create docker image
docker build -t tbuchi888/gobotblue:v1 ./ --force-rm=true

# create k8s service and deploymnet
kubectl create -f k8s/gobotsv.yml
kubectl create -f k8s/gobotblue.yml 

# get url of service on minikube 
minikube service gobotsv --url

# below mattermost is for testing
kubectl create -f k8s/mattermost.yml
minikube service mattermostsv --url
```

When you get slash command token , you need to change GOBOTTOKEN(Your_slashcommand_token) on k8s/gobotblue.yml or edit deployment.

```
kubectl edit deployment gobotblue
```
## on Docker

```
docker build -t tbuchi888/gobotblue:v1 ./ --force-rm=true
docker run --name golang-bot -d -p 3000:3000 -e GOBOTTOKEN=Your_slashcommand_token tbuchi888/gobotblue:v1
```

## golang

```
export GOBOTTOKEN="Your_slashcommand_token"
go get -u github.com/favclip/ucon
go install main
./main
```

## Thanks
- https://docs.mattermost.com/developer/slash-commands.html
- https://github.com/favclip/ucon
- https://developers.eure.jp/tech/gae-slashcommands-example/
