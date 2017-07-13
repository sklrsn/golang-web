IMAGE:=sklrsn/nginx

all: build run

build:
	  sudo docker build -t $(IMAGE) .

run:
		docker run -d --name nginx -p 80:80 --link app:app $(IMAGE)
