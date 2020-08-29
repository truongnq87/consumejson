# consumejson

#Build Docker Image
docker build -t consumejson .

#Push Registry
docker tag consumejson:latest <registry_host>/consumejson:latest
docker push <registry_host>/consumejson:latest

#Deploy To K8s Cluster
