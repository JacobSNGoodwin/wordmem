module.exports = {
  publicPath: "/account/",
  devServer: {
    // host: "0.0.0.0",
    // port: "8080",
    public: "wordmem.test",
    disableHostCheck: true,
    sockPath: "/account/sock-js"
    // watchOptions: {
    //   poll: true
    // }
  }
};
