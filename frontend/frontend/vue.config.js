// vue.config.js (projenin kökünde)
const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  devServer: {
    host: '0.0.0.0',
    port: 8084,      // npm run serve komutunu çalıştırdığın port
    proxy: {
      '/api': {
        target: 'http://localhost:8083',  // backend’in portu
        changeOrigin: true,
        secure: false,
        ws: true
      }
    }
  }
})
