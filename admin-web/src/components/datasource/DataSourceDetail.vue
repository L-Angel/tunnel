<template>
  <div>
    <Row>
      <Col>
        <Card>

        </Card>
      </Col>
    </Row>
    <Row :style="{marginTop: '10px'}">
      <Col span="24">
        <Card :style="{marginTop:'10px'}">
          <Table :columns="datasource_columns"
                 :data="datasource_origin"
                 size="small"></Table>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
export default {
  name: "DataSourceDetail",
  data() {
    return {
      datasource_origin: [{
        db_name: 'db_shard_01',
        host: '127.0.0.1',
        port: 3301,
        username: 'demo_user_01',
        password: '******'
      }, {
        db_name: 'db_shard_02',
        host: '127.0.0.1',
        port: 3302,
        username: 'demo_user_02',
        password: '******'
      }, {
        db_name: 'db_shard_03',
        host: '127.0.0.1',
        port: 3303,
        username: 'demo_user_03',
        password: '******'
      }, {
        db_name: 'db_shard_04',
        host: '127.0.0.1',
        port: 3304,
        username: 'demo_user_04',
        password: '******'
      }, {
        db_name: 'db_shard_05',
        host: '127.0.0.1',
        port: 3305,
        username: 'demo_user_05',
        password: '******'
      }],
      datasource_columns: [{
        title: 'Db_Name',
        key: 'db_name',

      }, {
        title: 'Host',
        key: 'host'
      }, {
        title: 'Port',
        key: 'port'
      }, {
        title: 'Username',
        key: 'username'
      }, {
        title: 'Password',
        key: 'password'
      }]
    }
  }
}
</script>

<style scoped>

</style>
