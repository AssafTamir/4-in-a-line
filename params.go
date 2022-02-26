package main

const numOfGames = 1000

const gameOver = "Game over"
const colorReset = string("\033[0m")
const colorRed = string("\033[31m")
const colorGreen = string("\033[32m")
const colorYellow = string("\033[33m")
const colorBlue = string("\033[34m")
const colorPurple = string("\033[35m")
const colorCyan = string("\033[36m")
const boardSize = 8

var X = &Token{"X", colorBlue}
var O = &Token{"O", colorRed}
var W = &Token{"W", colorGreen}
var D = &Token{"D", colorYellow}

var OpenPosition = &Token{" ", colorReset}
