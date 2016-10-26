var module  = require('../module');

module.exports = module.directive('mainMenu',[
    '$location','THEMES_DIR',
    function($location,THEMES_DIR) {
        return {
            replace:true,
            restrict:'E',
            templateUrl: THEMES_DIR+'/core/navigation/main-menu.html',
            controller:['$scope','$rootScope', function ($scope, $rootScope) {
                $scope.menu = [
                    {"title":"Plans", "url":"/palns"},
                    {"title":"Tasks", "url":"/tasks"}
                ];
            }]
        }
    }
]);