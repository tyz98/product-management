import Vue from 'vue';
import VueRouter from 'vue-router';
import ProductList from '../views/ProductList.vue';
import EditProduct from '../views/EditProduct.vue';

Vue.use(VueRouter);

const routes = [
  { path: '/', component: ProductList },
  { path: '/edit', component: EditProduct }
];

export default new VueRouter({
  routes
});
