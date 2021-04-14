#!/bin/sh

kubectl delete -f deployment_all_in_one.yaml
kubectl delete secret preset-pod -n preset