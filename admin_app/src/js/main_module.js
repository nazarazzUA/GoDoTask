
require('angular');
require('angular-resource');
require('angular-ui-router');
require('angular-file-upload');
require('jquery');

require ('./core/config');

module.exports = function() {

    var app = angular.module('nz-application',[
        'ngResource',
        'nz-core'
    ]);

    app.config(['$locationProvider','$httpProvider',
        function($locationProvider,$httpProvider) {
            $locationProvider.html5Mode(true);
            $httpProvider.defaults.headers.common["X-Requested-With"] = "GoDoTaskAdminApp";
//            $httpProvider.interceptors.push('HttpInterceptor');
        }]);

    app.run(['$rootScope','$http','$location','$window',
        function($rootScope,$http, $location,$window){

            console.log("Root application is start")

        }

    ]);
    
    return app;
};