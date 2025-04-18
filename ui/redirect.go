package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/troygilman/gong"
)

type RedirectComponent string

func (c RedirectComponent) View() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return gong.Redirect(ctx, string(c))
	})
}
