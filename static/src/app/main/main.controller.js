'use strict';

angular.module('static')
  .controller('MainCtrl', function ($scope, $http) {
    // world size
    var x = 20;
    var y = 20;

    // Make cells fit the width
    var world = angular.element('#world');
    var width = world.width() / x;
    var height = world.width() /y;

    // and configure the cell style
    $scope.cellStyle = {
      'height': height+'px',
      'width': width+'px'
    };

    // Initialize the cell array, a.k.a. world
    $scope.cells = [];
    for (var j = 0; j < y; j++) {
      $scope.cells[j] = [];
      for (var i = 0; i < x; i++) {
        $scope.cells[j][i] = false;
      }
    }

    $scope.play = function () {
      $http.post('/play', $scope.cells).
      success(function (data) {
        $scope.cells = data;
      });
    };
  });
