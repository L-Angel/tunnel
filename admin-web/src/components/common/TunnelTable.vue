<template>
  <div>
    <Card>
      <Row>
        <Col span="8">
          <Input placeholder="Searching..."
                 v-model="keyword"
                 @on-change="onFilterChange"></Input>
        </Col>
      </Row>
    </Card>
    <Card :style="{marginTop:'10px'}">
      <Table :data="display" :columns="columns" :size="size"></Table>
      <Page :page-size="page.size"
            :page-size-opts="page.opts"
            :current="page.current"
            :total="page.total"
            @on-change="onPageChange"
            @on-page-size-change="onPageSizeChange"
            :style="{marginTop: '10px'}"
            show-elevator
            show-sizer
            show-total>

      </Page>
    </Card>

  </div>
</template>

<script>
export default {
  name: "TunnelTable",
  data() {
    return {
      keyword: "",
      origin: [],
      display: [],
      filtered: [],
      page: {
        size: 10,
        current: 1,
        total: 0,
        opts: [10, 20, 50, 100]
      }
    }
  },
  props: {
    data: {
      type: Array,
      default() {
        return []
      }
    },
    columns: {
      type: Array,
      default() {
        return [];
      }
    },
    size: {
      type: String,
      default() {
        return "small";
      }
    }
  },
  mounted() {

  },
  created() {
  },
  computed: {},
  methods: {
    onFilterChange() {
      this.search(this.keyword)
    },
    search(keyword) {
      let fd = []

      if (keyword.length <= 1) {
        fd = this.origin;
      } else {
        for (let idx = 0; idx < this.origin.length; idx++) {
          let d = this.origin[idx];
          Object.values(d).forEach((v) => {
            if ((typeof v === "string") && keyword !== null && String(v).indexOf(keyword) !== -1) {
              fd.push(d)
              return
            }
          })
        }
      }
      this.filtered = fd
      this.page.total = fd.length;
      this.splitPage()
    },
    splitPage() {
      let start = (this.page.current - 1) * this.page.size;
      let end = (this.page.current) * this.page.size;
      this.display = this.filtered.slice(start, end)
    },
    onPageChange(page) {
      this.page.current = page;
      console.log(this.page)
      this.splitPage();
    },
    onPageSizeChange(pageSize) {
      this.page.size = pageSize;
      this.splitPage();
    }
  },
  watch: {
    data: function (newValue) {
      this.origin = newValue;
      this.filtered = newValue
      this.page.total = newValue.length;
      console.log(newValue)
      this.splitPage();
    }
  }
}
</script>

<style scoped>

</style>
