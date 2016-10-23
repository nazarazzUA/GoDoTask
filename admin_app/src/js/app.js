var main = require('./main_module.js');

module.exports = function (){
    var app = main();
    angular.bootstrap(document.documentElement,[app.name]);
};

