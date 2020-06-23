# rbacdemo
## casbin+gorm访问鉴权


最简单，最基本的model是ACL,如下面类似的

```
# Request definition
[request_definition]
r = sub, obj, act

# Policy definition
[policy_definition]
p = sub, obj, act

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))

# Matchers
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act

```

一个策略示例针对上面的model,如下:
```
p, alice, data1, read
p, bob, data2, write
```

这意味着：
- alice可以读取data1
- bob可以写data2

有关RBAC的Model示例可以查看本项目中的配置文件`.rbac/rbac.conf`

增加了用户和角色的关联定义,在对应的policy文件中的rule就是：

`g,zhangsan,admin`


使用数据连接存储Policy，需要的注意事项：
- 在数据库连接时，指定的dbName参数,`gormadapter`软件包增加了`casbin`后缀，这个需要特别注意一下
- 在修改policy之后，需要从新loadPolicy一下，防止新增的，或者删除的rule不生效



如下为数据库中`casbin_rule`数据库表的示例：


| p_type | v0 | v1| v2| v3 | v4 | v5|
|------|------|-----|------|----|------|---|
| p | admin | /member/* | GET | | | |
| g | zhaojunwei | admin | | | | |
| g | gaolin | vip | | | | | 
| p | goalin | /member/* | GET | | | |