<template>
  <div class="copyright2">
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
.copyright2 {
  white-space: nowrap;
  position: absolute;
  height: 10%;
  padding: 1%;
  width: 50%;
  top: 78%;
  left: 30%;
  color: #999;
  font-size: 1.25vw;

  a {
    color: #999;
    line-height: 1%;
    padding: 0 1%;
    border-right: 1px solid #999;

    &:last-child {
      border-right: none;
    }
  }
}
</style>
