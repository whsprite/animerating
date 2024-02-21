module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  publicPath: "./",
  devServer: {
    port: 80,
    proxy: {
      '/v1/': {
        target: 'your sever ip address',
        changeOrigin: true,
      }
    }
  }
}
