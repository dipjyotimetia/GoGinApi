#!/bin/bash
@echo "Clean Deploy"
  cd config
  helm delete goginapi
  kubectl delete deployments --all
  kubectl delete services --all
  kubectl delete pods --all
  kubectl delete daemonset --all