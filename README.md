# Chess game engine
Using go version 1.20.5

## Models

### The pieces

Representing the game pieces with a struct like this:

```go
type Piece struct {
    Type  PieceType
    Color Color
}
```

For console Interface, pieces representated using [unicode](https://en.wikipedia.org/wiki/Chess_symbols_in_Unicode) like these: ♖ ♕.

### The board

To represent 8 x 8 tiles layout and the state of the game.

#### Two-Dimensional Slice Structure: An 8x8 slice ([][]Piece), where Piece is a custom struct representing a chess piece.

- **Pros**: Intuitive and easy to understand. Directly maps to the physical layout of a chessboard.
- **Cons**: Can be less efficient in terms of performance, especially when copying the board state.

#### Console user interface of the board
```
    a    b    c    d    e    f    g    h  
8 [ ♖ ][ ♘ ][ ♗ ][ ♕ ][ ♔ ][ ♗ ][ ♘ ][ ♖ ]
7 [ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ]
6 [   ][   ][   ][   ][   ][   ][   ][   ]
5 [   ][   ][   ][   ][   ][   ][   ][   ]
4 [   ][   ][   ][   ][   ][   ][   ][   ]
3 [   ][   ][   ][   ][   ][   ][   ][   ]
2 [ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ]
1 [ ♜ ][ ♞ ][ ♝ ][ ♛ ][ ♚ ][ ♝ ][ ♞ ][ ♜ ]
```

### Move

Representing the game moves with these

```go
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
```

## Services

Handles smallest game logic such as validating moves, updating the game state, and any other rules of chess.

### Moves

- **Responsibility**: 
    - Manages the move rules of each piece. It should be able to return possible moves for each piece on their position.
    - Validate the end position of a move from user input.

### Game state

- **Responsibility**: Validate game states such as check and checkmates after each move

## Controllers

Handles user input and flow control.

### New game

Creates a new game state.

### Update game

#### User Input

- **Responsibility**: The controller layer will handle user interactions. In this case, it's responsible for taking user input, interpreting it, and passing it on to the model layer.
- **Console input**: receives straightforward `start` and `end` position separately like "e2" and then "e4", which means move piece in `e2` to `e4`.