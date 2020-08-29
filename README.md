# consumejson

#Build Docker Image  
docker build -t consumejson .

#Push Registry  
docker tag consumejson:latest <registry_host>/consumejson:latest  
docker push <registry_host>/consumejson:latest

#Deploy To K8s Cluster  
Change image name into <registry_host>/consumejson:latest in k8s deployment manifest  
kubectl  apply -f k8s_deployment.yaml
