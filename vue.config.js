module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  publicPath: "./",
  devServer: {
    port: 80,
    proxy: {
      '/v1/': {
        target: 'http://182.92.0.107:8081',
        // target: 'http://api.live.bilibili.com/room',
        changeOrigin: true,
      }
    }
  }
}