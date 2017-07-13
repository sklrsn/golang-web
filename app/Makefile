IMAGE:=sklrsn/app

all: build run

build:
	  sudo docker build -t $(IMAGE) .

run:
		docker run -d --name app -p 8080 --link redis:redis $(IMAGE)
