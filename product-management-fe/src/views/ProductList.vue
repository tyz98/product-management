<template>
  <div>
    <el-form :inline="true" :model="filters" class="filter-form">
      <el-form-item label="Name">
        <el-input v-model="filters.name" @input="updateUrl" placeholder="Search by name" />
      </el-form-item>

      <el-form-item label="Type">
        <el-select v-model="filters.type" @input="updateUrl" placeholder="Search by type">
            <el-option v-for="productType in productTypes" :key="productType" :label="productType" :value="productType"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="Price">
        <div style="display: flex;">
          <el-input v-model="filters.priceMin" @input="updateUrl" placeholder="Min price" />
          <span style="font-size: 18px;">&nbsp;~&nbsp;</span>
          <el-input v-model="filters.priceMax" @input="updateUrl" placeholder="Max price" />
        </div>
      </el-form-item>

      <el-button @click="resetFilters">Reset</el-button>
    </el-form>

    <el-table :data="products" style="width: 100%" @sort-change="handleSort">
      <el-table-column prop="id" label="ID" sortable="custom"></el-table-column>
      <el-table-column prop="name" label="Name" sortable="custom"></el-table-column>
      <el-table-column prop="type" label="Type" sortable="custom"></el-table-column>
      <el-table-column prop="price" label="Price" sortable="custom"></el-table-column>
      <el-table-column prop="description" label="Description">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="scope.row.description" placement="top">
            <span class="description-text">{{ scope.row.description }}</span>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="image_path" label="Image">
        <template slot-scope="scope">
          <img :src="fullImageUrl(scope)" alt="Product Image" style="width: 50px; height: 50px; cursor: pointer;" @click="showImagePreview(fullImageUrl(scope))" />
        </template>
      </el-table-column>
      <el-table-column label="Action">
        <template slot-scope="scope">
          <el-button 
            type="primary" 
            size="small" 
            @click="goToEditPage(scope.row.id)">
            Update
          </el-button>
          <el-button 
            type="danger" 
            size="small" 
            @click="deleteProduct(scope.row)">
            Delete
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <div style="display: flex; justify-content: center; margin-top: 10px;">
      <el-pagination
        :current-page="currentPage"
        :page-size="pageSize"
        :page-sizes="[5, 10, 20, 50, 100]"
        :total="totalProducts"
        @current-change="handlePageChange"
        @size-change="handlePageSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
    </div>

     <!-- Image Preview Dialog -->
     <el-dialog 
      :visible.sync="dialogVisible" 
      width="50%" 
      @close="dialogVisible = false"
      class="image-dialog"
    >
      <img :src="imagePreviewUrl" alt="Image Preview" class="image-preview"/>
    </el-dialog>
  </div>
</template>

<script>
import { getProducts, deleteProduct } from '@/apis/product';
import { mapGetters } from 'vuex';
import { BACKEND_BASE_URL, DEFAULT_LIST_PAGE, DEFAULT_LIST_PAGESIZE } from "@/config";

export default {
  data() {
    return {
      baseUrl: BACKEND_BASE_URL, // Base URL for the backend server
      dialogVisible: false,
      imagePreviewUrl: '',
      products: [],
      totalProducts: 0,
      currentPage: DEFAULT_LIST_PAGE,
      pageSize: DEFAULT_LIST_PAGESIZE,
      filters: {
        name: '',
        type: '',
        priceMin: '',
        priceMax: ''
      },
      sortBy: '',
      sortOrder: ''
    };
  },
  mounted() {
    this.updateFiltersFromUrl();
    this.fetchProducts();
    this.$store.dispatch('fetchProductTypes');
  },
  methods: {
    fetchProducts() {
      const params = {
        page: this.currentPage,
        page_size: this.pageSize,
        name: this.filters.name,
        type: this.filters.type,
        price_min: this.filters.priceMin,
        price_max: this.filters.priceMax,
        sort_by: this.sortBy,
        sort_order: this.sortOrder
      };

      getProducts(params)
      .then((resp)=> {
        this.products = resp.data.products;
        this.totalProducts = resp.data.total;
      })
      .catch(err=>{
        this.$message.error(`Get products error: ${err.message}`)
      })
    },
    handleSort({order, prop}) {
      console.log("handleSort:", order, prop)
      this.sortBy = prop;
      this.sortOrder = order === 'ascending' ? 'asc' : 'desc';
      this.updateUrl();
    },
    handlePageChange(page) {
      this.currentPage = page;
      this.updateUrl();
    },
    handlePageSizeChange(size) {
      this.pageSize = size;
      this.updateUrl();
    },
    resetFilters() {
      this.filters = {
        name: '',
        type: '',
        priceMin: '',
        priceMax: ''
      };
      this.currentPage = DEFAULT_LIST_PAGE;
      this.updateUrl();
    },
    updateUrl() {
      const queryParams = {
        ...this.filters,
        page: this.currentPage,
        page_size: this.pageSize,
        sort_by: this.sortBy,
        sort_order: this.sortOrder
      };
      console.log("query params=", queryParams)
      // Convert queryParams to a query string
      const queryString = new URLSearchParams(queryParams).toString();
      const currentQueryString = new URLSearchParams(this.$route.query).toString();

      // Compare the query strings to avoid redundant navigation
      if (queryString !== currentQueryString) {
        this.$router.push({ path: '/', query: queryParams });
      }
    },
    updateFiltersFromUrl() {
      console.log("updatefiltersfromurl", this.$route.query)
      const query = this.$route.query;
      this.filters.name = query.name || '';
      this.filters.type = query.type || '';
      this.filters.priceMin = query.priceMin || '';
      this.filters.priceMax = query.priceMax || '';
      this.currentPage = parseInt(query.page) || 1;
      this.pageSize = parseInt(query.page_size) || 10;
      this.sortBy = query.sort_by || '';
      this.sortOrder = query.sort_order || '';
    },
    fullImageUrl(scope) {
      if (scope.row.image_path.startsWith("/")) {
        return `${this.baseUrl}${scope.row.image_path}`
      }
      return `${this.baseUrl}/${scope.row.image_path}`
    },
    showImagePreview(imageUrl) {
      this.imagePreviewUrl = imageUrl;  // Set the clicked image as the preview source
      this.dialogVisible = true; // Open the dialog
    },
    goToEditPage(productId) {
      this.$router.push({ path: '/edit', query: { id: productId } });
    },
    deleteProduct(row) {
      this.$confirm(`Are you sure you want to delete product "${row.name}"?`)
      .then(()=>{
        deleteProduct(row.id)
        .then(() => {
            this.$message.success("Product deleted successfully!")
            this.fetchProducts()
        })
        .catch((err) => {
            this.$message.error(`Product deletion failed: ${err.message}`)
        })
      })
      .catch((a)=>{
        console.log(a)
      })
    }
  },
  watch: {
    '$route.query': 'fetchProducts'
  },
  computed: {
    ...mapGetters(['productTypes']),
  }
};
</script>

<style>
.filter-form {
  margin-bottom: 20px;
}
.image-dialog .el-dialog__body {
  display: flex;
  justify-content: center;
}
.image-preview {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.description-text {
  display: inline-block;
  max-width: 200px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
div.el-tooltip__popper {
  max-width: 600px;
  word-wrap: break-word;
}
</style>
