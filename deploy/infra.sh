#!/bin/bash
# grafana
helm install grafana grafana-labs/grafana --set persistence.enabled=true --set service.type=NodePort --set service.nodePort=30002
# loki, promtail
helm install loki grafana-labs/loki-stack