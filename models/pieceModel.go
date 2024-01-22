package model

import "fmt"

type Color string
type PieceType string

const (
    Black Color = "black"
    White Color = "white"
    None  Color = "none"
)

const (
    Pawn   PieceType = "pawn"
    Knight PieceType = "knight"
    Bishop PieceType = "bishop"
    Rook   PieceType = "rook"
    Queen  PieceType = "queen"
    King   PieceType = "king"
    Empty  PieceType = "empty"
)

type Piece struct {
    Type  PieceType
    Color Color
}

func (p Piece) String() string {
    if p.Type == Empty {
        return "  "
    }
    return fmt.Sprintf("%s:%s", p.Color, p.Type)
}
