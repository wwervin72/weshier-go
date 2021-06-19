import vue from 'vue'
import vuex from 'vuex'
import getters from './getters'

vue.use(vuex)

const modulesFiles = require.context("./modules", true, /\.js$/);

const modules = modulesFiles.keys().reduce((modules, modulePath) => {
  const moduleName = modulePath.replace(/^\.\/(.*)\.\w+$/, "$1");
  const value = modulesFiles(modulePath);
  modules[moduleName] = value.default;
  return modules;
}, {});

const store = new vuex.Store({
	modules,
	getters
  });

  export default store;
