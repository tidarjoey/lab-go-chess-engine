# Chess game engine
Using go version 1.20.5

## Data Structures

### The pieces

Representing the game pieces with a struct like this:
```go
type Piece struct {
    Type  PieceType
    Color Color
}
```

### The board

To represent 8 x 8 tiles layout

#### Two-Dimensional Slice Structure: An 8x8 slice ([][]Piece), where Piece is a custom struct representing a chess piece.
- **Pros**: Intuitive and easy to understand. Directly maps to the physical layout of a chessboard.
- **Cons**: Can be less efficient in terms of performance, especially when copying the board state.
