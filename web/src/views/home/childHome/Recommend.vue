<template>
  <div class="total">
    <div class="layer">
      <div class="recommend" v-for="(item, index) in goodsList" :key="index">
        <router-link :to="{ name: 'Details', params: { goodId: item.id } }">
          <ul>
            <li>
              <img :src="item.image" />
            </li>
            <li class="introduce">{{ item.introduce }}</li>
            <li class="price">￥{{ item.price }}</li>
          </ul>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import goodsList from '../../../static/goodsList.json'
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'recommend',
  data () {
    return {
      goodsList
    }
  }
}
</script>

<style scoped>
.total {
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
}

.layer {
  flex-grow: 1;
  margin-bottom: 64px;
}

.recommend {
  margin: 8px;
  display: inline-block;
}

li {
  list-style: none;
}

.recommend img {
  height: 300px;
  text-align: center;
}

.introduce {
  width: 300px;
  text-align: center;
  font-size: 14px;
  margin-top: 6px;
  margin-bottom: 6px;
}

.price {
  width: 300px;
  text-align: center;
  font-size: 18px;
  color: #f00;
}
</style>
