<template>
  <div class="product-detail">
    <div class="product-image-container">
      <!-- 商品图片 -->
      <img class="product-image" :src="product.image" alt="Product Image" />
    </div>
    <div class="product-info-container">
      <!-- 商品标题 -->
      <h1 class="product-title">{{ product.title }}</h1>

      <!-- 商品价格 -->
      <p class="product-price">价格：{{ product.price }}</p>

      <!-- 商品评分 -->
      <div class="product-rating">
        <el-rate v-model="product.rating" :max="5" :allow-half="true"></el-rate>
        <span class="rating-value">({{ product.rating }})</span>
      </div>

      <!-- 商品描述 -->
      <p class="product-description">{{ product.description }}</p>

      <!-- 商品参数 -->
      <div class="product-parameters">
        <h3>商品参数</h3>
        <table>
          <tr v-for="parameter in product.parameters" :key="parameter.name">
            <td>{{ parameter.name }}</td>
            <td>{{ parameter.value }}</td>
          </tr>
        </table>
      </div>

      <!-- 添加到购物车按钮 -->
      <el-button type="primary" @click="addToCart">加入购物车</el-button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      product: {
        image: 'https://gd1.alicdn.com/imgextra/i1/60075666/O1CN01k1iUIj1rj4oG1zWZv_!!60075666.jpg',
        title: '商品标题',
        price: '¥19.99',
        rating: 4.5,
        description: '商品描述',
        parameters: [
          { name: '参数1', value: '值1' },
          { name: '参数2', value: '值2' },
          { name: '参数3', value: '值3' }
        ]
      }
    };
  },
  methods: {
    addToCart() {
      // 添加到购物车的逻辑
    }
  }
};
</script>

<style>
.product-detail {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  margin-top: 50px;
}

.product-image-container {
  margin-right: 20px;
}

.product-image {
  width: 400px;
  height: auto;
}

.product-info-container {
  flex: 1;
}

.product-title {
  font-size: 24px;
  margin-bottom: 10px;
}

.product-price {
  font-size: 18px;
  margin-bottom: 10px;
}

.product-rating {
  margin-bottom: 10px;
}

.rating-value {
  margin-left: 5px;
}

.product-description {
  margin-bottom: 20px;
}

.product-parameters {
  margin-bottom: 20px;
}

.product-parameters h3 {
  font-size: 18px;
  margin-bottom: 10px;
}

.product-parameters table {
  width: 100%;
}

.product-parameters td {
  padding: 5px 0;
}

.el-button {
  margin-top: 20px;
}
</style>
