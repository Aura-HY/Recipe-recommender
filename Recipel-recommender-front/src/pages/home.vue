<template>
  <div class="profile-page">
    <!-- 顶部标题栏 -->
    <div class="header">
      <h1>我的</h1>
    </div>
    
<!-- 用户信息卡片 -->
<userInfoCard/>

<!-- 健康指数部分 -->
<health/>
<!-- 收藏的食谱 -->
<recipes/>    
<!-- 底部导航栏 -->
<navigation/>

</div> 
</template>


<script setup>
import userInfoCard from '../components/userInfoCard.vue';
//const name = '李';
</script>

<script setup>
import health from '../components/health.vue';
</script>

<script setup>
import recipes from '../components/recipes.vue';
</script>

<script setup>
import navigation from '../components/navigation.vue';
</script>


<style scoped>
/* 您的样式 */
</style>

<style scoped>


.profile-page {
  font-family: 'PingFang SC', 'Helvetica Neue', Arial, sans-serif;
  background-color: #f5f5f5;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
  padding-bottom: 60px; /* 为底部导航留出空间 */
}

/* 顶部标题栏 */
.header {
  background-color: #FFF2CC;
  padding: 15px 0;
  text-align: center;
}

.header h1 {
  margin: 0;
  font-size: 18px;
  color: #8B4513;
}

/* 用户卡片 */
.user-card {
  background-color: white;
  margin: 15px;
  border-radius: 10px;
  padding: 15px;
  display: flex;
  align-items: center;
}

.avatar {
  width: 60px;
  height: 60px;
  background-color: #FFF2CC;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  color: #8B4513;
}

.user-info {
  flex: 1;
  margin-left: 15px;
}

.user-info h2 {
  margin: 0;
  font-size: 16px;
  color: #333;
}

.user-info p {
  margin: 5px 0 0;
  font-size: 14px;
  color: #666;
}

.edit-btn button {
  background-color: #FFF2CC;
  border: none;
  border-radius: 20px;
  padding: 8px 15px;
  font-size: 14px;
  color: #8B4513;
  cursor: pointer;
}

/* 健康指数部分 */
.health-index {
  background-color: white;
  margin: 15px;
  border-radius: 10px;
  padding: 15px;
}

.section-title {
  font-size: 16px;
  color: #333;
  font-weight: bold;
  margin-bottom: 15px;
}

.metrics-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.metric-card {
  background-color: #FFF9E6;
  border-radius: 10px;
  padding: 15px;
  position: relative;
}

.metric-icon {
  position: absolute;
  top: 10px;
  left: 0px;
}

.metric-name {

  margin-top: 15px;
  font-size: 20px;
  color: #666;
  margin-bottom: 6px;
}

.metric-value {
  margin-top: 12px;
  font-size: 22px;
  color: #8B4513;
  font-weight: bold;
}

.metric-unit {
  font-size: 12px;
  color: #999;
  text-align: right;
}

/* 图标样式 */
.icon-fire::before { content: "🔥"; }
.icon-protein::before { content: "🥩"; }
.icon-water::before { content: "💧"; }
.icon-exercise::before { content: "🏃"; }

/* 收藏的食谱 */
.recipes {
  background-color: white;
  margin: 15px;
  border-radius: 10px;
  padding: 15px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.view-all {
  font-size: 14px;
  color: #FF9800;
}

.recipes-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.recipe-card {
  background-color: #FFECB3;
  border-radius: 10px;
  height: 120px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #8B4513;
}

/* 底部导航栏 */
.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 60px;
  background-color: white;
  display: flex;
  border-top: 1px solid #eee;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: #999;
  font-size: 12px;
}

.nav-item i {
  font-size: 22px;
  margin-bottom: 5px;
}

.nav-item.active {
  color: #FF9800;
}

/* 图标样式 */
.icon-home::before { content: "🏠"; }
.icon-cart::before { content: "🛒"; }
.icon-profile::before { content: "👤"; }
</style>