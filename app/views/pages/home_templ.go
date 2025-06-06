// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"ecomm-go/app/views/components"
	"ecomm-go/db/goqueries"
	"fmt"
)

func Home(products []goqueries.GetProductsByCategoriesAndNameRow) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"pt-4 pb-10 flex flex-col gap-4\"><div class=\"flex md:flex-row flex-col bg-hawaiian-tan-200 rounded-lg md:py-16 py-4 px-4 lg:px-32 sm:px-16 gap-10 justify-center items-center\"><div class=\"flex flex-col md:w-2/3 pt-10 gap-4 text-center md:text-left\"><div class=\"text-hawaiian-tan-800 font-bold text-lg\">hendrerit ullamcorper. Donec </div><div class=\"font-bold text-5xl\">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</div><div class=\"text-lg py-4\">eu magna massa, scelerisque ac lobortis et, posuere eu metus. Aenean fringilla, felis id hendrerit ullamcorper, lectus leo faucibus tellus, vel ullamcorper</div><a href=\"/category\" class=\"text-2xl font-bold underline\">List products by categories</a></div><div class=\"md:w-1/3 p-4 md:p-0\"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 510))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 21, Col: 97}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\"></div></div><div class=\"grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-10\"><div class=\"rounded-lg flex bg-hunter-green-200 transform duration-300 hover:scale-105 hover:shadow-xl\"><div class=\"w-1/2 flex flex-col  p-10\"><div class=\"font-bold text-lg\">scelerisque ac lobortis</div><div class=\"text-md\">scelerisque ac lobortis et</div><div class=\"font-bold text-lg\">scelerisque </div></div><div class=\"w-1/2 \"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 501))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 33, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" class=\"object-cover rounded-r-lg w-full h-full\"></div></div><div class=\"rounded-lg flex bg-timber-green-200 transform duration-300 hover:scale-105 hover:shadow-xl\"><div class=\"w-1/2 flex flex-col p-10\"><div class=\"font-bold text-lg\">scelerisque ac lobortis</div><div class=\"text-md\">scelerisque ac lobortis et</div><div class=\"font-bold text-lg\">scelerisque </div></div><div class=\"w-1/2 \"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 502))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 43, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" class=\"object-cover rounded-r-lg w-full h-full\"></div></div><div class=\"rounded-lg flex bg-heavy-metal-200 transform duration-300 hover:scale-105 hover:shadow-xl\"><div class=\"w-1/2 flex flex-col p-10\"><div class=\"font-bold text-lg\">scelerisque ac lobortis</div><div class=\"text-md\">scelerisque ac lobortis et</div><div class=\"font-bold text-lg\">scelerisque </div></div><div class=\"w-1/2 h-full\"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 503))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 53, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" class=\"object-cover rounded-r-lg w-full h-full\"></div></div></div><div class=\"mb-10\"><div class=\"font-bold text-3xl mb-4\">New Products</div><div class=\"flex flex-wrap md:justify-normal justify-center transform duration-500 gap-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ProductList(products, 1, false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div></div><div class=\"grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-10\"><div class=\"rounded-lg flex bg-timber-green-200 transform duration-300 hover:scale-105 hover:shadow-xl\"><div class=\"w-1/2 flex flex-col  p-10\"><div class=\"font-bold text-lg\">scelerisque ac lobortis</div><div class=\"text-md\">scelerisque ac lobortis et</div><div class=\"font-bold text-lg\">scelerisque </div></div><div class=\"w-1/2 \"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 504))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 76, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" class=\"object-cover rounded-r-lg w-full h-full\"></div></div><div class=\"rounded-lg flex bg-heavy-metal-200 transform duration-300 hover:scale-105 hover:shadow-xl\"><div class=\"w-1/2 flex flex-col p-10\"><div class=\"font-bold text-lg\">scelerisque ac lobortis</div><div class=\"text-md\">scelerisque ac lobortis et</div><div class=\"font-bold text-lg\">scelerisque </div></div><div class=\"w-1/2 \"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 505))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 86, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" class=\"object-cover rounded-r-lg w-full h-full\"></div></div><div class=\"rounded-lg flex bg-hunter-green-200 transform duration-300 hover:scale-105 hover:shadow-xl\"><div class=\"w-1/2 flex flex-col p-10\"><div class=\"font-bold text-lg\">scelerisque ac lobortis</div><div class=\"text-md\">scelerisque ac lobortis et</div><div class=\"font-bold text-lg\">scelerisque </div></div><div class=\"w-1/2 h-full\"><img alt=\"Product\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("https://picsum.photos/seed/%d/600/600", 506))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/pages/home.templ`, Line: 96, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" class=\"object-cover rounded-r-lg w-full h-full\"></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
