FROM ubuntu:22.04

ADD /cni-bin /opt/cni/bin
ADD /bin/copier copier

ENTRYPOINT ["/copier"]
