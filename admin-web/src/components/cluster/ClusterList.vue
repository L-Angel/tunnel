<template>
  <div>
    Cluster List
    <Table :data="clusters.display"></Table>
  </div>
</template>

<script>
export default {
  name: "ClusterList",
  data() {
    return {
      clusters: {
        origin: [],
        display: []
      }
    }
  },
  methods: {
    
  }
}
</script>

<style scoped>

</style>
