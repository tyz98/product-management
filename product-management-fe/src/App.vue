<template>
<el-container style="border: 1px solid #eee; height: 100vh">
  <el-aside width="200px">
    <el-menu
      :default-active="activeMenu"
      class="el-menu-vertical-demo"
      @select="handleMenuSelect"
    >
      <el-menu-item index="/">
        <i class="el-icon-menu"></i>
        <span>Product List</span>
      </el-menu-item>
      <el-menu-item index="/edit">
        <i class="el-icon-edit"></i>
        <span>Create/Update Product</span>
      </el-menu-item>
    </el-menu>
  </el-aside>
  
  <el-container>
    <el-header style="display: flex; align-items: center; font-size: 24px; font-weight: bold;">
      <span>Product Management</span>
    </el-header>
    
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</el-container>
</template>

<style scoped>
  .el-header {
    background-color: #B3C0D1;
  }
  .el-menu {
    height: 100%;
  }
</style>

<script>

export default {
  name: 'App',
  data() {
    return {
      activeMenu: this.$route.path
    };
  },
  methods: {
    handleMenuSelect(index) {
      this.$router.push(index);
    }
  },
  watch: {
    $route(to) {
      this.activeMenu = to.path;
    }
  }
};
</script>