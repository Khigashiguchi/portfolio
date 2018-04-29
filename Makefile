build:
	docker build -t profile-site .

run:
	docker run -p 8080:8080 profile-site

deploy:
	docker build --tag=Khigashiguchi/profile-site .
	docker push Khigashiguchi/profile-site