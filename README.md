Tunnel
====
基于binlog的数据同步组件
----

# 功能诉求

- 阶段一
    - 单节点Worker，其他worker backup
    - 基于zk的任务协调与恢复
    - 多mysql 实例监听

- 阶段二
    - 多worker分发

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