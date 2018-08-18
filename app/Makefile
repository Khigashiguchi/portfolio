build:
	docker build -t profile-site .

run:
	docker run -p 8080:8080 profile-site

deploy:
	docker build --tag=khigashiguchi/profile .
	docker push khigashiguchi/profile