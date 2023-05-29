<!-- eslint-disable vue/no-multiple-template-root -->
<template>
  <div
    class="wholepage"
    @mouseover="stopInv()"
    @mouseout="runInv()"
    @touchstart="stopInv()"
    @touchend="runInv()"
  >
    <!--滚动图片，后面在这里添加点击跳转-->
    <div class="item">
      <img :src="dataList[currentIndex]" />
    </div>
    <!--手动切换图片-->
    <div class="page" v-if="this.dataList.length > 1">
      <ul>
        <li @click="gotoPage(prevIndex)">&lt;</li>
        <li
          v-for="(item, index) in dataList"
          @click="gotoPage(index)"
          :class="{ current: currentIndex == index }"
          :key="index"
        >
          {{ index + 1 }}
        </li>
        <li @click="gotoPage(nextIndex)">&gt;</li>
      </ul>
    </div>
    <!--后端拉取显示商品，尚未实现-->
    <div class="product-list">
      <div v-for="product in products" :key="product.id" class="product">
        <img :src="product.image" :alt="product.introduce" />
        <h3>{{ product.introduce }}</h3>
        <p>价格: {{ product.price }}</p>
      </div>
    </div>
  </div>
  <Recommend></Recommend>
</template>


<script>
import axios from 'axios'
import Recommend from './childHome/Recommend.vue'
export default {
  data () {
    return {
      searchQuery: '',
      products: [],
      // 这里只是一个测试，可以从后端拉去对应的图片然后放在这里的datalist
      dataList: [
        'https://i1.mifile.cn/a4/xmad_15535933141925_ulkYv.jpg',
        'https://i1.mifile.cn/a4/xmad_15532384207972_iJXSx.jpg',
        'https://i1.mifile.cn/a4/xmad_15517939170939_oiXCK.jpg'
      ],
      currentIndex: 0, // 默认显示图片
      timer: null // 定时器
    }
  },
  mounted () {
    this.fetchProducts()
    setTimeout(() => {
      this.runInv()
    })
  },
  // 处理图片切换
  computed: {
    // 上一张
    prevIndex () {
      if (this.currentIndex === 0) {
        return this.dataList.length - 1
      } else {
        return this.currentIndex - 1
      }
    },
    // 下一张
    nextIndex () {
      if (this.currentIndex === this.dataList.length - 1) {
        return 0
      } else {
        return this.currentIndex + 1
      }
    }
  },
  methods: {
    // 定时滚动图片
    stopInv () {
      clearInterval(this.timer)
    },
    runInv () {
      this.timer = setInterval(() => {
        this.gotoPage(this.nextIndex)
      }, 3000)
    },
    gotoPage (index) {
      this.currentIndex = index
    },
    // 拉取后端数据
    fetchProducts () {
      axios
        .get('/api/products') // 根据你的实际API路由进行调整
        .then(response => {
          this.products = response.data
        })
        .catch(error => {
          console.error(error)
        })
    },
    /*
    search () {
      axios
        .get('/api/products', {
          params: {
            searchQuery: this.searchQuery
          }
        })
        .then(response => {
          this.products = response.data
        })
        .catch(error => {
          console.error(error)
        })
    }
  },*/
},
  components: {
    Recommend
  }
}
</script>

<style scoped>
.flex-grow {
  flex-grow: 1;
}

* {
  margin: 0;
  padding: 0;
}

header {
  position: relative;
  margin-bottom: 20px;
  margin-top: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

input[type="text"] {
  width: 300px;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
  margin-right: 10px;
}

button {
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.product-list {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-gap: 20px;
}

.product {
  text-align: center;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 4px;
}

.product img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  margin-bottom: 10px;
}

.product h3 {
  margin-bottom: 5px;
}

.product p {
  color: #888;
  font-size: 14px;
}

ul li {
  list-style: none;
  float: left;
  width: 30px;
  height: 40px;
  line-height: 40px;
  text-align: center;
  cursor: pointer;
  color: rgba(255, 255, 255, 0.8);
  font-size: 14px;
}

.wholepage {
  max-width: 1800px;
  margin: 0 auto;
  position: relative;
}

.wholepage img {
  width: 100%;
  display: block;
}

.wholepage .page {
  background: rgba(0, 0, 0, 0.5);
  position: absolute;
  right: 0;
  bottom: 0;
  width: 100%;
}

.wholepage .page ul {
  float: right;
}

.current {
  color: #ff6700;
}
</style>
