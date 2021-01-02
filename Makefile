aws:
	docker-compose up website-aws

aws-build:
	docker-compose up --build website-aws

dev:
	docker-compose up website-dev

dev-build:
	docker-compose up --build website-dev

remove:
	docker-compose rm -f
