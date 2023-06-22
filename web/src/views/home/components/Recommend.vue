<template>
  <div class="total">
    <div class="layer_good">
      <div class="recommend" v-for="(item, index) in goodsList" :key="index">
        <router-link
          :to="{ name: 'Details', params: { goodId: item.productId } }"
        >
          <ul >
            <li>
              <img v-lazy="item.productLink" class="lazy" />
            </li>
            <li class="introduce">{{ item.productName }}</li>
            <li class="price">￥{{ item.productPrice }}</li>
          </ul>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "recommend",
  mounted() {
    axios
      .get("/api/goodslist")
      .then((response) => {
        this.goodsList = response.data;
      })
      .catch((error) => {
        console.error(error);
      });
  },
  data() {
    return {
      goodsList: null,
    };
  },
};
</script>

<style scoped>
.total {
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
}

.layer_good {
  text-align: center;
  flex-grow: 1;
  margin-bottom: 64px;
}

.recommend {
  display: inline-block;
  width: 300px; /* 设置每个推荐项的宽度 */
  margin: 12px; /* 设置推荐项之间的间距 */
  text-align: center;
}

li {
  list-style: none;
}

.recommend img {
  width: 100%; /* 让图片宽度充满容器 */
  height: auto; /* 根据宽度自适应调整高度 */
  text-align: center;
  box-shadow: 5px 5px 5px rgba(57, 56, 56, 0.5);
}

.introduce {
  width: 100%;
  text-align: center;
  font-size: 14px;
  margin-top: 6px;
  margin-bottom: 6px;
  color: black;
}

.price {
  width: 100%;
  text-align: center;
  font-size: 18px;
  color: #f00;
}
</style>
