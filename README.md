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
8 [ ♖ ][ ♘ ][ ♗ ][ ♕ ][ ♔ ][ ♗ ][ ♘ ][ ♖ ] 8
7 [ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ][ ♙ ] 7
6 [   ][   ][   ][   ][   ][   ][   ][   ] 6
5 [   ][   ][   ][   ][   ][   ][   ][   ] 5
4 [   ][   ][   ][   ][   ][   ][   ][   ] 4
3 [   ][   ][   ][   ][   ][   ][   ][   ] 3
2 [ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ][ ♟ ] 2
1 [ ♜ ][ ♞ ][ ♝ ][ ♛ ][ ♚ ][ ♝ ][ ♞ ][ ♜ ] 1
    a    b    c    d    e    f    g    h  
```

### Move Validations
- **Responsibility**: TManages the rules of the chess game. It should validate whether moves are legal and update the game state accordingly.

## Controllers

### User Input
- **Responsibility**: The controller layer will handle user interactions. In this case, it's responsible for taking user input, interpreting it, and passing it on to the model layer.
- **Console input**: receives straightforward `start` and `end` position like "e2 e4", which means move piece in `e2` to `e4`.