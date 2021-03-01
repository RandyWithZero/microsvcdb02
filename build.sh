image="168.1.9.1/test/microsvcdb02:latest"
docker build -t $image .
docker push $image