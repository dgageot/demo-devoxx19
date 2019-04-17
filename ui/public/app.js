"use strict";

angular.module('quizz', [])

.controller('QuizzCtrl', function($scope, $http) {
  $http.get('/title').then(function(reponse) {
    $scope.title = reponse.data;
  }, function errorCallback(response) {
    $scope.error = true
  });

  $http.get('/game').then(function(reponse) {
    var game = reponse.data;

    $scope.guess = game.guess;
    $scope.flavor = game.flavor
    $scope.playAgain = false;

    var win = function() {
      if ($scope.playAgain) return;
      $scope.correct = true;
      $scope.wrong = false;
      $scope.playAgain = true;
    };
    var lose = function() {
      if ($scope.playAgain) return;
      $scope.correct = false;
      $scope.wrong = true;
      $scope.playAgain = true;
    };

    if (Math.random() >= 0.5) {
      $scope.candidate1 = game.guess;
      $scope.candidate2 = game.choice;
      $scope.playLeft = win;
      $scope.playRight = lose;
    } else {
      $scope.candidate1 = game.choice;
      $scope.candidate2 = game.guess;
      $scope.playLeft = lose;
      $scope.playRight = win;
    }
  }, function errorCallback(response) {
    $scope.error = true
  });
})
