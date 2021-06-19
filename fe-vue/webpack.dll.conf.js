const path = require("path");
const webpack = require("webpack");

module.exports = {
  mode: "production",
  // 想要打包的模块的数组
  entry: {
    vendor: ["axios", "vue", "vuex", "js-cookie", "vue-router"]
  },
  output: {
    // 打包后文件输出的位置
    path: path.resolve(__dirname, "public"),
    // 生成的文件名字 默认为vendor.dll.js
    filename: "vendor.dll.js",
    // 生成文件的映射关系，与下面的 DLLPlugin 配置相对应
    library: "vendor_library"
  },
  plugins: [
    new webpack.DllPlugin({
      // 生成一个json文件 里面是关于dll.js的配置信息
      path: path.join(__dirname, "public", "vendor-manifest.json"),
      // 与上面output中的配置对应
      name: "vendor_library",
      // 上下文环境路径，必须填写，为了与DLLReferencePlugin存在于同一上下文中，否则undefined
      context: path.join(__dirname, "public")
    })
  ]
};
