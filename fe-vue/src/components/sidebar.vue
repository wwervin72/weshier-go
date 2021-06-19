<template>
    <a-menu
      style="width: 200px"
      :default-selected-keys="defaultSelectedKeys"
      :open-keys.sync="openKeys"
      mode="inline"
      @click="handleClick"
    >
        <template v-for="menu in menus">
            <a-menu-item v-if="!menu.children || !menu.children.length" :key="menu.path">{{menu.title}}</a-menu-item>
            <a-sub-menu v-else :key="menu.path">
                <span slot="title"><a-icon :type="menu.icon" /><span>{{menu.title}}</span></span>
                <a-menu-item v-for="cMenu in menu.children || []" :key="cMenu.path">{{cMenu.title}}</a-menu-item>
            </a-sub-menu>
        </template>
    </a-menu>
</template>

<script>
import { Menu, Icon } from 'ant-design-vue'
export default {
	name: "WsSidebar",
    components: {
        AIcon: Icon,
        AMenu: Menu,
        AMenuItem: Menu.Item,
        ASubMenu: Menu.SubMenu,
    },
	data() {
        const firstMenu = this.$route.meta.menu[0]
		return {
            openKeys: [firstMenu.path],
		}
	},
    computed: {
        menus () {
            return this.$route.meta.menu || []
        },
        defaultSelectedKeys () {
			let routeRefMenu
			this.menus.forEach(el => {
				if (el.children && el.children.length) {
					el.children.forEach(ele => {
						if (ele.path === this.$route.fullPath) {
							routeRefMenu = ele
							this.openKeys = [el.path]
						}
					})
				} else if (el.path === this.$route.fullPath) {
					routeRefMenu = el
				}
			})
            let firstMenu = this.menus[0]
			if (routeRefMenu) return [routeRefMenu.path]
            firstMenu = firstMenu && firstMenu.children ? firstMenu.children[0] : null
            return firstMenu ? [firstMenu.path] : []
        }
    },
    methods: {
        handleClick(e) {
            if (this.$route.fullPath !== e.key) {
                this.$router.push(e.key)
            }
        },
    }
}
</script>
<style lang="scss" scoped>

</style>
