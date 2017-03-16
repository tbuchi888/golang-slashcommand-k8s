# golang-slashcommand-k8s
This BOT made by gollang for slash command of mattermost and Slack. 
And it run on Docker and k8s. 

But these are still being verified.
This is Sample program.

## Usage

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

When you get slash command token , you need to change GOBOTTOKEN(yourtoken) on k8s/gobotblue.yml or edit deployment.

Example of After that.

```
kubectl edit deployment gobotblue
```
## on Docker

```
docker build -t tbuchi888/gobotblue:v1 ./ --force-rm=true
docker run --name golang-bot -d -p 3000:3000 tbuchi888/gobotblue:v1 -e GOBOTTOKEN=your_token
```

## golang

```
export GOBOTTOKEN="your_token"
go get -u github.com/favclip/ucon
go install main
./main
```

## Thanks
- https://docs.mattermost.com/developer/slash-commands.html
- https://github.com/favclip/ucon
- https://developers.eure.jp/tech/gae-slashcommands-example/
