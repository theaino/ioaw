const path = require("path");
const webpack = require("webpack");

module.exports = {
	context: __dirname,
	entry: "./assets/js/index",
	module: {
		rules: [
			{
				test: /\.s[ac]ss$/i,
				use: ["style-loader", "css-loader", "sass-loader"],
			},
		],
	},
	output: {
		path: path.resolve(__dirname, "dist/"),
		publicPath: "auto",
		filename: "bundle.js",
	},
	plugins: [
		new webpack.ProvidePlugin({
			$: "jquery",
			jQuery: "jquery",
		}),
	],
	cache: {
		type: "filesystem",
		allowCollectingMemory: true,
	},
};
