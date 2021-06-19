const path = require("path");
const webpack = require("webpack");
const CompressionPlugin = require("compression-webpack-plugin");
const proxyHost = 'http://localhost:3333/'

function resolve(dir) {
	return path.join(__dirname, dir);
}

module.exports = {
  devServer: {
    port: '3000',
    host: '0.0.0.0',
    compress: true,
    disableHostCheck: true,
    overlay: true,
    headers: {
      'X-Custom-Foo': 'bar',
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, PATCH, OPTIONS',
      'Access-Control-Allow-Headers':
        'X-Requested-With, content-type, Authorization'
    },
    historyApiFallback: {
      verbose: true
    },
	proxy: {
		'': {
		  target: proxyHost,
		  changOrigin: true,
		  headers: {
			origin: proxyHost
		  }
		}
	  }
  },
  configureWebpack: {
    output: {
      // 打包后文件输出的位置
      path: path.join(__dirname, "dist"),
      // 生成文件的映射关系，与下面的DLLPlugin配置相对应
      library: "[name]_library"
    },
    resolve: {
      alias: {
        "@": resolve("src")
      }
    },
    plugins: [
      new webpack.DllReferencePlugin({
        // 与 DllPlugin 中的 context 保持一致
        context: path.join(__dirname, "public"),
        name: "vendor_library",
        /* 这个地址对应webpack.dll.conf.js中生成的那个json文件的路径，这样webpack打包的时候
        会检测当前文件中的映射，不会把已经存在映射的包再次打包进bundle.js */
        manifest: path.resolve(__dirname, "./public/vendor-manifest.json")
      }),
      new CompressionPlugin({
        algorithm: "gzip"
      })
    ]
  },
}
