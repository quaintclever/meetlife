## 构建 opa-server
```bash
cd opasrv
make run
```

### 如何获取 bundle.tar.gz

> 推荐参考官方文档, HTTP APIs
https://www.openpolicyagent.org/docs/latest/http-api-authorization/
 
或者编写 rbac.rego 文件
执行下面命令

```bash
opa build rbac.rego
```
 