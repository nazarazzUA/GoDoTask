var webpack = require('webpack');

module.exports = {
    watch : true,
    context: __dirname + "/src",
    entry: {
        admin : './js/index'
    },
    output: {
        path: __dirname + "/dist/",
        filename: "[name]_bundle.js"
    },
    plugins: [
        new webpack.ProvidePlugin({
            $: "jquery", jQuery: "jquery"
        })
    ]
};