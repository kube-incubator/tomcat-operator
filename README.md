# tomcat-operator

Tomcat operator enables managing multiple tomcat installments at scale.

## Quick Start

### Controller deploy

```sh
$ git clone https://github.com/kube-incubator/tomcat-operator.git
$ cd tomcat-operator

$ kubectl apply -f deploy/crds/tomcat_v1alpha1_tomcat_crd.yaml

$ kubectl apply -f deploy/service_account.yaml
$ kubectl apply -f deploy/role.yaml
$ kubectl apply -f deploy/role_binding.yaml
$ kubectl apply -f deploy/operator.yaml
```

### Deploying a tomcat cluster

```sh
$ kubectl apply -f deploy/crds/tomcat_v1alpha1_tomcat_cr.yaml
```

```sh
$ kubectl get tomcat
NAME     AGE
tomcat   51m

$ kubectl get pod
NAME                               READY   STATUS    RESTARTS   AGE
tomcat-669459654b-qcsr8            1/1     Running   0          51m
tomcat-669459654b-wzclc            1/1     Running   0          51m
tomcat-operator-65bd87b969-mlrpr   1/1     Running   0          51m

$ kubectl get svc
NAME              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
kubernetes        ClusterIP   10.96.0.1        <none>        443/TCP    69m
tomcat            ClusterIP   10.110.157.42    <none>        80/TCP     51m
tomcat-operator   ClusterIP   10.108.128.161   <none>        8383/TCP   51m
```

Then you can view the web service at http://10.110.157.42/sample/.

![](http://ww1.sinaimg.cn/large/007uElTfly1g1ajks6vewj30q90f1q41.jpg)

## References

- The basic framework used in this project: [operator-sdk](https://github.com/operator-framework/operator-sdk)
- The workflow to build this project: [user-guide](https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md)