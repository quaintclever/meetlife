## k8sOperator 

**replicas demo** k8sOperator 控制 replicas demo

执行下面命令时, 首先进入到代码目录
```bash
% cd study/k8sopt/cr-replicas
```

注: `bash` 代码块, `%` 后面跟的是终端需要输入的命令, 没有 `%` 表示是终端输出的内容

### 初始化项目

```bash
% kubebuilder init --domain github.com --repo github.com/quaintclever/meetlife
```

### 创建API

```bash
# GVK  组, 版本, 种类 (crd) 
% kubebuilder create api --group paas --version v1 --kind ReplicasDemo

---

Create Resource [y/n]
y
Create Controller [y/n]
y
```

### CRUD -> 创建一个deployment

replicasdemo_types.go
```go
package v1

import (
	v1 "k8s.io/api/apps/v1"
)

// ReplicasDemoSpec defines the desired state of ReplicasDemo
type ReplicasDemoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// BatchSize 用来控制 replicas 启动时每次分批的数量
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	BatchSize int32 `json:"batchSize,omitempty"`

	// DeploymentSpec  k8s 自带 deploymentSpec
	DeploymentSpec v1.DeploymentSpec `json:"deploymentSpec,omitempty"`
}

// ReplicasDemoStatus defines the observed state of ReplicasDemo
type ReplicasDemoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Ready        string `json:"ready,omitempty"`
	CurrentBatch int    `json:"currentBatch,omitempty"`
}
```

replicasdemo_controller.go
```go
// Reconcile
func (r *ReplicasDemoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)
    
    // TODO(user): your logic here
    // make install 之后, 获取 cr
    cr := &paasv1.ReplicasDemo{}
    if err := r.Get(ctx, req.NamespacedName, cr); err != nil {
        logger.Error(err, fmt.Sprintf("cr get fail!, crName: %s, namespace: %s", cr.Name, cr.Namespace))
        return ctrl.Result{}, err
    }
    logger.Info("======= Perform Reconcile =======", "name", cr.Name, "namespace", cr.Namespace)
    
    // 查询 deployment 是否存在, 不存在则创建
    foundDeployment := &corev1.Deployment{}
    err := r.Get(ctx, types.NamespacedName{Name: cr.Name, Namespace: cr.Namespace}, foundDeployment)
    if err != nil && errors.IsNotFound(err) {
        // 新建 deployment
        deployment := &corev1.Deployment{}
        deployment.APIVersion = cr.APIVersion
        deployment.Kind = "Deployment"
        deployment.ObjectMeta.Name = cr.ObjectMeta.Name
        deployment.ObjectMeta.Namespace = cr.ObjectMeta.Namespace
        deployment.Spec = cr.Spec.DeploymentSpec
        var initReplicas int32 = 0
        deployment.Spec.Replicas = &initReplicas
        // 如果一批次小于总 replicas, 启动一批次的 pod
        // 设置 ownerReferences
        if cr.Spec.BatchSize < *cr.Spec.DeploymentSpec.Replicas {
            *deployment.Spec.Replicas = cr.Spec.BatchSize
        }
        trueVal := true
        falseVal := false
        deployment.OwnerReferences = append(deployment.OwnerReferences, v1.OwnerReference{
        APIVersion:         cr.APIVersion,
        Kind:               cr.Kind,
        Name:               cr.Name,
        UID:                cr.UID,
        Controller:         &trueVal,
        BlockOwnerDeletion: &falseVal,
        })
        logger.Info("======= Creating Deployment =======", "deployment", cr.Name)
        if err := r.Create(ctx, deployment); err != nil {
            return ctrl.Result{}, err
        }
    
        cr.Status.CurrentBatch = 1
    } else if err == nil {
        // 如果deploy 里 ready 的数量 等于总数量, 根据 cr, 修改 deployment 状态
        if *foundDeployment.Spec.Replicas != *cr.Spec.DeploymentSpec.Replicas &&
        foundDeployment.Status.ReadyReplicas == *foundDeployment.Spec.Replicas {
            // 下一批次 到顶了.
            if *foundDeployment.Spec.Replicas+cr.Spec.BatchSize >= *cr.Spec.DeploymentSpec.Replicas {
            *foundDeployment.Spec.Replicas = *cr.Spec.DeploymentSpec.Replicas
            } else {
            *foundDeployment.Spec.Replicas = *foundDeployment.Spec.Replicas + cr.Spec.BatchSize
            }
            logger.Info("======= Updating Deployment =======", "deployment", cr.Name)
            err = r.Update(ctx, foundDeployment)
            // 修改cr 状态
            cr.Status.CurrentBatch = cr.Status.CurrentBatch + 1
        }
    } else {
        logger.Error(err, "======= Reconcile Error =======")
        return ctrl.Result{}, err
    }
    cr.Status.Ready = fmt.Sprintf("%d/%d", foundDeployment.Status.ReadyReplicas, foundDeployment.Status.Replicas)
    logger.Info("======= Update CR Status =======", "crName:", cr.Name)
    err = r.Status().Update(ctx, cr)
    return ctrl.Result{
        // 10s 之后再次调用
        RequeueAfter: time.Second * 10,
    }, nil
}
```

