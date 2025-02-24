<template>
    <div>
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm">
        <el-form-item label="Product Name" prop="name">
            <el-input v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="Product Type" prop="type">
            <el-select v-model="ruleForm.type" placeholder="Product Type">
                <el-option v-for="productType in productTypes" :key="productType" :label="productType" :value="productType"></el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="Price" prop="price">
            <el-input v-model="ruleForm.price"></el-input>
        </el-form-item>
        <el-form-item label="Description" prop="description">
            <el-input type="textarea" v-model="ruleForm.description"></el-input>
        </el-form-item>
        <el-form-item label="Image" prop="image_path">
            <el-upload
            ref="upload"
            class="avatar-uploader"
            :action="imageUploadUrl"
            :show-file-list="false"
            :before-upload="beforeImageUpload"
            :on-success="handleImageSuccess"
            :on-error="handleImageError"
            >
                <img v-if="ruleForm.image_path" :src="fullImageUrl" class="avatar">
                <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
        </el-form-item>
        <el-form-item class="button-container">
            <el-button type="primary" @click="submitForm('ruleForm')">{{ isEditMode ? 'Update' : 'Create' }}</el-button>
            <el-button @click="resetForm">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
</template>

<style scoped>
  .button-container {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
  .avatar-uploader img {
    width:100%;
  }
  </style>
  
<script>
  import { createProduct, getProduct, updateProduct } from '@/apis/product'
  import { mapGetters } from 'vuex';
  import { BACKEND_BASE_URL, IMG_UPLOAD_URL } from '@/config';

  export default {
    mounted() {
      this.$store.dispatch('fetchProductTypes');
      const productId = this.$route.query.id;
      if (productId) {
        this.productId = productId;
        this.isEditMode = true;
        this.getProductDetails(productId);
      }
    },
    data() {
      return {
        baseUrl: BACKEND_BASE_URL,
        isEditMode: false,
        productId: 0,
        ruleForm: {
          name: '',
          type: '',
          price: '',
          description: '',
          image_path: '',
        },
        imageUploadUrl: IMG_UPLOAD_URL,
        rules: {
          name: [
            { required: true, message: 'Please input product name', trigger: 'blur' },
            { min: 1, max: 20, message: 'Length should be 1 to 20', trigger: 'blur' }
          ],
          type: [
            { required: true, message: 'Please select product type', trigger: 'change' }
          ],
          price: [
            { 
                required: true, 
                message: 'Please input product price', 
                trigger: 'blur' 
            },
            { 
                pattern: /^\d+(\.\d{1,2})?$/, 
                message: 'Price must be a valid number with at most 2 decimal places', 
                trigger: 'blur'
            }
          ],
          description: [
            { required: false, message: 'Please input description', trigger: 'blur' }
          ],
          image_path: [
            { required: true, message: 'Please upload a product image', trigger: 'change' }
          ]
        }
      };
    },
    computed: {
        ...mapGetters(['productTypes']),
        fullImageUrl() {
            if (this.ruleForm.image_path.startsWith("/")) {
                return `${this.baseUrl}${this.ruleForm.image_path}`
            }
            return `${this.baseUrl}/${this.ruleForm.image_path}`
        }
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if (!this.isEditMode) {
                createProduct({
                    name: this.ruleForm.name,
                    type: this.ruleForm.type,
                    price: Number(this.ruleForm.price),
                    description: this.ruleForm.description,
                    image_path: this.ruleForm.image_path
                })
                .then(() => {
                    this.$message.success("Product created successfully!")
                    this.resetForm()
                    this.$router.push({ path: '/' })
                })
                .catch((err) => {
                    this.$message.error(`Product creation failed: ${err.message}`)
                })
            } else {
                updateProduct(this.productId, {
                    name: this.ruleForm.name,
                    type: this.ruleForm.type,
                    price: Number(this.ruleForm.price),
                    description: this.ruleForm.description,
                    image_path: this.ruleForm.image_path
                })
                .then(() => {
                    this.$message.success("Product updated successfully!")
                    this.$router.push({ path: '/' })
                })
                .catch((err) => {
                    this.$message.error(`Product update failed: ${err.message}`)
                })
            }
            
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm() {
        this.ruleForm = {
          name: '',
          type: '',
          price: '',
          description: '',
          image_path: '',
        }
      },
      // Handle the image upload success, get the URL
      handleImageSuccess(response) {
        this.ruleForm.image_path = `${response.url}`; // Assuming the response contains the image URL
      },
      // Handle image upload failure
      handleImageError(err) {
        this.$message.error('Image upload failed', err);
      },
      // Validate that the uploaded file is an image before upload
      beforeImageUpload(file) {
        console.log("file.type=", file.type)
        const isImage = file.type.startsWith('image/');
        if (!isImage) {
            this.$message.error('You can only upload image files!');
        }
        return isImage;
      },
      getProductDetails(id) {
        getProduct(id)
        .then((resp) => {
            const data = resp.data
            this.ruleForm.name = data.name
            this.ruleForm.type = data.type
            this.ruleForm.price = String(data.price)
            this.ruleForm.description = data.description
            this.ruleForm.image_path = data.image_path
        })
        .catch((err) => {
            this.$message.error(`Get Product ${id} failed: ${err.message}`)
        })
      }
    }
  }
</script>
  