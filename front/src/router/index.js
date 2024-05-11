import {createRouter, createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/", component: () => import("@/page/home.vue"),
            children: [
                {path: "index", component: () => import("@/page/index.vue")},
                {path: "/chat/:friendId", component: () => import("@/page/chat.vue")},
            ]

        },
        {
            path: "/login", component: () => import("@/page/login.vue"),
        },
    ]
});


router.beforeEach((to, from, next) => {
    const isLogin = localStorage.getItem("token");
    if (to.path === "/login") {
        next();
    } else {
        isLogin ? next() : next("/login");
    }
})


export default router;