## install to K8s

install cni-copier to Kubernetes cluster as Daemonset, cni-copier will automatically install the CNI plugins not present on specific path

```
kubectl apply -f copier-daemonset.yml
```

## install on local machine

```
docker run --rm -v /<your>/<host>/<path>:/out -e CNI_BIN_DST=/out zengxu/cni-copier:ac86731
```

current plugin sources
- https://github.com/containernetworking/plugins