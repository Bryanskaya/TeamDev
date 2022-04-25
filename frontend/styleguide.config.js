const path = require('path');
const glob = require('glob');
var webpack = require('webpack');

module.exports = {
//   title: 'React Style Guide Example',
//   components: function () {
//     return glob.sync(path.resolve(__dirname, 'src/components/**/*.tsx'))
//       .filter(function (module) {
//         return /\/[A-Z]\w*\.tsx$/.test(module);
//       });
//   },
//   resolver: require('react-docgen').resolver.findAllComponentDefinitions,
//   propsParser: require('react-docgen-typescript').withDefaultConfig({ propFilter: { skipPropsWithoutDoc: true } }).parse,

  sections: [
    {
      name: 'Styleguide',
      components: 'src/components/Boxes/**/[A-Z]*.tsx',
    }
    ],

    webpackConfig: {
        // entry: {
        //     bundle: ['./index.ts'],
        // },
        entry:'./src/index.tsx',
        context: path.resolve(__dirname),
        
        output: {
          filename: 'bundle.js',
          path: path.join(__dirname, 'temp')
        },
        resolve: {
            // symlinks: false,

            alias: {
                components: path.resolve(__dirname, './src/components/'),
                styles: path.resolve(__dirname, './src/styles/'),
                postAPI: path.resolve(__dirname, './src/postAPI/'),
                functions: path.resolve(__dirname, './src/functions/'),
                context: path.resolve(__dirname, './src/context/'),
                img: path.resolve(__dirname, './src/img/'),
            },
            extensions: [".tsx", ".ts", ".js", ".jsx", '.scss'],
        },
        module: {
          rules: [
            {
              test: /\.tsx?$/,
              loader: 'ts-loader',
              exclude: /node_modules/,
            },
            {
                test: /\.s[ac]ss$/i,
                use: [
                    // Creates `style` nodes from JS strings
                    "style-loader",
                    // Translates CSS into CommonJS
                    "css-loader",
                    // Compiles Sass to CSS
                    "sass-loader",
                ],
            },
          ]
        },
    }
};
