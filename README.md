## install to K8s

install cni-copier to Kubernetes cluster as Daemonset, cni-copier will automatically install the CNI plugins not present on specific path

```
kubectl apply -f https://raw.githubusercontent.com/phosae/cni-copier/main/copier-daemonset.yml
```

## install on local machine

```
docker run --rm -v /opt/cni/bin:/out -e CNI_BIN_DST=/out -e OVERRIDE=true zengxu/cni-copier:230425
```

current plugin sources
- https://github.com/containernetworking/plugins, commit 10b5639
- https://github.com/phosae/cniplugins, commit 5a70664
