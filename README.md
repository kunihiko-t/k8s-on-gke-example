# Simple Kubernetes examples

## Create a cluster

```
gcloud container clusters create example --num-nodes 2
```

## Build a new image

```
cd app
docker build -t asia.gcr.io/<GCP-PROJECT-ID>/app .
```

## Push a Docker image to the private repo

```
gcloud docker -- push asia.gcr.io/<GCP-PROJECT-ID>/app
```

## Create PV, PVC, Deployment, Ingress, Services..

```

kubectl apply -f <each yml file>

```

## Run data migration

```
kubectl get pods

NAME                      READY     STATUS    RESTARTS   AGE
app-3191165209-6h38h      1/1       Running   0          1d
app-3191165209-l8f51      1/1       Running   0          1d
.....
.....

kubectl exec -it app-755235758-9k08z /bin/bash
goose up

```
## You can access to the db


```
kubectl get pods

NAME                     READY     STATUS    RESTARTS   AGE
app-2651697751-87q1f     1/1       Running   0          8m
db-803988798-2kkrf       1/1       Running   0          21s
...
...

kubectl exec -it db-803988798-2kkrf  /bin/bash
psql -U postgres

```


## Access to the web app

```

kubectl get ingress

NAME          HOSTS     ADDRESS        PORTS     AGE
app-ingress   *         xxx.xxx.xxx.xxx   80        1d


Open xxx.xxx.xxx.xxx

```

## Rolling Update

```
Edit some files & rebuild docker image & run gcloud docker -- push command again

And then

kubectl set image deployment/app app=asia.gcr.io/<GCP-PROJECT-ID>/app:latest
```


## Clean up

```
kubectl delete deployment --all
kubectl delete ingress --all
kubectl delete service --all
kubectl delete pv --all
kubectl delete pvc --all
....
....

gcloud container clusters delete example

```