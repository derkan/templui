// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/axzilla/templui/pkg/icons"
	"github.com/axzilla/templui/pkg/utils"
	"strings"
)

type DropdownMenuItem struct {
	Label      string             // Display text
	Value      string             // Optional value
	Href       string             // Optional link URL
	Target     string             // Optional link target
	IconLeft   templ.Component    // Optional icon on the left
	IconRight  templ.Component    // Optional icon on the right
	SubItems   []DropdownMenuItem // Nested submenu items
	Disabled   bool               // Disables the item
	Attributes templ.Attributes   // Additional HTML attributesß
}

type DropdownMenuProps struct {
	Items    []DropdownMenuItem // Menu items
	Trigger  templ.Component    // Custom trigger element
	Class    string             // Additional CSS classes
	Position string             // Preferred placement
}

func (d DropdownMenuItem) modifierClasses() string {
	classes := []string{}
	if d.Disabled {
		classes = append(classes, "text-muted-foreground cursor-not-allowed")
	} else {
		classes = append(classes, "text-foreground hover:bg-accent hover:text-accent-foreground")
	}
	return strings.Join(classes, " ")
}

func renderMenuItem(item DropdownMenuItem, index int, depth int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if len(item.SubItems) > 0 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"relative group\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 = []any{
				utils.TwMerge(
					"w-full text-left flex items-center justify-between px-4 py-2 text-sm",
					item.modifierClasses(),
				),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<button class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" role=\"menuitem\" tabindex=\"-1\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("menu-item-%d", index))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 51, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" aria-haspopup=\"true\" aria-expanded=\"false\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " disabled")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, item.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "><span class=\"flex items-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconLeft != nil {
				templ_7745c5c3_Err = item.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 61, Col: 17}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconRight != nil {
				templ_7745c5c3_Err = item.IconRight.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = icons.ChevronRight(icons.IconProps{Size: "16"}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if depth < 3 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<div class=\"absolute left-full top-0 hidden group-hover:block\"><div class=\"py-1 bg-popover rounded-md shadow-lg border border-border\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for subIndex, subItem := range item.SubItems {
					templ_7745c5c3_Err = renderMenuItem(subItem, subIndex, depth+1).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if item.Href != "" {
			var templ_7745c5c3_Var6 = []any{
				"px-4 py-2 text-sm flex items-center",
				templ.KV("text-foreground hover:bg-accent hover:text-accent-foreground", !item.Disabled),
				templ.KV("text-muted-foreground cursor-not-allowed", item.Disabled),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.SafeURL = templ.SafeURL(item.Href)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" target=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(item.Target)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 82, Col: 23}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "\" role=\"menuitem\" tabindex=\"-1\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("menu-item-%d", index))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 90, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, item.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "><span class=\"flex items-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconLeft != nil {
				templ_7745c5c3_Err = item.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 97, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconRight != nil {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<span class=\"ml-auto\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = item.IconRight.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			var templ_7745c5c3_Var12 = []any{
				"w-full text-left flex items-center justify-between px-4 py-2 text-sm",
				templ.KV("text-foreground hover:bg-accent hover:text-accent-foreground", !item.Disabled),
				templ.KV("text-muted-foreground cursor-not-allowed", item.Disabled),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var12...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "<button class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var12).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "\" role=\"menuitem\" tabindex=\"-1\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("menu-item-%d", index))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 114, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, " disabled")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, item.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "><span class=\"flex items-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconLeft != nil {
				templ_7745c5c3_Err = item.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 122, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconRight != nil {
				templ_7745c5c3_Err = item.IconRight.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

func DropdownScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		handler := templ.NewOnceHandle()
		templ_7745c5c3_Var17 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 134, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "\">\n            document.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('dropdown', () => ({\n\t\t\t\t\tisOpen: false,\n\t\t\t\t\tposition: null,\n\t\t\t\t\tverticalPosition: 'bottom',\n\n\t\t\t\t\tinit() {\n\t\t\t\t\t\tthis.position = this.$el.dataset.position;\n\t\t\t\t\t},\n\n\t\t\t\t\tupdatePosition() {\n\t\t\t\t\t\tif (!this.isOpen) return;\n\n\t\t\t\t\t\tconst menu = this.$refs.menu;\n\t\t\t\t\t\tconst rect = menu.getBoundingClientRect();\n\t\t\t\t\t\tconst viewportWidth = window.innerWidth || document.documentElement.clientWidth;\n\t\t\t\t\t\tconst viewportHeight = window.innerHeight || document.documentElement.clientHeight;\n\n\t\t\t\t\t\tif (this.position === 'left' && rect.left < 0) {\n\t\t\t\t\t\t\tthis.position = 'right';\n\t\t\t\t\t\t} else if (this.position !== 'left' && rect.right > viewportWidth) {\n\t\t\t\t\t\t\tthis.position = 'left';\n\t\t\t\t\t\t}\n\n\t\t\t\t\t\tif (this.verticalPosition === 'bottom' && rect.bottom > viewportHeight) {\n\t\t\t\t\t\t\tthis.verticalPosition = 'top';\n\t\t\t\t\t\t} else if (this.verticalPosition === 'top' && rect.top < 0) {\n\t\t\t\t\t\t\tthis.verticalPosition = 'bottom';\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tgetPositionClass() {\n\t\t\t\t\t\tconst classes = [];\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (this.position === 'left') {\n\t\t\t\t\t\t\tclasses.push('right-0');\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tclasses.push('left-0');\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (this.verticalPosition === 'top') {\n\t\t\t\t\t\t\tclasses.push('bottom-full mb-2');\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tclasses.push('top-full mt-2');\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\treturn classes.join(' ');\n\t\t\t\t\t},\n\n\t\t\t\t\tsetClose() {\n\t\t\t\t\t\tthis.isOpen = false;\n\t\t\t\t\t},\n\n\t\t\t\t\tsetOpen() {\n\t\t\t\t\t\tthis.isOpen = true;\n\t\t\t\t\t},\n\n\t\t\t\t\topenMenu() {\n\t\t\t\t\t\tthis.isOpen = !this.isOpen;\n\t\t\t\t\t\tif(this.isOpen) this.$nextTick(() => this.updatePosition());\n\t\t\t\t\t}\n\t\t\t\t}))\n\t\t\t});\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handler.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var17), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

// Floating menu for displaying a list of actions or options.
func DropdownMenu(props DropdownMenuProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var19 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var19 == nil {
			templ_7745c5c3_Var19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var20 = []any{utils.TwMerge("relative inline-block text-left", props.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var20...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "<div x-data=\"dropdown\" @resize.window=\"updatePosition\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var20).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "\" data-position=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 string
		templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(props.Position)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 208, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "\"><div @click=\"openMenu\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Trigger != nil {
			templ_7745c5c3_Err = props.Trigger.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = Button(ButtonProps{
				Text:      "Options",
				Variant:   "outline",
				IconRight: icons.ChevronDown(icons.IconProps{Size: "16"}),
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var23 = []any{
			"absolute z-50 mt-2 w-56 rounded-md shadow-lg bg-popover ring-1 ring-black ring-opacity-5 focus:outline-none",
			"border border-border",
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var23...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, "<div x-ref=\"menu\" x-show=\"isOpen\" @click.away=\"setClose\" @keydown.escape.window=\"setClose\" x-transition:enter=\"transition ease-out duration-100\" x-transition:enter-start=\"transform opacity-0 scale-95\" x-transition:enter-end=\"transform opacity-100 scale-100\" x-transition:leave=\"transition ease-in duration-75\" x-transition:leave-start=\"transform opacity-100 scale-100\" x-transition:leave-end=\"transform opacity-0 scale-95\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 string
		templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var23).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, "\" x-bind:class=\"getPositionClass\" role=\"menu\" aria-orientation=\"vertical\" aria-labelledby=\"dropdown-menu-button\" tabindex=\"-1\"><div class=\"py-1\" role=\"none\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for index, item := range props.Items {
			templ_7745c5c3_Err = renderMenuItem(item, index, 0).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, "</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
