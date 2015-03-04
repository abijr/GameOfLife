'use strict';

angular.module('static')
  .controller('MainCtrl', function ($scope) {
    var x = 20;
    var y = 20;


    var world = $("#world");
    console.log(world.width());
    var width = world.width() / x;
    var height = world.width() /y;

    $scope.cellStyle = {
      "height": height+"px",
      "width": width+"px"
    }

    $scope.cells = [];
    for (var j = 0; j < y; j++) {
      $scope.cells[j] = [];
      for (var i = 0; i < x; i++) {
        $scope.cells[j][i] = false;
      }
    }

  });
