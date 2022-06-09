<template>
  <div>
    <TunnelTable :data="datasources.origin" :columns="datasources.columns"></TunnelTable>
  </div>
</template>

<script>
import TunnelTable from "../common/TunnelTable";

export default {
  name: "DataSourceList",
  components: {TunnelTable},
  data() {
    return {
      datasources: {
        columns: [],
        origin: [],
      }
    }
  },
  created() {
    this.datasources.columns = [{
      title: 'ID',
      key: 'id'
    }, {
      title: '名称',
      key: 'name'
    }, {
      title: '描述',
      key: 'desc'
    }, {
      title: '负责人',
      key: 'owner'
    }, {
      title: '状态',
      key: 'status_name'
    }, {
      title: '操作',
      key: 'action',
      render: (h, params) => {
        return h('div', [
          h('Button', {
            props: {
              type: 'primary',
              size: 'small'
            },
            style: {
              marginRight: '5px'
            },
            on: {
              click: () => {
                this.$router.push({name: 'db_datasource_detail', params: {id: params.id}})
              }
            }
          }, '明细')]);
      }
    }
    ];
  },
  mounted() {
    this.datasources.origin = [{
      id: 1,
      name: "Name1",
      status_name: '正常'
    }, {
      id: 2,
      name: "Name1",
      status_name: "部分异常"
    }, {
      id: 1,
      name: "Name1"
    }, {
      id: 3,
      name: "Name1"
    }, {
      id: 4,
      name: "Name1"
    }, {
      id: 1,
      name: "Name1"
    }, {
      id: 6,
      name: "Name1"
    }, {
      id: 7,
      name: "Name1"
    }, {
      id: 1,
      name: "Name1"
    }, {
      id: 8,
      name: "Name1"
    }, {
      id: 1,
      name: "Name1"
    }, {
      id: 12,
      name: "Name1"
    }, {
      id: 11,
      name: "Name1"
    }, {
      id: 10,
      name: "Name2"
    }, {
      id: 9,
      name: "Name1"
    }
    ];
  },
  methods: {}
}
</script>

<style scoped>

</style>
