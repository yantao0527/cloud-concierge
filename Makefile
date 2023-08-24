run:
	docker-compose up --build cloud-concierge

python-dep:
	pip3 install -r main/internal/python_scripts/requirements.txt

#CONCIERGE_IMAGE=dragondropcloud/cloud-concierge:latest
#CONCIERGE_IMAGE=dragondropcloud/cloud-concierge-dev:latest
CONCIERGE_IMAGE=cloud-concierge-local:latest

build:
	docker build -t $(CONCIERGE_IMAGE) main

rmi:
	docker rmi $(CONCIERGE_IMAGE)


build-pytest:
	docker build -f ./main/internal/python_scripts/pytest-dockerfile -t cloud-concierge-pytest:latest main

credentials-aws:
	-mkdir main/credentials
	ln -s $(HOME)/.aws main/credentials/aws
