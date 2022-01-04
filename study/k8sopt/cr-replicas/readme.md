## k8sOperator 

replicas demo

### 初始化项目

```bash
kubebuilder init --domain github.com --repo github.com/quaintclever/meetlife
```

### 创建API

```bash
# GVK  组, 版本, 种类 (crd) 
kubebuilder create api --group paas --version v1 --kind ReplicasDemo

---

Create Resource [y/n]
y
Create Controller [y/n]
y
```