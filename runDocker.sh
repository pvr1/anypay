go build ./
docker build -t pervonrosen/anypay:latest ./
#docker tag anypay:latest pervonrosen/anypay:latest
docker push pervonrosen/anypay:latest
docker container run -p 8080:8080 pervonrosen/anypay:latest