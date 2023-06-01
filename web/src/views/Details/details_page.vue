<template>
  <div v-if="goodsList[good_id]">
    <!-- 图片 -->
    <div class="pic">
      <img :src="goodsList[good_id].productLink" />
    </div>
    <!-- 文字 -->
    <div class="fonts">
      <!-- 标题 -->
      <div class="titles">
        <slot:titles>
          <h3>
            {{ goodsList[good_id].productName }}
          </h3>
        </slot:titles>
      </div>
      <!-- 价格 -->
      <div class="price">
        <slot:price>
          <h2>￥{{ goodsList[good_id].productPrice }}</h2>
        </slot:price>
      </div>
      <!-- 购买数量 -->
      <div class="num">
        <div>
          <button @click="add(good_id)">+</button>
          <span>
            <!--购买数量：{{ this.$store.state.goodsList[good_id].counter }}-->
          </span>
          <button @click="sub(good_id)" :disabled="dis">-</button>
        </div>
        <!-- 购物车 -->
        <button
          class="stopcart"
          :class="{ activeBgColor: isActive == true }"
          @click="addToCart"
        >
          加入购物车
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  // eslint-disable-next-line vue/multi-word-component-names, vue/no-reserved-component-names
  name:'Details',
  data () {
    return {
      good_id: this.$route.params.goodId,
      goodsList:null,
      isActive: false,
      dis: true
    }
  },
  mounted () {
    axios.get('/api/goodslist').then(response => {
      this.goodsList = response.data
    }).catch(error => {
      console.error(error)
    })
  },
  methods: {  
    add (good_id) {
      this.$store.state.goodsList[good_id].counter++
      if (this.$store.state.goodsList[good_id].counter >= 0) {
        this.dis = false
      }
    },
    sub (good_id) {
      this.$store.state.goodsList[good_id].counter--
      if (this.$store.state.goodsList[good_id].counter == 0) {
        this.dis = true
      } else {
        this.dis = false
      }
    },
    goHome () {
      this.$router.push('/home')
    },
    addToCart () {
      if (this.$store.state.goodsList[this.good_id].counter >= 1) {
        this.isActive = true
        // 传入商品id
        const goodsListID = this.$store.state.goodsList[this.good_id].id
        const cart = this.$store.state.cart
        const result = cart.find(ele => ele.id === goodsListID)
        if (result === undefined) {
          cart.push(this.$store.state.goodsList[this.good_id])
        } else {
          document.querySelector('.stopcart').innerHTML = '请勿重复添加'
        }
      }
    }
  }
}
</script>

<style scoped>
.nav-details {
  background-color: rgb(248, 94, 73);
  color: #fff;
}
.nav-details div img {
  width: 25px;
  position: relative;
  top: 6px;
}
.nav-left {
  margin-left: 10px;
}
.pic {
  text-align: center;
  margin-top: 20px;
}
.pic img {
  width: 300px;
  height: 300px;
}
.fonts {
  text-align: center;
  margin-top: 20px;
}

.price {
  color: #f00;
}
.num {
  margin-top: 6px;
  position: relative;
  /* margin-bottom: 100px; */
}
.num span {
  font-size: 20px;
  font-weight: 700;
  margin-left: 10px;
  margin-right: 10px;
}
button {
  width: 30px;
}
.stopcart {
  width: 150px;
  height: 30px;
  margin-top: 10px;
  color: #fff;
  background-color: rgb(248, 94, 73);
}
.activeBgColor {
  background-color: rgb(90, 138, 76);
}
</style>
