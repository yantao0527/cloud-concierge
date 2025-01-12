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

test:
	cd main; go test ./...

pytest:
	cd main/internal; \
		coverage run -m pytest --cov ./  --cov-branch --cov-report term-missing --cov-config=./python_scripts/tests/.coveragerc

clean-main:
	-rm -rf main/repo
	-rm -rf main/state_files
	-rm -rf main/current_cloud
	-rm -rf main/outputs

ci-lint-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.54.2

ci-lint:
	cd main; golangci-lint run

tools-install:
	docker cp concierge:/usr/local/bin/tfswitch /tmp/tfswitch
	docker cp concierge:/usr/local/bin/terraformer /tmp/terraformer
	docker cp concierge:/usr/local/bin/tfsec /tmp/tfsec
	sudo mv /tmp/tfswitch /usr/local/bin/tfswitch
	sudo mv /tmp/terraformer /usr/local/bin/terraformer
	sudo mv /tmp/tfsec /usr/local/bin/tfsec
	sudo chown root:root /usr/local/bin/tfswitch
	sudo chown root:root /usr/local/bin/terraformer
	sudo chown root:root /usr/local/bin/tfsec

python-link:
	sudo ln -s $(PWD)/main/internal/python_scripts /python_scripts
