IMAGE:=sklrsn/golang-rest-app

all: build run

build:
	  sudo docker build -t $(IMAGE) .

run:
	  sudo docker run $(IMAGE) /bin/sh -c "go run main/*"

clean:
	  sudo docker rmi -f $(IMAGE)
