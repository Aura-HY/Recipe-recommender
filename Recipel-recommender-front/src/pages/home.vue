<template>
  <div class="profile-page">
    <!-- 顶部标题栏 -->
    <div class="header">
      <h1>我的</h1>
    </div>

    <!-- 用户信息卡片 -->
    <div class="user-card">
      <div class="avatar">
        <img :src="account.avatar" alt="头像" class="avatar-responsive" />
      </div>
      <div class="user-info">
        <h2>{{ account.name }}</h2>
        <p>健康饮食爱好者</p>
      </div>
      <div class="edit-btn">
        <button @click="goToEdit">编辑资料</button>
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

<script>
import * as CasdoorSDK from "casdoor-vue-sdk";
import backend from "../backend/casdoorBackend";
import { getMyProfileUrl } from "casdoor-vue-sdk";

export default {
  data() {
    return {
      token:
        "eyJhbGciOiJSUzI1NiIsImtpZCI6ImNlcnRfZm9vZCIsInR5cCI6IkpXVCJ9.eyJvd25lciI6ImZvcl9oZWFsdGgiLCJuYW1lIjoiajEyMyIsImNyZWF0ZWRUaW1lIjoiMjAyNS0wMy0xN1QyMDoxMDoyMyswODowMCIsInVwZGF0ZWRUaW1lIjoiMjAyNS0wMy0xOFQyMjo0OTowOCswODowMCIsImRlbGV0ZWRUaW1lIjoiIiwiaWQiOiJiMzQ1N2ZiOS0wZjU0LTRhNzEtYjJhOS03MTI5MTdhYWQyMTYiLCJ0eXBlIjoibm9ybWFsLXVzZXIiLCJwYXNzd29yZCI6IiIsInBhc3N3b3JkU2FsdCI6IiIsInBhc3N3b3JkVHlwZSI6InBsYWluIiwiZGlzcGxheU5hbWUiOiJ3d3ciLCJmaXJzdE5hbWUiOiIiLCJsYXN0TmFtZSI6IiIsImF2YXRhciI6Imh0dHBzOi8vcmVjaXBlLTEzMjYyNjk1NzcuY29zLmFwLWd1YW5nemhvdS5teXFjbG91ZC5jb20vYXZhdGFyL2Zvcl9oZWFsdGgvajEyMy53ZWJwIiwiYXZhdGFyVHlwZSI6IiIsInBlcm1hbmVudEF2YXRhciI6IiIsImVtYWlsIjoiIiwiZW1haWxWZXJpZmllZCI6ZmFsc2UsInBob25lIjoiIiwiY291bnRyeUNvZGUiOiJVUyIsInJlZ2lvbiI6IiIsImxvY2F0aW9uIjoiIiwiYWRkcmVzcyI6W10sImFmZmlsaWF0aW9uIjoiIiwidGl0bGUiOiIiLCJpZENhcmRUeXBlIjoiIiwiaWRDYXJkIjoiIiwiaG9tZXBhZ2UiOiIiLCJiaW8iOiIiLCJsYW5ndWFnZSI6InpoIiwiZ2VuZGVyIjoiIiwiYmlydGhkYXkiOiIiLCJlZHVjYXRpb24iOiIiLCJzY29yZSI6MCwia2FybWEiOjAsInJhbmtpbmciOjEsImlzRGVmYXVsdEF2YXRhciI6ZmFsc2UsImlzT25saW5lIjpmYWxzZSwiaXNBZG1pbiI6ZmFsc2UsImlzRm9yYmlkZGVuIjpmYWxzZSwiaXNEZWxldGVkIjpmYWxzZSwic2lnbnVwQXBwbGljYXRpb24iOiJyZWNpcGUtcmVjb21tZW5kZXIiLCJoYXNoIjoiIiwicHJlSGFzaCI6IiIsImFjY2Vzc0tleSI6IiIsImFjY2Vzc1NlY3JldCI6IiIsImdpdGh1YiI6IiIsImdvb2dsZSI6IiIsInFxIjoiIiwid2VjaGF0IjoiIiwiZmFjZWJvb2siOiIiLCJkaW5ndGFsayI6IiIsIndlaWJvIjoiIiwiZ2l0ZWUiOiIiLCJsaW5rZWRpbiI6IiIsIndlY29tIjoiIiwibGFyayI6IiIsImdpdGxhYiI6IiIsImNyZWF0ZWRJcCI6IiIsImxhc3RTaWduaW5UaW1lIjoiIiwibGFzdFNpZ25pbklwIjoiIiwicHJlZmVycmVkTWZhVHlwZSI6IiIsInJlY292ZXJ5Q29kZXMiOm51bGwsInRvdHBTZWNyZXQiOiIiLCJtZmFQaG9uZUVuYWJsZWQiOmZhbHNlLCJtZmFFbWFpbEVuYWJsZWQiOmZhbHNlLCJsZGFwIjoiIiwicHJvcGVydGllcyI6e30sInJvbGVzIjpbXSwicGVybWlzc2lvbnMiOltdLCJncm91cHMiOltdLCJsYXN0U2lnbmluV3JvbmdUaW1lIjoiIiwic2lnbmluV3JvbmdUaW1lcyI6MCwibWFuYWdlZEFjY291bnRzIjpudWxsLCJ0b2tlblR5cGUiOiJhY2Nlc3MtdG9rZW4iLCJ0YWciOiIiLCJzY29wZSI6InByb2ZpbGUiLCJhenAiOiI0MDAzYzQ4YzY5NTg4Y2I3N2Q3NSIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODAwMCIsInN1YiI6ImIzNDU3ZmI5LTBmNTQtNGE3MS1iMmE5LTcxMjkxN2FhZDIxNiIsImF1ZCI6WyI0MDAzYzQ4YzY5NTg4Y2I3N2Q3NSJdLCJleHAiOjE3NDI5MTUwMjMsIm5iZiI6MTc0MjMxMDIyMywiaWF0IjoxNzQyMzEwMjIzLCJqdGkiOiJhZG1pbi9hMDg0ODUwZi1hNzk2LTQ3OGItOTllZi1jY2QwNWJmZTNmYmIifQ.ln9SGkMslvl-Eut3OkJbskTKGeaC9HeA9DevBX5SZ-aDK7k2E1gvWkuHMtY_Rzs0AdYMQeYjEfMolMcSjo6Spj7KMkcZmtNhXnhBaZjq5rYQrWBo4vORGK17evJKXWlVmq1wRnVgCu4ZlU6dScEIZDWajBur5hDHXmHbWaZrMXpATyWeZ6PuPlLIg2BXzl2RQ_6FXLrly8KSGCrbFyw4mP2gTPHj4bcoH_lnQx8pL3K7YaqT0m5Sou14JGBYGXU4-k3g8aXG9EU1iiF58dVeuwih-x8d7eG1SNTBHBQ_qU8bkp5NLAOi3eoYluW7U213B9EClAFhU_csh1B81Q4kHDsjj0_6B5pq6EUmhPNuuQGzqnsY6jm1y2aW74S4wAKj36_1OrszSet_3Z4SX0Kd-wPV71Dku_rqrRfw9L8LSXlWZvQb_T4As055pl0V949M-dEd7gT8m-ra-ii5FwisxxbByw4BicUXDmiaeT8zMDpi8GdmJQ9Z5DOolbPuZK1jwflTO3RhdFsjw32upu4UJYbKmVrY4vmED16TNwXHkMZ7Avtu3FZxW8P_JGjehW_-e0xKKIVyP4VpD78u_hn5zeodRLtZNGUBAFASA2NLvqnpoFal1h9plihHT6DS4rcLsdfbI6rnsdSPf6mnzw68W6BpLUMG523_DFGUqcL3eIc", // 需要替换成用户的 OAuth Token
      account: {
        name: "李",
        avatar: "",
        email: "",
      },
    };
  },
  methods: {
    goToEdit() {
      window.location.href = this.getMyProfileUrl(this.account);
    },
  },
  mounted() {
    // 获取用户信息
    backend.getAccount().then((res) => {
      console.log(res);
      if (res["status"] === "ok") {
        this.account = res["data"];
        console.log(this.getMyProfileUrl(this.account));
        // window.location.href = this.getMyProfileUrl(this.account);
        console.log("success:", this.account);
      } else {
        console.log("fail:", res);
      }
    });
  },
};

