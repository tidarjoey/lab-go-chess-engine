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

func (piece Piece) Unicode() string {
	if piece.Color == White {
		switch piece.Type {
		case Pawn:
			return " ♟ "
		case Knight:
			return " ♞ "
		case Bishop:
			return " ♝ "
		case Rook:
			return " ♜ "
		case Queen:
			return " ♛ "
		case King:
			return " ♚ "
		}
	} else if piece.Color == Black {
		switch piece.Type {
		case Pawn:
			return " ♙ "
		case Knight:
			return " ♘ "
		case Bishop:
			return " ♗ "
		case Rook:
			return " ♖ "
		case Queen:
			return " ♕ "
		case King:
			return " ♔ "
		}
	}
	return "   " // Space for empty square
}

func (piece Piece) String() string {
	if piece.Type == Empty {
		return "  "
	}
	return fmt.Sprintf("%s:%s", piece.Color, piece.Type)
}
