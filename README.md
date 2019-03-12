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
tomcat   50m

$ kubectl get pod
NAME                               READY   STATUS    RESTARTS   AGE
tomcat-669459654b-qcsr8            1/1     Running   0          51m
tomcat-669459654b-wzclc            1/1     Running   0          51m
tomcat-operator-65bd87b969-mlrpr   1/1     Running   0          51m

$ kubectl get svc
NAME              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
kubernetes        ClusterIP   10.96.0.1        <none>        443/TCP    54m
tomcat            ClusterIP   10.109.176.183   <none>        80/TCP     52m
tomcat-operator   ClusterIP   10.106.176.219   <none>        8383/TCP   52m
```

## References

- The basic framework used in this project: [operator-sdk](https://github.com/operator-framework/operator-sdk)
- The workflow to build this project: [user-guide](https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md)