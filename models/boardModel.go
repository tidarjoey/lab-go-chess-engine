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

    placeNonPawns(&board)
    placePawns(&board)

    return &board
}

func (board *Board) Print() {
    printColumnLabels()

    for i, row := range board {
        // Print row number at the start of each row
        fmt.Printf("%d ", 8-i)

        for _, piece := range row {
            fmt.Printf("[%s]", piece.Unicode())
        }

        fmt.Println()
    }
}

func placeNonPawns(board *Board) {
    // Black Pieces
    board[0][0] = Piece{
        Type: Rook,
        Color: Black,
    }
    board[0][7] = Piece{
        Type: Rook,
        Color: Black,
    }
    board[0][1] = Piece{
        Type: Knight,
        Color: Black,
    }
    board[0][6] = Piece{
        Type: Knight,
        Color: Black,
    }
    board[0][2] = Piece{
        Type: Bishop,
        Color: Black,
    }
    board[0][5] = Piece{
        Type: Bishop,
        Color: Black,
    }
    board[0][3] = Piece{
        Type: Queen,
        Color: Black,
    }
    board[0][4] = Piece{
        Type: King,
        Color: Black,
    }

    // White Pieces
    board[7][0] = Piece{
        Type: Rook,
        Color: White,
    }
    board[7][7] = Piece{
        Type: Rook,
        Color: White,
    }
    board[7][1] = Piece{
        Type: Knight,
        Color: White,
    }
    board[7][6] = Piece{
        Type: Knight,
        Color: White,
    }
    board[7][2] = Piece{
        Type: Bishop,
        Color: White,
    }
    board[7][5] = Piece{
        Type: Bishop,
        Color: White,
    }
    board[7][3] = Piece{
        Type: Queen,
        Color: White,
    }
    board[7][4] = Piece{
        Type: King,
        Color: White,
    }
}

func placePawns(board *Board) {
    for col := 0; col < 8; col++ {
        board[1][col] = Piece{Type: Pawn, Color: Black}
        board[6][col] = Piece{Type: Pawn, Color: White}
    }
}

func printColumnLabels() {
    fmt.Print("  ")
    for col := 'a'; col <= 'h'; col++ {
        fmt.Printf("  %c  ", col)
    }
    fmt.Println()
}