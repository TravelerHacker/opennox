//go:build !server

package opennox

import (
	"context"
	"strings"

	"github.com/opennox/libs/console"
)

func init() {
	noxConsole.Register(&console.Command{
		Token: "unstick",
		HelpID: "unstickhelp",
		Flags:  console.Client,
		Func: func(ctx context.Context, c *console.Console, tokens []string) bool {
			if len(tokens) != 0 {
				return false
			}
			actions := clientRecoverStuckInput()
			if len(actions) == 0 {
				c.Println(console.ColorLightYellow, "no known stuck-input flags were set")
				return true
			}
			c.Println(console.ColorLightGreen, "recovered: "+strings.Join(actions, "; "))
			return true
		},
	})
}
