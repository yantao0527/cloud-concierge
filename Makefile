run:
	docker-compose up --build cloud-concierge

#CONCIERGE_IMAGE=dragondropcloud/cloud-concierge:latest
#CONCIERGE_IMAGE=dragondropcloud/cloud-concierge-dev:latest
CONCIERGE_IMAGE=cloud-concierge-local:latest

build:
	docker build -t $(CONCIERGE_IMAGE) main

rmi:
	docker rmi $(CONCIERGE_IMAGE)