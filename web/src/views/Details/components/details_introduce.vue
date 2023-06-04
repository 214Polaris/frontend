<!--商品信息介绍-->
<template>
  <!--拿去商品信息后放在这里-->
  <p class="g-name">{{ GoodName }}</p>
  <p class="g-desc">{{ GoodIntroduce }}</p>
  <p class="g-price">
    <span>{{ GoodPrice - 20 }}</span>
    <span>{{ GoodPrice }}</span>
  </p>
  <div class="g-service">
    <dl>
      <dt>促销</dt>
      <dd>新商城开业！所有商品折扣促销！！！/dd></dd>
    </dl>

    <dl>
      <dt>配送</dt>
      <dd>至</dd>
    </dl>
    <dl>
      <dt>服务</dt>
      <dd>
        <span>无忧退货</span>
        <span>快速退款</span>
        <span>到额包邮</span>
        <a href="javascript:;">了解详情</a>
      </dd>
    </dl>
  </div>
</template>

<script>
export default {
  name: "GoodsIntroduce",
  //props处获取参数
  props: {
    GoodName: String,
    GoodIntroduce: String,
    GoodPrice: int,
  },
};
</script>

<style lang="scss" scoped>
@import "@/assets/variables.scss";
.g-name {
  font-size: 22px;
}
.g-desc {
  color: #999;
  margin-top: 10px;
}
.g-price {
  margin-top: 10px;
  span {
    &::before {
      content: "¥";
      font-size: 14px;
    }
    &:first-child {
      color: $priceColor;
      margin-right: 10px;
      font-size: 22px;
    }
    &:last-child {
      color: #999;
      text-decoration: line-through;
      font-size: 16px;
    }
  }
}
.g-service {
  background: #f5f5f5;
  width: 500px;
  padding: 20px 10px 0 10px;
  margin-top: 10px;
  dl {
    padding-bottom: 20px;
    display: flex;
    align-items: center;
    dt {
      width: 50px;
      color: #999;
    }
    dd {
      color: #666;
      &:last-child {
        span {
          margin-right: 10px;
          &::before {
            content: "•";
            color: $xtxColor;
            margin-right: 2px;
          }
        }
        a {
          color: $xtxColor;
        }
      }
    }
  }
}
</style>
