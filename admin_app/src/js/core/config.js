var core = require('./module');

require('./directives/mainMenu');

module.exports = core.config([
    '$stateProvider','THEMES_DIR',
    function($stateProvider,THEMES_DIR) {

        $stateProvider.state('homepage',{
            url:"/",
            templateUrl: THEMES_DIR + '/core/main-page.html'
        });
    }
]);