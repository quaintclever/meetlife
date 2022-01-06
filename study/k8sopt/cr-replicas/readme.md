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
// ReplicasDemoSpec defines the desired state of ReplicasDemo
type ReplicasDemoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// do quaint spec define
	// BatchSize 用来控制 replicas 启动时每次分批的数量
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	BatchSize int32 `json:"batchSize,omitempty"`

	// DeploymentSpec  k8s 自带 deploymentSpec
	DeploymentSpec *v1.DeploymentSpec `json:"deploymentSpec,omitempty"`
}
```

replicasdemo_controller.go
```go
func (r *ReplicasDemoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	cr := &paasv1.ReplicasDemo{}
	if err := r.Get(ctx, req.NamespacedName, cr); err != nil {
		logger.Error(err, fmt.Sprintf("cr get fail!, crName: %s, namespace: %s", cr.Name, cr.Namespace))
		return ctrl.Result{}, err
	}
	logger.Info("perform reconcile", "name", cr.Name, "namespace", cr.Namespace)

	// 新建 deployment
	if cr.Spec.DeploymentSpec != nil {
		deployment := &corev1.Deployment{}
		deployment.APIVersion = cr.APIVersion
		deployment.Kind = "Deployment"
		deployment.ObjectMeta = cr.ObjectMeta
		deployment.Spec = *cr.Spec.DeploymentSpec
		if err := r.Create(ctx, deployment); err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}
```

### 尝试运行它
```bash
# 确保目录正确
% cd cr-replicas
% make run

# 你会发现这里会出现一个错误.
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