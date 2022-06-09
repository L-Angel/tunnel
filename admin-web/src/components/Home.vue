<style scoped>
.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
  height: inherit;
}

.layout-logo {
  width: 100px;
  height: 30px;
  background: #5b6270;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 15px;
  left: 20px;
  color: #f5f7f9;
  line-height: 30px;
  font-size: larger;
  font-weight: bolder;
}

.layout-nav {
  width: 420px;
  margin: 0 auto;
  margin-right: 20px;
}
</style>
<template>
  <div class="layout">
    <Layout style="display: flex">
      <Header>
        <Menu mode="horizontal" theme="dark" active-name="1">
          <div class="layout-logo">Tunnel</div>
          <div class="layout-nav">
            <!--            <MenuItem name="1">-->
            <!--              <Icon type="ios-navigate"></Icon>-->
            <!--              Item 1-->
            <!--            </MenuItem>-->
            <!--            <MenuItem name="2">-->
            <!--              <Icon type="ios-keypad"></Icon>-->
            <!--              Item 2-->
            <!--            </MenuItem>-->
            <!--            <MenuItem name="3">-->
            <!--              <Icon type="ios-analytics"></Icon>-->
            <!--              Item 3-->
            <!--            </MenuItem>-->
            <!--            <MenuItem name="4">-->
            <!--              <Icon type="ios-paper"></Icon>-->
            <!--              Item 4-->
            <!--            </MenuItem>-->
          </div>
        </Menu>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff',height:'100%'}">
          <Menu active-name="1-2" theme="light" width="auto" :open-names="['1']">
            <Submenu name="sys-cluster-management">
              <template slot="title">
                <Icon type="ios-navigate"></Icon>
                集群管理
              </template>
              <MenuItem name="sys-cluster-list" :to="{name:'cluster_list'}">
                <!--                <router-link :to="{name:'cluster_list'}">-->
                集群列表
                <!--                </router-link>-->
              </MenuItem>
              <MenuItem name="sys-machine-list" :to="{name:'instance_list'}">实例列表</MenuItem>
            </Submenu>
            <Submenu name="sys-task-management">
              <template slot="title">
                <Icon type="ios-keypad"></Icon>
                任务管理
              </template>
              <MenuItem name="sys-task-list" :to="{name:'task_list'}">任务列表</MenuItem>
            </Submenu>
            <Submenu name="sys-datasource-management">
              <template slot="title">
                <Icon type="logo-buffer"/>
                数 据 源
              </template>
              <MenuItem name="sys-datasource-list" :to="{name:'db_datasource_list'}">数据源列表</MenuItem>
            </Submenu>
            <Submenu name="sys-configuration">
              <template slot="title">
                <Icon type="ios-cog-outline"/>
                系统设置
              </template>
              <MenuItem name="sys-configuration-basic">基础配置</MenuItem>
            </Submenu>
          </Menu>
        </Sider>
        <Layout :style="{padding: '12px 12px 12px'}">
          <Content :style="{padding: '5px', background: '#fff',height:'100%',overflow:'scroll'}">
            <router-view></router-view>
          </Content>
        </Layout>
      </Layout>
      <Footer>2022 &copy; xxx</Footer>
    </Layout>
  </div>
</template>
<script>
export default {
  name: "Home"
}
</script>
