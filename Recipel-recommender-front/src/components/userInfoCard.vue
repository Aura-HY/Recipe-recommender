<template>
    <!-- 用户信息卡片 -->
    <div class="user-card">
    <div class="avatar">
        <span></span>
    </div>
    <div class="user-info">
        <h2>{{  name }}</h2>
        <p>健康饮食爱好者</p>
    </div>
    <div class="edit-btn">
        <button>编辑资料</button>
    </div>
    </div>
</template>

<script>
export default {
 name: '', // 组件的名称
 components: {}, // 用于注册子组件的地方
 data() { // 定义组件的响应式数据
return {
    name: '李', // 用户名
};
},
 computed: {}, // 定义计算属性
 created() {}, // 组件的生命周期钩子，创建时调用
 mounted() {}, // 组件的生命周期钩子，挂载时调用
 updated() {}, // 组件的生命周期钩子，更新时调用
 methods: {}, // 定义组件的方法
};
</script>
<!-- css样式 -->
<style scoped>
</style>