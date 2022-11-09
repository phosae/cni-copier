## How to

install cni-copier to Kubernetes cluster as Daemonset, cni-copier will automatically install the CNI plugins not present on every node

```
kubectl apply -f copier-daemonset.yml
```

current plugin sources
- https://github.com/containernetworking/plugins