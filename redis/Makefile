IMAGE:=redis

all: build run

build:
	  sudo docker build -t $(IMAGE) .

run:
	  docker run -d --name redis -p 6379:6379 $(IMAGE)
