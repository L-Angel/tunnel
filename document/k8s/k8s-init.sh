kubectl create -f pv-zk.yaml
kubectl get pv
kubectl create -f k8s-zk.yaml
kubectl get pods
for i in 0 1 2; do kubectl exec zk-$i zkServer.sh status; done
