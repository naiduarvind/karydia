## Overview of Features provided by Karydia
### Invert the Pod Defaults
The following screenshot shows the pod specification without the usage of Karydia:
* Service accout token is mounted
* A user is not specified (so the pod uses root by default)
* No seccomp profile is assigned
```
kubectl run -it --rm --restart=Never alpine --image=alpine sh -n demo
kubectl edit  pod/alpine -n demo
```
![](../images/pod-without-karydia.png)

If you create a pod after the installation of Karydia, the pod description is different, even if you use the same commands:
* No service account token is mounted
* A user is specified (the root user is not used)
* Privilege escalation is denied
* The seccomp profile runtime/default is assigned

```
kubectl run -it --rm --restart=Never alpine --image=alpine sh -n demo
kubectl edit  pod/alpine -n demo
```
![](../images/pod-with-karydia-and-container-sec-context.png)

### Add a Network Policy
Karydia adds one or multiple default network policies to each namespace and reconciles it.
```
kubectl get networkpolicy -n demo
```
![](../images/networkpolicy.png)

## Features and Configuration Options
You can configure each feature to meet the needs of your applications:
* A custom seccomp profile
* One or multiple default network policies
* A specific network policy per namespace
* The usage of a root user if necessary
* The privilege escalation container option

See all [features and options](../features.md).
