package model

type Position struct {
	X int // X represents the column (0 to 7 corresponding to a to h)
	Y int // Y represents the row (0 to 7 corresponding to 1 to 8)
}

type Move struct {
	Piece         Piece
	Start         Position
	End           Position
	Captured      Piece // Represents a piece that was captured during the move
	IsCastling    bool
	IsEnPassant   bool
	PromotionType PieceType // For pawn promotion, indicates the piece type to which the pawn is promoted
}
