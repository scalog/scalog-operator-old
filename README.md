# scalog-operator

### Running the operator

There are two ways to run the Kubernetes operator, in cluster and locally. Running the operator locally is much faster than deploying it into the cluster when developing. 

##### Local Deploy

There are a few prerequisites to running the kubernetes operator properly.

Ensure that your local VM is running...
```minikube start```

Inside of the scalog/scalog repository, bootstrap the namespace and manually provision volume space for the data layer...
```kubectl create -f data/k8s/namespace.yaml```
```kubectl create -f data/k8s/rbac.yaml```
```kubectl create -f data/k8s/volumes.yaml```

In the scalog-operator repository, register the custom resource with kubernetes...
```kubectl create -f deploy/crds/scalog_v1alpha_scalogservice_crd.yaml```

Run the operator locally...
```operator-sdk up local --namespace=scalog```

Create the scalogservice 
```kubectl create -f deploy/crds/scalog_v1alpha_scalogservice_cr.yaml```

At this point, one shard with 2 replicas should have been spun up in your kubernetes cluster. You can scale up the number of shards by modifying 
`scalog_v1alpha_scalogservice_cr.yaml` and then running `kubectl apply -f deploy/crds/scalog_v1alpha_scalogservice_cr.yaml`. 

##### In Cluster Deploy

This is typically used for production deployments of operators, so no.