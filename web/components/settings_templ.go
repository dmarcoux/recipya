// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/reaper47/recipya/internal/templates"
)

func Settings(data templates.Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if data.IsHxRequest {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title hx-swap-oob=\"true\">Settings | Recipya</title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = settings(data).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Err = settings(data).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = layoutMain("Settings", data).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func settings(data templates.Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid place-content-center md:place-content-stretch md:grid-flow-col md:h-full\" style=\"grid-template-columns: min-content\"><div class=\"hidden md:grid text-sm md:text-base bg-gray-200 max-w-[6rem] mt-[1px] dark:bg-gray-600 dark:border-r dark:border-r-gray-500\" role=\"tablist\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !data.IsAutologin {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"px-2 bg-gray-300 hover:bg-gray-300 dark:bg-gray-800 dark:hover:bg-gray-800\" hx-get=\"/settings/tabs/profile\" hx-target=\"#settings-tab-content\" role=\"tab\" aria-selected=\"false\" aria-controls=\"tab-content\" _=\"on click remove .bg-gray-300 .dark:bg-gray-800 from &lt;div[role=&#39;tablist&#39;] button/&gt; then add .bg-gray-300 .dark:bg-gray-800\">Profile</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		var templ_7745c5c3_Var4 = []any{"px-2 hover:bg-gray-300 dark:hover:bg-gray-800", templ.KV("bg-gray-300", data.IsAutologin)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var4).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-get=\"/settings/tabs/recipes\" hx-target=\"#settings-tab-content\" role=\"tab\" aria-selected=\"false\" aria-controls=\"tab-content\" _=\"on click remove .bg-gray-300 .dark:bg-gray-800 from &lt;div[role=&#39;tablist&#39;] button/&gt; then add .bg-gray-300 .dark:bg-gray-800\">Recipes</button> <button class=\"px-2 hover:bg-gray-300\" hx-get=\"/settings/tabs/advanced\" hx-target=\"#settings-tab-content\" role=\"tab\" aria-selected=\"false\" aria-controls=\"tab-content\" _=\"on click remove .bg-gray-300 .dark:bg-gray-800 from &lt;div[role=&#39;tablist&#39;] button/&gt; then add .bg-gray-300 .dark:bg-gray-800\">Advanced</button></div><div id=\"settings_bottom_tabs\" class=\"btm-nav btm-nav-sm z-20 md:hidden\" _=\"on click remove .active from &lt;button/&gt; in settings_bottom_tabs then add .active to event.srcElement\"><button class=\"active\" hx-get=\"/settings/tabs/profile\" hx-target=\"#settings-tab-content\">Profile</button> <button hx-get=\"/settings/tabs/recipes\" hx-target=\"#settings-tab-content\">Recipes</button> <button hx-get=\"/settings/tabs/advanced\" hx-target=\"#settings-tab-content\">Advanced</button></div><div id=\"settings-tab-content\" role=\"tabpanel\" class=\"w-[90vw] text-sm md:text-base p-4 auto-rows-min md:w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.IsAutologin {
			templ_7745c5c3_Err = SettingsTabsRecipes(data.Settings).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = SettingsTabsProfile().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SettingsTabsAdvanced(data templates.SettingsData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mb-2 md:grid md:grid-cols-2 md:gap-4\"><p class=\"mb-1 font-semibold select-none md:text-end md:mb-0\">Restore backup:</p><form class=\"grid gap-1 grid-flow-col w-fit\" hx-post=\"/settings/backups/restore\" hx-include=\"select[name=&#39;date&#39;]\" hx-swap=\"none\" hx-indicator=\"#fullscreen-loader\" hx-confirm=\"Continue with this backup? Today&#39;s data will be backed up if not already done.\"><label><select required id=\"file-type\" name=\"date\" class=\"select select-bordered select-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, b := range data.Backups {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(b.Value))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" selected>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(b.Display)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `settings.templ`, Line: 95, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label> <button class=\"btn btn-sm btn-outline\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = iconRocketLaunch().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SettingsTabsRecipes(data templates.SettingsData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mb-4 md:mb-2 md:grid md:grid-cols-2 md:gap-4\"><p class=\"mb-1 font-semibold md:text-end\">Export data:<br><span class=\"font-light text-sm\">Download your recipes in the selected file format.</span></p><form class=\"grid gap-1 grid-flow-col w-fit\" hx-get=\"/settings/export/recipes\" hx-include=\"select[name=&#39;type&#39;]\" hx-swap=\"none\"><label class=\"form-control w-full max-w-xs\"><select required id=\"file-type\" name=\"type\" class=\"w-fit select select-bordered select-sm\"><optgroup label=\"Recipes\"><option value=\"json\" selected>JSON</option> <option value=\"pdf\">PDF</option></optgroup></select></label> <button class=\"btn btn-outline btn-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = iconDownload().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></form></div><div class=\"mb-4 md:mb-2 md:grid md:grid-cols-2 md:gap-4\"><p class=\"mb-1 font-semibold md:text-end\">Measurement system:</p><label class=\"form-control w-full max-w-xs\"><select name=\"system\" hx-post=\"/settings/measurement-system\" hx-swap=\"none\" class=\"w-fit select select-bordered select-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, system := range data.MeasurementSystems {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(system.String()))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if system == data.UserSettings.MeasurementSystem {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(system.String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `settings.templ`, Line: 132, Col: 116}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label></div><div class=\"flex mb-4 md:mb-2 md:grid md:grid-cols-2 md:gap-4\"><label for=\"convert\" class=\"mb-1 font-semibold md:text-end\">Convert automatically:<br><span class=\"font-light text-sm\">Convert new recipes to your preferred measurement system.</span></label> <input type=\"checkbox\" name=\"convert\" id=\"convert\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.UserSettings.ConvertAutomatically {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"checkbox\" hx-post=\"/settings/convert-automatically\" hx-trigger=\"click\"></div><div class=\"flex mb-4 md:mb-2 md:grid md:grid-cols-2 md:gap-4\"><label for=\"calculate-nutrition\" class=\"mb-1 font-semibold md:text-end\">Calculate nutrition facts:<br><span class=\"font-light text-sm md:max-w-96 md:inline-block\">Calculate the nutrition facts automatically when adding a recipe. The processing will be done in the background.</span></label> <input id=\"calculate-nutrition\" type=\"checkbox\" name=\"calculate-nutrition\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.UserSettings.CalculateNutritionFact {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"checkbox\" hx-post=\"/settings/calculate-nutrition\" hx-trigger=\"click\"></div><div class=\"md:grid md:grid-cols-2 md:gap-4\"><label class=\"font-semibold md:text-end\">Integrations:<br><span class=\"font-light text-sm\">Import recipes from the selected solution.</span></label><div class=\"grid gap-1 grid-flow-col w-fit h-fit mt-1 md:mt-0\"><label class=\"form-control w-full max-w-xs\"><select name=\"integrations\" class=\"w-fit select select-bordered select-sm\"><option value=\"nextcloud\" selected>Nextcloud</option></select></label> <button class=\"btn btn-outline btn-sm\" onmousedown=\"integrations_nextcloud_dialog.showModal()\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = iconDownload2().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div></div><dialog id=\"integrations_nextcloud_dialog\" class=\"modal\"><div class=\"modal-box\"><form method=\"dialog\"><button class=\"btn btn-sm btn-circle btn-ghost absolute right-2 top-2\">✕</button></form><h3 class=\"font-semibold underline text-center\">Import from Nextcloud</h3><form method=\"dialog\" hx-post=\"/integrations/import/nextcloud\" hx-swap=\"none\" onsubmit=\"integrations_nextcloud_dialog.close()\"><label class=\"form-control w-full\"><div class=\"label\"><span class=\"label-text font-medium\">Nextcloud URL</span></div><input type=\"url\" name=\"url\" placeholder=\"https://nextcloud.mydomain.com\" class=\"input input-bordered w-full\" required></label> <label class=\"form-control w-full\"><div class=\"label\"><span class=\"label-text font-medium\">Username</span></div><input type=\"text\" name=\"username\" placeholder=\"Enter your Nextcloud username\" class=\"input input-bordered w-full\" required></label> <label class=\"form-control w-full pb-2\"><div class=\"label\"><span class=\"label-text font-medium\">Password</span></div><input type=\"password\" name=\"password\" placeholder=\"Enter your Nextcloud password\" class=\"input input-bordered w-full\" required></label> <button class=\"btn btn-block btn-primary btn-sm mt-2\">Import</button></form></div></dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SettingsTabsProfile() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mb-4 md:mb-2 md:grid md:grid-cols-2 md:gap-4\"><p class=\"mb-1 font-semibold md:text-end\">Change password:</p><div class=\"card card-bordered card-compact w-96 bg-base-100 max-w-xs\"><div class=\"card-body pt-2\"><form hx-post=\"/auth/change-password\" hx-indicator=\"#fullscreen-loader\" hx-swap=\"none\"><label class=\"form-control w-full\"><div class=\"label\"><span class=\"label-text\">Current password?</span></div><input type=\"password\" placeholder=\"Enter current password\" class=\"input input-bordered input-sm w-full max-w-xs\" name=\"password-current\" required></label> <label class=\"form-control w-full\"><div class=\"label\"><span class=\"label-text\">New password?</span></div><input type=\"password\" placeholder=\"Enter new password\" class=\"input input-bordered input-sm w-full max-w-xs\" name=\"password-new\" required></label> <label class=\"form-control w-full\"><div class=\"label\"><span class=\"label-text\">Confirm password?</span></div><input type=\"password\" placeholder=\"Retype new password\" class=\"input input-bordered input-sm w-full max-w-xs\" name=\"password-confirm\" required></label><div type=\"submit\" class=\"card-actions justify-end mt-2\"><button class=\"btn btn-primary btn-block btn-sm\">Update</button></div></form></div></div></div><div class=\"mb-2 grid grid-cols-2 gap-4\"><p class=\"mb-1 font-semibold md:text-end\">Delete Account:<br><span class=\"font-light text-sm\">This will delete all your data.</span></p><button type=\"submit\" hx-delete=\"/auth/user\" hx-confirm=\"Are you sure you want to delete your account? This action is irreversible.\" class=\"btn btn-error w-28\">Delete</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}