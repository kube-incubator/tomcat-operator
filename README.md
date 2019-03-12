# tomcat-operator

## Workflow

The following workflow is for a new Kubernetes operator:
1. Create a new operator project using the SDK Command Line Interface(CLI)
2. Define new resource APIs by adding Custom Resource Definitions(CRD)
3. Define Controllers to watch and reconcile resources
4. Write the reconciling logic for your Controller using the SDK and controller-runtime APIs
5. Use the SDK CLI to build and generate the operator deployment manifests

## Quick Start

```sh
$ git clone https://github.com/kube-incubator/tomcat-operator.git
$ cd tomcat-operator

$ kubectl apply -f deploy/crds/tomcat_v1alpha1_tomcat_crd.yaml

$ kubectl apply -f deploy/service_account.yaml
$ kubectl apply -f deploy/role.yaml
$ kubectl apply -f deploy/role_binding.yaml
$ kubectl apply -f deploy/operator.yaml

$ kubectl apply -f deploy/crds/tomcat_v1alpha1_tomcat_cr.yaml
```

```sh
$ kubectl get tomcat
NAME     AGE
tomcat   50m

$ kubectl get pod
NAME                               READY   STATUS    RESTARTS   AGE
tomcat-669459654b-qcsr8            1/1     Running   1          51m
tomcat-669459654b-wzclc            1/1     Running   1          51m
tomcat-operator-65bd87b969-mlrpr   1/1     Running   3          51m

$ kubectl get svc
NAME              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
kubernetes        ClusterIP   10.96.0.1        <none>        443/TCP    54m
tomcat            ClusterIP   10.109.176.183   <none>        80/TCP     52m
tomcat-operator   ClusterIP   10.106.176.219   <none>        8383/TCP   52m
```