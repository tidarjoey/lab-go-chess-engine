package model

import "fmt"

type Board [8][8]Piece

func NewBoard() *Board {
    var board Board
    // Initialize the board with empty pieces
    for i := range board {
        for j := range board[i] {
            board[i][j] = Piece{Type: Empty, Color: None}
        }
    }
    board[0][0] = Piece{
        Type: Rook,
        Color: Black,
    }
    board[0][7] = Piece{
        Type: Rook,
        Color: Black,
    }
    return &board
}

func (b *Board) Print() {
    for _, row := range b {
        for _, piece := range row {
            fmt.Printf("[%s] ", piece)
        }
        fmt.Println()
    }
}
