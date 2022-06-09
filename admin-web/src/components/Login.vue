<template>
  <div class="layout" style="background: #f5f7f9" >
    <Layout :style="{backgroundImage:`url(${require('../assets/background.jpeg')})`,backgroundRepeat:'no-repeat',backgroundSize:'100% 100%'}">
      <Row style="margin-top: 15%">
        <Col :offset="14" :span="8" style="min-width: 400px;max-width: 600px">
          <Card title="登陆">
            <Row class-name="login-row">
              <Input placeholder="username......" prefix="ios-contact"></Input>
            </Row>
            <Row class-name="login-row">
              <Input password placeholder="password......" type="password"></Input>
            </Row>
            <Row class-name="login-row">
              <Button long type="success">登陆</Button>
            </Row>
          </Card>
        </Col>
      </Row>
      <Footer :style="{position:'fixed',bottom:'0px',left:'0px',width:'100%',backgroundColor:'transparent',color:'black'}">2022 &copy; XXX</Footer>
    </Layout>
  </div>
</template>

<script>
export default {
  name: "Login"
}
</script>

<style scoped>
.login-row {
  margin-top: 40px;
}


</style>
