########## consumejson



CI with JENKIN PIPELINE

Jenkins is a open source CI/CD tool i.e. continuous integration/continuous delivery tool. It is used to integrate and automate your product development and testing processes. The purpose of using this tool was to build and test project continuously. In an ascent of agile, this could help developers to integrate the changes to the project as quickly as possible and obtain fresh builds ready for testing. Easy to learn and maintain, Free with tons of stackoverflow hint.

#Build Docker Image

git clone https://github.com/truongnq87/consumejson.git

cd consumejson

docker build -t consumejson .

#Push Registry  

docker tag consumejson:latest <registry_host>/consumejson:latest  

docker push <registry_host>/consumejson:latest

###########################################################

CD with ARGOCD

Why Argo CD

Application definitions, configurations, and environments should be declarative and version controlled. Application deployment and lifecycle management should be automated, auditable, and easy to understand.

https://argoproj.github.io/argo-cd/

#Deploy To K8s Cluster

#Pull image from <registry_host>

#Change image name into <registry_host>/consumejson:latest in k8s deployment manifest  

kubectl  apply -f k8s_deployment.yaml

########## CI/CD Diagram


#Tobe loadbalance we set replica=n (n is base on require)

kubectl scale deploy consumejson-deployment --replicas=$n


By theory, rolling update keep our services handle zero-downtime upgrades, but in fact, we can tunning it to adapt the productivity.

https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/


![alt text](https://github.com/truongnq87/consumejson/blob/master/Diagram.png?raw=true)
