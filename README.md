# Example 1

Create namespace:

```bash
kubectl create namespace backend
```
Sve kubernetes fajlove pokrenuti da dobijemo configmap, secret i mongo service i statefulSet
```shell
kubectl -n backend apply -f mongo-configmap.yml
kubectl -n backend apply -f mongo-secret.yml
kubectl -n backend apply -f mongo.yml
kubectl -n backend apply -f search-configmap.yml 
kubectl apply -f search-service.yml 
```
Testing load balancing and service:
```shell
kubectl -n backend run -it --rm  --image curlimages/curl:8.00.1 curl -- sh
```
Inside the container execute `curl http://hotel:8083/hotels` (hotel jer je to naziv servisa)
```shell
/ $ curl http://hotel:8083/hotels
```

Testing connection from another namespace:
```shell
kubectl create namespace tmp
kubectl -n demo run -it --rm  --image curlimages/curl:8.00.1 curl -- sh

/ $ curl http://counter.demo.svc.cluter.local:8000
Counter:  16
```

Ingress setup:
Deploy ingress:
```shell
kubectl -n demo apply -f ingress.yaml
```