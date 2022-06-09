Tunnel
====
基于binlog的数据同步组件
----

# 产品介绍

# 系统部署架构
![部署架构](./document/tunnel架构.png)
# 产品特性
- cluster-group-worker 三级架构模式，支持集群无限扩容
- master-slave高可用部署架构
- 基于zk的任务协调与恢复



# 注册中心节点设计

```shell 
-------
  |--tunnel
  |     |
  |     |--[cluster_code]
  |     |     |
  |     |     |--[group_code]
  |     |     |     |--election
  |     |     |     |     |(leader IP)  
  |     |     |     |
  |     |     |     |--nodes
  |     |     |     |     |--Node Ip1
  |     |     |     |     |--Node Ip2
  |     |     |     |     |--Node Ip3
  |     |     |     |     
  |     |     |     |--postion
  |     |     |     |     |(byte data，mysql Position数据)
  |     |     |
  |     |     |

```

# 数据模式

* **UPDATE**

```json
{
  "schema": "",
  "table": "xxx",
  "action": "update",
  "rows": [
    [
      {
        "column": "xx",
        "before": "xx",
        "after": "xx"
      },
      {
        "column": "xx2",
        "before": "xx",
        "after": "xx"
      }
    ],
    [
      {
        "column": "xx",
        "before": "xx",
        "after": "xx"
      },
      {
        "column": "xx2",
        "before": "xx",
        "after": "xx"
      }
    ]
  ]
}
```

* **INSERT**
```json
{
	"schema": "",
	"table": "xxx",
	"action": "insert",
	"rows": [
		[{
			"column": "xx",
			"value": "xx"
		}, {
			"column": "xx2",
			"value": "xx"
		}],
		[{
			"column": "xx",
			"value": "xx"
		}, {
			"column": "xx2",
			"value": "xx"
		}]
	]
}
```

* **DELETE**
```json
{
	"schema": "",
	"table": "xxx",
	"action": "delete",
	"rows": [
		[{
			"column": "xx",
			"value": "xx"
		}, {
			"column": "xx2",
			"value": "xx"
		}],
		[{
			"column": "xx",
			"value": "xx"
		}, {
			"column": "xx2",
			"value": "xx"
		}]
	]
}
```