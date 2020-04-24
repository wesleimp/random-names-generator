const { join } = require("path")
const HtmlWebPackPlugin = require('html-webpack-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')

const babelLoader = {
  test: /\.js$/,
  exclude: /node_modules/,
  use: { loader: 'babel-loader' }
}

const fileLoader = {
  test: /\.(eot|svg|otf|ttf|woff|woff2)$/,
  use: 'file-loader'
}

const htmlLoader = {
  test: /\.html$/,
  use: { loader: 'html-loader' }
}

module.exports = {
 entry: join(__dirname, 'src', 'index.js'),
  output: {
    path: join(__dirname, 'dist'),
    filename: 'main.[hash:8].js',
    chunkFilename: '[name].js',
  },
  optimization: {
    namedModules: true,
    splitChunks: {
      chunks: "all",
      cacheGroups: {
        default: false,
        vendors: false,

        vendor: {
          chunks: "all",
          test: /[\\/]node_modules[\\/]/,
          name: "vendor",
          priority: 20
        }
      }
    }
  },
  devtool: 'source-map',
  module: {
    rules: [babelLoader, fileLoader, htmlLoader]
  },
  devServer: {
    inline: true,
    port: 3005,
    contentBase: join(__dirname, 'public'),
    historyApiFallback: true
  },
  plugins: [
    new CleanWebpackPlugin(),
    new HtmlWebPackPlugin({
      inject: true,
      template: join(__dirname, "public", "index.html"),
    }),
  ],
}