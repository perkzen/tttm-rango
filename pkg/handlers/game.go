package handlers

import (
	"fmt"
	tictactoe "github.com/perkzen/tttm-go/pkg/game"
	"net/http"
	"net/url"
	"strconv"
)

type GetMoveQueryParams struct {
	Gid     string
	Size    int
	Playing tictactoe.Symbol
	Moves   string
}

func NewGetMoveQueryParams(query url.Values) (*GetMoveQueryParams, error) {
	size, err := strconv.Atoi(query.Get("size"))
	if err != nil {
		size = 3
	}

	player, err := tictactoe.ToPlayerSymbol(query.Get("playing"))
	if err != nil {
		return nil, err
	}

	return &GetMoveQueryParams{
		Gid:     query.Get("gid"),
		Size:    size,
		Playing: player,
		Moves:   query.Get("moves"),
	}, nil
}

func HandleGetMove(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	moveQueryParams, err := NewGetMoveQueryParams(query)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error:Sorry mate, can't do it bro."))
		return
	}

	game := tictactoe.NewGame(moveQueryParams.Gid, moveQueryParams.Size, moveQueryParams.Moves)
	row, col := game.GetRandomMove(moveQueryParams.Playing)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Move:%s_%d_%d", moveQueryParams.Playing, row, col)))
}
