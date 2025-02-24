// src/store/index.js
import Vue from'vue';
import Vuex from 'vuex';
import { getProductTypes } from '@/apis/product';

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    productTypes: []
  },
  mutations: {
    setProductTypes(state, productTypes) {
      state.productTypes = productTypes;
    }
  },
  actions: {
    fetchProductTypes({ commit }) {
      getProductTypes()
      .then(resp => {
        commit('setProductTypes', resp.data.types);
      })
      .catch(error => {
        console.error('Failed to fetch product types:', error);
      });
    }
  },
  getters: {
    productTypes: state => state.productTypes
  }
});
