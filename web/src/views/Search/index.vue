<template>
  <div class="headText">
    <p>以下是 {{ key }} 的搜索结果 :</p>
  </div>
  <el-empty description="没有找到相关结果" v-if="no_result" :image-size="300" />
  <div class="total">
    <div class="layer_good">
      <div
        class="recommend"
        v-for="(item, index) in goodsList.products"
        :key="index"
      >
        <router-link
          :to="{ name: 'Details', params: { goodId: item.productId } }"
        >
          <ul>
            <li>
              <img :src="item.productLink" class="lazy" />
            </li>
            <li class="introduce">{{ item.productName }}</li>
            <li class="price">￥{{ item.productPrice }}</li>
          </ul>
        </router-link>
      </div>
    </div>
    <el-pagination
      background="true"
      layout="prev, pager, next"
      :total="total_page"
      class="pageSelect"
      v-model:current-page="current_page"
      @current-change="onCurrentChange()"
    />
  </div>
</template>

<script setup>
import axios from "axios";
import { ElMessage } from "element-plus";
import { onMounted, ref, reactive } from "vue";
import { useRoute } from "vue-router";
const goodsList = reactive({ products: [], totalPage: "" });
const route = useRoute();
const key = ref(route.params.prom);
const current_page = ref();
const total_page = ref();
const no_result = ref(false);

const getgoodList = () => {
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/Search",
    data: {
      name: key.value,
      pageNum: 1,
    },
  })
    .then((response) => {
      goodsList.products = response.data.products;
      goodsList.totalPage = response.data.totalPage;
      total_page.value = parseInt(goodsList.totalPage) * 10;
      console.log(total_page);
      console.log(goodsList);
    })
    .catch(() => {
      console.log("搜索失败");
      no_result.value = true;
    });
};

const onCurrentChange = () => {
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/Search",
    data: {
      name: key.value,
      pageNum: current_page.value,
    },
  })
    .then((response) => {
      goodsList.products = response.data.products;
      goodsList.totalPage = response.data.totalPage;
      total_page.value = parseInt(goodsList.totalPage) * 10;
    })
    .catch((error) => {
      console.log("搜索失败");
    });
};

onMounted(() => {
  getgoodList();
});
</script>

<style scoped>
.headText {
  position: relative;
  top: 30px;
  left: 20px;
  font-size: 20px;
}

.total {
  position: relative;
  top: 50px;
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

.pageSelect {
  position: relative;
  align-content: center;
  text-align: center;
  bottom: 40px;
}
</style>
