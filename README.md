## install to K8s

install cni-copier to Kubernetes cluster as Daemonset, cni-copier will automatically install the CNI plugins not present on specific path

```
kubectl apply -f copier-daemonset.yml
```

## install on local machine

```
docker run --rm -v /opt/cni/bin:/out -e CNI_BIN_DST=/out -e OVERRIDE=true zengxu/cni-copier:221226-ec76e3c
```

current plugin sources
- https://github.com/containernetworking/plugins
- https://github.com/phosae/cniplugins