### 尝试运行它
```bash
# 确保目录正确
% cd cr-replicas
% make run

# 你会发现这里会出现一个错误.
main.go:34:2: package git.ddxq.mobi/cs-oss-internal/k8s-operator-example/cr-replicas/api/v1 imports k8s.io/api/apps/v1 from implicitly required module; to add missing requirements, run:
        go get k8s.io/api@v0.22.1
# 执行 下面命令, 检查go mod 是否生成
% go get k8s.io/api@v0.22.1

```

### 接着, 你会出现第二个错误
```bash
error: if kind is a CRD, it should be installed before calling Start
```

### 安装 crd 到 k8s 集群中
```bash
# 确保目录正确
% cd cr-replicas
% make install
...
...
customresourcedefinition.apiextensions.k8s.io/replicasdemoes.paas.github.com created

# 查看crd
% kubectl get crd | grep "replicas"
replicasdemoes.paas.github.com                    2022-01-04T09:43:37Z
```

### 再次尝试运行它
```bash
# 确保目录正确
% cd cr-replicas
% make run

# 运行成功会有以下提示
go fmt ./...
go vet ./...
go run ./main.go
I0105 11:12:35.994117    7019 request.go:665] Waited for 1.030986356s due to client-side throttling, not priority and fairness, request: GET:https://cls-4il4gkgn.ccs.tencent-cloud.com/apis/scheduling.k8s.io/v1?timeout=32s
2022-01-05T11:12:36.308+0800    INFO    controller-runtime.metrics      metrics server is starting to listen    {"addr": ":8080"}
2022-01-05T11:12:36.309+0800    INFO    setup   starting manager
2022-01-05T11:12:36.309+0800    INFO    starting metrics server {"path": "/metrics"}
2022-01-05T11:12:36.309+0800    INFO    controller.replicasdemo Starting EventSource    {"reconciler group": "paas.github.com", "reconciler kind": "ReplicasDemo", "source": "kind source: /, Kind="}
2022-01-05T11:12:36.309+0800    INFO    controller.replicasdemo Starting Controller     {"reconciler group": "paas.github.com", "reconciler kind": "ReplicasDemo"}
2022-01-05T11:12:36.413+0800    INFO    controller.replicasdemo Starting workers        {"reconciler group": "paas.github.com", "reconciler kind": "ReplicasDemo", "worker count": 1}
2022-01-05T11:12:36.414+0800    INFO    controller.replicasdemo perform reconcile       {"reconciler group": "paas.github.com", "reconciler kind": "ReplicasDemo", "name": "replicasdemo-sample", "namespace": "default", "name": "replicasdemo-sample", "namespace": "default"}
```

### 查看 命名空间 demo 里面的情况
```bash
% kubectl get po -n quaint
No resources found in demo namespace.

% kubectl get rs -n quaint  
No resources found in demo namespace.

% kubectl get deploy -n quaint                               
No resources found in demo namespace.
```

### 部署 samples 测试

修改 config/samples 下的 paas_v1_replicasdemo.yaml 文件

paas_v1_replicasdemo.yaml
```yaml
apiVersion: paas.github.com/v1
kind: ReplicasDemo
metadata:
  name: replicasdemo-sample
  namespace: quaint
spec:
  # TODO(user): Add fields here
  batchSize: 1
  deploymentSpec:
    replicas: 3
    selector:
      matchLabels:
        app: replicasdemo-sample
    template:
      metadata:
        labels:
          app: replicasdemo-sample
      spec:
        containers:
          - name: java-tomcat
            image: tomcat:8.5.43-jdk8-openjdk
```

```bash
kubectl apply -f config/samples/paas_v1_replicasdemo.yaml
replicasdemo.paas.ddmc-inc.com/replicasdemo-sample created
```


```bash
# 可以看到 每隔10s 启动了一个.一共启动了3个.
% kubectl get po -n quaint    
NAME                                   READY   STATUS    RESTARTS   AGE
replicasdemo-sample-65c75d5b68-dghtl   1/1     Running   0          42s
replicasdemo-sample-65c75d5b68-nsw55   1/1     Running   0          52s
replicasdemo-sample-65c75d5b68-qlpg8   1/1     Running   0          32s

% kubectl get deploy -n quaint
NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
replicasdemo-sample   3/3     3            3           39s

% kubectl get rs -n demo
NAME                             DESIRED   CURRENT   READY   AGE
replicasdemo-sample-65c75d5b68   3         3         3       24s

% kubectl get crd | grep "replicas"
replicasdemoes.paas.github.com                    2022-01-07T04:24:42Z

% kubectl get replicasdemoes.paas.github.com -n quaint
NAME                  READY   CURRENTBATCH
replicasdemo-sample   3/3     3
```