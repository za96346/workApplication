const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const webpack = require("webpack")
const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');
// const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = {
  entry: './src/index.tsx',
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
    plugins:[
      new TsconfigPathsPlugin({
          configFile: 'tsconfig.json',
          extensions: ['.ts', '.js', 'tsx', 'jsx'],
          baseUrl: './src'
      }),
    ],
  },
  output: {
    path: path.join(__dirname, '/dist'),
    filename: 'bundle.min.js'
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
      {
        loader: 'css-loader',
        test: /\.css$/i,
        options: {
          modules: { localIdentName: '[name]__[local]___[hash:base64:5]' },
        },
      },
      {
        test: /\.scss$/i,
        use: [MiniCssExtractPlugin.loader, "css-loader", "sass-loader"],
      },
      {
        test: /\.(png|jpe?g|gif)$/i,
        use: [
          {
            loader: 'file-loader',
          },
        ],
      },
    ]
  },
  plugins: [
    // new HtmlWebpackPlugin({
    //   template: './src/index.html'
    // }),
    new MiniCssExtractPlugin({
      filename: 'index.css',
    }),
    new webpack.DefinePlugin({
      PRODUCTION: JSON.stringify(true),
      VERSION: JSON.stringify('5fa3b9'),
      BROWSER_SUPPORTS_HTML5: true,
      TWO: '1+1',
      'typeof window': JSON.stringify('object'),
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV),
    }),
  ],
};