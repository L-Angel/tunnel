# K8S-Zookeeper

----

## 脚本
```shell
kubectl create -f pv-zk.yaml
# 查看PV
# kubectl get pv

kubectl create -f k8s-zk.yaml
# kubectl get pods

kubectl expose service zk-cs --port=2181 --target-port=2181 --external-ip=192.168.56.105 --name use-zk

```

## Tcp端口转发
* https://www.cnblogs.com/victorbu/p/14780037.html
## K8S 部署zookeeper
* https://blog.csdn.net/wslyk606/article/details/90720424