<template>
    <div class="profile-page">
    <!-- 顶部标题栏 -->
    <div class="header">
        <h1>我的</h1>
    </div>
    
    <!-- 用户信息卡片 -->
    <div class="user-card">
        <div class="avatar">
        <span>李</span>
        </div>
        <div class="user-info">
        <h2>李小厨</h2>
        <p>健康饮食爱好者</p>
        </div>
        <div class="edit-btn">
        <button>编辑资料</button>
        </div>
    </div>
    
    <!-- 健康指数部分 -->
    <div class="health-index">
        <div class="section-title">健康指数</div>
        <div class="metrics-grid">
        <!-- 卡路里 -->
        <div class="metric-card">
            <div class="metric-icon calories">
            <i class="icon-fire"></i>
            </div>
            <div class="metric-name">卡路里</div>
            <div class="metric-value">1850</div>
            <div class="metric-unit">千卡/天</div>
        </div>
        
        <!-- 蛋白质 -->
        <div class="metric-card">
            <div class="metric-icon protein">
            <i class="icon-protein"></i>
            </div>
            <div class="metric-name">蛋白质</div>
            <div class="metric-value">75</div>
            <div class="metric-unit">克/天</div>
        </div>
        
        <!-- 水分 -->
        <div class="metric-card">
            <div class="metric-icon water">
            <i class="icon-water"></i>
            </div>
            <div class="metric-name">水分</div>
            <div class="metric-value">2000</div>
            <div class="metric-unit">毫升/天</div>
        </div>
        
        <!-- 运动 -->
        <div class="metric-card">
            <div class="metric-icon exercise">
            <i class="icon-exercise"></i>
            </div>
            <div class="metric-name">运动</div>
            <div class="metric-value">45</div>
            <div class="metric-unit">分钟/天</div>
        </div>
        </div>
    </div>
    
    <!-- 收藏的食谱 -->
    <div class="recipes">
        <div class="section-header">
        <div class="section-title">收藏的食谱</div>
        <div class="view-all">查看全部</div>
        </div>
        <div class="recipes-grid">
        <div class="recipe-card">食谱图片</div>
        <div class="recipe-card">食谱图片</div>
        </div>
    </div>
    
    <!-- 底部导航栏 -->
    <div class="bottom-nav">
        <div class="nav-item">
        <i class="icon-home"></i>
        <span>首页</span>
        </div>
        <div class="nav-item">
        <i class="icon-cart"></i>
        <span>购物车</span>
        </div>
        <div class="nav-item active">
          <i class="icon-profile"></i>
          <span>我的</span>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  // 这里可以添加组件逻辑
  </script>
  
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
    top: 15px;
    left: 15px;
  }
  
  .metric-name {
    font-size: 14px;
    color: #666;
    margin-bottom: 5px;
  }
  
  .metric-value {
    font-size: 24px;
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