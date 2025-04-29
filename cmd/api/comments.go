package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pedrogawa/social-go/internal/store"
)

type CommentPayload struct {
	Content string `json:"content" validate:"max=255"`
}

// CreateComment godoc
//
//	@Summary		Creates a comment
//	@Description	Creates a comment in a post
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//
//	@Param			postID	path	int				true	"Post ID"
//
//	@Param			payload	body	CommentPayload	true	"Comment payload"
//
//	@Success		200
//
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID}/comments [post]
func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")

	ctx := r.Context()

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var payload CommentPayload

	err = readJSON(w, r, &payload)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.store.Comments.Create(ctx, &store.Comment{
		UserID:  1303,
		PostID:  id,
		Content: payload.Content,
	})

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
}
