// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func UiAmountInput() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"w-1/2 flex flex-col gap-2\"><h1 class=\"text-gray-600\">Amount</h1><div class=\"flex items-center\"><button type=\"button\" class=\"w-8 h-8 bg-white border border-gray-800 text-gray-800 text-lg font-bold flex justify-center items-center hover:opacity-80\" onclick=\"changeQuantity(-1)\">-</button> <input type=\"number\" id=\"amount\" name=\"amount\" class=\"w-12 h-8 bg-white text-center text-lg border border-white\" value=\"1\" min=\"1\" hx-get=\"/api/update-product-info\" hx-trigger=\"input changed, load\" hx-target=\"#dummy-target\" hx-include=\"[name], #amount\"> <button type=\"button\" class=\"w-8 h-8 bg-white border border-gray-800 text-gray-800 text-lg font-bold flex justify-center items-center hover:opacity-80\" onclick=\"changeQuantity(+1)\">+</button></div></div><script>\r\n      function changeQuantity(amount){\r\n        let input = document.getElementById(\"amount\");\r\n        let newValue = parseInt(input.value) + amount;\r\n        if (newValue >= 1) {\r\n          input.value = newValue;\r\n        }\r\n        input.dispatchEvent(new Event(\"input\", { bubbles: true }));\r\n      }\r\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
