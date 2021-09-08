import vue from "vue";
import vueRouter from "vue-router";
import store from "@/store/index";
import { roles } from "@/utils/variables.js";

vue.use(vueRouter);

const settingPageMenus = [
	{
		title: "分类管理",
		icon: "diff",
		path: "cataMg",
		children: [
			{
				title: "分类管理",
				icon: "mail",
				path: "/setting/cata"
			},
			{
				title: "添加分类",
				icon: "mail1",
				path: "/setting/addCata"
			}
		]
	},
	{
		title: "标签管理",
		icon: "tag",
		path: "tagMg",
		children: [
			{
				title: "标签管理",
				path: "/setting/tag"
			},
			{
				title: "添加标签",
				path: "/setting/addTag"
			}
		]
	},
	{
		title: "缩略图管理",
		icon: "tag",
		path: "thumbnailMg",
		children: [
			{
				title: "缩略图管理",
				path: "/setting/thumbnailMg"
			}
		]
	}
];

const loginPath = "/login";
const router = new vueRouter({
	mode: "history",
	routes: [
		{
			path: "/",
			name: "home",
			component: () =>
				import(/* webpackChunkName: "home" */ "../views/home/index.vue")
		},
		{
			path: loginPath,
			name: "login",
			component: () =>
				import(
					/* webpackChunkName: "login" */ "../views/login/index.vue"
				)
		},
		{
			path: "/githubLogin",
			name: "githubLogin",
			component: () =>
				import(
					/* webpackChunkName: "githubLogin" */ "../views/login/githubLogin.vue"
				)
		},
		{
			path: "/",
			name: "layout",
			redirect: "/articles",
			component: () =>
				import(
					/* webpackChunkName: "layout" */ "../components/layout.vue"
				),
			children: [
				{
					path: "tarticles",
					name: "tarticles",
					component: () =>
						import(
							/* webpackChunkName: "tarticles" */ "../components/articles/index.vue"
						)
				},
				{
					path: "articles",
					name: "articles",
					component: () =>
						import(
							/* webpackChunkName: "articles" */ "../views/articles/index.vue"
						)
				},
				{
					path: "tag/:tagId",
					name: "tagArticles",
					props: true,
					component: () =>
						import(
							/* webpackChunkName: "tagArticles" */ "../views/tagArticles/index.vue"
						)
				},
				{
					path: "a/:articleId",
					name: "article",
					props: true,
					component: () =>
						import(
							/* webpackChunkName: "article" */ "../views/article/index.vue"
						)
				},
				{
					path: "user/:userId",
					name: "user",
					meta: {
						role: [roles.admin]
					},
					component: () =>
						import(
							/* webpackChunkName: "user" */ "../views/user/index.vue"
						)
				},
				{
					path: "user/:userId/tag/:tagId",
					name: "userTagArticles",
					props: true,
					meta: {
						role: [roles.admin]
					},
					component: () =>
						import(
							/* webpackChunkName: "userTagArticles" */ "../views/userTagArticles/index.vue"
						)
				},
				{
					path: "user/:userId/cata/:cataId",
					name: "userCataArticles",
					props: true,
					component: () =>
						import(
							/* webpackChunkName: "userTagArticles" */ "../views/cataArticles/index.vue"
						)
				},
				{
					path: "editor",
					name: "editor",
					meta: {
						footer: false,
						role: [roles.admin]
					},
					component: () =>
						import(
							/* webpackChunkName: "editor" */ "../views/editor/index.vue"
						)
				},
				{
					path: "update/:articleId",
					name: "updateArticle",
					props: true,
					meta: {
						footer: false,
						role: [roles.admin]
					},
					component: () =>
						import(
							/* webpackChunkName: "editor" */ "../views/editor/index.vue"
						),
				},
				{
					path: "setting",
					name: "setting",
					redirect: "/setting/cata",
					component: () =>
						import(
							/* webpackChunkName: "setting" */ "../components/sidebarLayout.vue"
						),
					children: [
						{
							path: "cata/:id",
							name: "editCata",
							component: () =>
								import(
									/* webpackChunkName: "updateCata" */ "../views/setting/cataMg/add.vue"
								),
							props: true,
							meta: {
								role: [roles.admin],
								footer: false,
								menu: settingPageMenus
							}
						},
						{
							path: "cata",
							name: "cataMg",
							component: () =>
								import(
									/* webpackChunkName: "cataMg" */ "../views/setting/cataMg/index.vue"
								),
							meta: {
								role: [roles.admin],
								footer: false,
								menu: settingPageMenus
							}
						},
						{
							path: "addCata",
							name: "addCata",
							component: () =>
								import(
									/* webpackChunkName: "addCata" */ "../views/setting/cataMg/add.vue"
								),
							meta: {
								role: [roles.admin],
								footer: false,
								menu: settingPageMenus
							}
						},
						{
							path: "tag",
							name: "tagMg",
							component: () =>
								import(
									/* webpackChunkName: "tagMg" */ "../views/setting/tagMg/index.vue"
								),
							meta: {
								footer: false,
								role: [roles.admin],
								menu: settingPageMenus
							}
						},
						{
							path: "addTag",
							name: "addTag",
							component: () =>
								import(
									/* webpackChunkName: "addTag" */ "../views/setting/tagMg/add.vue"
								),
							meta: {
								footer: false,
								role: [roles.admin],
								menu: settingPageMenus
							}
						},
						{
							path: "tag/:id",
							name: "editTag",
							component: () =>
								import(
									/* webpackChunkName: "addTag" */ "../views/setting/tagMg/add.vue"
								),
							props: true,
							meta: {
								role: [roles.admin],
								footer: false,
								menu: settingPageMenus
							}
						},
						{
							path: "thumbnailMg",
							name: "thumbnailMg",
							component: () =>
								import(
									/* webpackChunkName: "thumbnailMg" */ "../views/setting/thumbnailMg/index.vue"
								),
							meta: {
								footer: false,
								role: [roles.admin],
								menu: settingPageMenus
							}
						}
					]
				},
				{
					path: "comment",
					name: "comment",
					component: () =>
						import(
							/* webpackChunkName: "comment" */ "../views/comment/index.vue"
						),
					meta: {
						role: [roles.admin, roles.tourist]
					}
				},
				{
					path: "admin",
					name: "admin",
					component: () =>
						import(
							/* webpackChunkName: "admin" */ "../views/admin/index.vue"
						),
					meta: {
						role: [roles.admin]
					}
				},
				{
					path: "changepwd",
					name: "changepwd",
					component: () =>
						import(
							/* webpackChunkName: "changepwd" */ "../views/changePwd/index.vue"
						),
					meta: {
						footer: false,
						role: [roles.admin]
					}
				},
				{
					path: "error",
					name: "error",
					component: () =>
						import(
							/* webpackChunkName: "error" */ "../views/error/index.vue"
						)
				}
			]
		},
		{
			path: "*",
			component: () =>
				import(
					/* webpackChunkName: "layout" */ "../components/layout.vue"
				),
			redirect: "/notFound",
			children: [
				{
					path: "/notFound",
					name: "/notFound",
					component: () =>
						import(
							/* webpackChunkName: "notFound" */ "../views/404/index.vue"
						)
				}
			]
		}
	]
});

router.beforeEach((to, from, next) => {
	const userModule = store.state.user;
	const token = userModule.token;
	const role = userModule.role;
	if (!token && to.meta.role) {
		// 需要重新登陆
		next(loginPath);
	} else if (token && to.path === loginPath) {
		next(from.path || "/");
	} else {
		if (to.meta.role && !to.meta.role.includes(role)) {
			// 使用 next() 会报错 vue-router 希望每个导航操作只有一个重定向
			router.push(from.path)
		} else {
			next();
		}
	}
});

export default router;
