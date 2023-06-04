<template>
  <div class="copyright">
    <p>
      <a href="javascript:;">关于我们</a>
      <a href="javascript:;">帮助中心</a>
      <a href="javascript:;">售后服务</a>
      <a href="javascript:;">配送与验收</a>
      <a href="javascript:;">商务合作</a>
      <a href="javascript:;">搜索推荐</a>
      <a href="javascript:;">友情链接</a>
    </p>
  </div>
</template>

<style lang="scss">
.copyright {
  height: 150px;
  padding-top: 40px;
  text-align: center;
  color: #999;
  font-size: 15px;

  p {
    line-height: 1;
    margin-bottom: 20px;
  }

  a {
    color: #999;
    line-height: 1;
    padding: 0 10px;
    border-right: 1px solid #999;

    &:last-child {
      border-right: none;
    }
  }
}
</style>
