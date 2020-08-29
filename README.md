# consumejson

#Build Docker Image__
docker build -t consumejson .

#Push Registry__
docker tag consumejson:latest <registry_host>/consumejson:latest
docker push <registry_host>/consumejson:latest

#Deploy To K8s Cluster__