// 这里可以添加组件逻辑
</script>

<style scoped>
.profile-page {
  font-family: "PingFang SC", "Helvetica Neue", Arial, sans-serif;
  background-color: #f5f5f5;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
  padding-bottom: 60px; /* 为底部导航留出空间 */
}

/* 顶部标题栏 */
.header {
  background-color: #fff2cc;
  padding: 15px 0;
  text-align: center;
}

.header h1 {
  margin: 0;
  font-size: 18px;
  color: #8b4513;
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
  object-fit: cover; /* 保持内容填充但不变形 */
  width: 60px;
  height: 60px;
  background-color: #fff2cc;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  color: #8b4513;
}
.avatar-responsive {
  width: 100%;
  height: auto;
  object-fit: contain;
  border-radius: 50%;
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
  background-color: #fff2cc;
  border: none;
  border-radius: 20px;
  padding: 8px 15px;
  font-size: 14px;
  color: #8b4513;
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
  background-color: #fff9e6;
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
  color: #8b4513;
  font-weight: bold;
}

.metric-unit {
  font-size: 12px;
  color: #999;
  text-align: right;
}

/* 图标样式 */
.icon-fire::before {
  content: "🔥";
}
.icon-protein::before {
  content: "🥩";
}
.icon-water::before {
  content: "💧";
}
.icon-exercise::before {
  content: "🏃";
}

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
  color: #ff9800;
}

.recipes-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.recipe-card {
  background-color: #ffecb3;
  border-radius: 10px;
  height: 120px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #8b4513;
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
  color: #ff9800;
}

/* 图标样式 */
.icon-home::before {
  content: "🏠";
}
.icon-cart::before {
  content: "🛒";
}
.icon-profile::before {
  content: "👤";
}
</style>
