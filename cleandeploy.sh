@echo "Clean Deploy"
    kubectl delete deployments --all
    kubectl delete services --all
    kubectl delete pods --all
    kubectl delete daemonset --all