IMAGE:=sklrsn/go-tutorial

all: build run

build:
	  sudo docker build -t $(IMAGE) .

run:
	  sudo docker run -p 8088:8080 $(IMAGE)

clean:
	  sudo docker rmi -f $(IMAGE)
