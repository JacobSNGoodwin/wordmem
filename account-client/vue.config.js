module.exports = {
  devServer: {
    port: 3000,
    proxy: {
      "^/": {
        target: "http://wordmem.dev",
        ws: true,
        changeOrigin: true
      }
    }
  }
};
