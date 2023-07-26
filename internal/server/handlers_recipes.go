package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/reaper47/recipya/internal/app"
	"github.com/reaper47/recipya/internal/models"
	"github.com/reaper47/recipya/internal/scraper"
	"github.com/reaper47/recipya/internal/templates"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func recipesAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Hx-Request") == "true" {
		parsedURL, err := url.Parse(r.Header.Get("HX-Current-Url"))
		if err == nil && parsedURL.Path == "/recipes/add/unsupported-website" {
			w.Header().Set("HX-Trigger", makeToast("Website requested.", infoToast))
		}

		templates.RenderComponent(w, "recipes", "add-recipe", nil)
	} else {
		page := templates.AddRecipePage
		templates.Render(w, page, templates.Data{
			IsAuthenticated: true,
			Title:           page.Title(),
		})
	}
}

func (s *Server) recipesAddImportHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 128<<20)

	if err := r.ParseMultipartForm(128 << 20); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not parse the uploaded files.", errorToast))
		return
	}

	files, filesOk := r.MultipartForm.File["files"]
	if !filesOk {
		w.Header().Set("HX-Trigger", makeToast("Could not retrieve the files or the directory from the form.", errorToast))
		return
	}

	recipes := s.Files.ExtractRecipes(files)
	userID := r.Context().Value("userID").(int64)

	count := 0
	for _, r := range recipes {
		if _, err := s.Repository.AddRecipe(&r, userID); err != nil {
			continue
		}
		count += 1
	}

	msg := fmt.Sprintf("Imported %d recipes. %d failed", count, len(recipes)-count)
	w.Header().Set("HX-Trigger", makeToast(msg, infoToast))
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) recipeAddManualHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	categories, err := s.Repository.Categories(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData := &templates.ViewRecipeData{Categories: categories}

	if r.Header.Get("Hx-Request") == "true" {
		templates.RenderComponent(w, "recipes", "add-recipe-manual", templates.Data{
			View: viewData,
		})
	} else {
		page := templates.AddRecipeManualPage
		templates.Render(w, page, templates.Data{
			IsAuthenticated: true,
			Title:           page.Title(),
			View:            viewData,
		})
	}
}

func (s *Server) recipeAddManualPostHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 128<<20)
	if err := r.ParseMultipartForm(128 << 20); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not parse the form.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	imageFile, ok := r.MultipartForm.File["image"]
	if !ok {
		w.Header().Set("HX-Trigger", makeToast("Could not retrieve the image from the form.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, err := imageFile[0].Open()
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could open the image from the form.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()

	imageUUID, err := s.Files.UploadImage(f)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error uploading image.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ingredients := make([]string, 0)
	i := 1
	for {
		key := fmt.Sprintf("ingredient-%d", i)
		if r.Form.Has(key) {
			ingredients = append(ingredients, r.FormValue(key))
			i++
		} else {
			break
		}
	}

	instructions := make([]string, 0)
	i = 1
	for {
		key := fmt.Sprintf("instruction-%d", i)
		if r.Form.Has(key) {
			instructions = append(instructions, r.FormValue(key))
			i++
		} else {
			break
		}
	}

	times, err := models.NewTimes(r.FormValue("time-preparation"), r.FormValue("time-cooking"))
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error parsing times.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	yield, err := strconv.ParseInt(r.FormValue("yield"), 10, 16)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error parsing yield.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recipe := &models.Recipe{
		Category:     strings.ToLower(r.FormValue("category")),
		CreatedAt:    time.Time{},
		Cuisine:      "",
		Description:  r.FormValue("description"),
		Image:        imageUUID,
		Ingredients:  ingredients,
		Instructions: instructions,
		Keywords:     nil,
		Name:         r.FormValue("title"),
		Nutrition: models.Nutrition{
			Calories:           r.FormValue("calories"),
			Cholesterol:        r.FormValue("cholesterol"),
			Fiber:              r.FormValue("fiber"),
			Protein:            r.FormValue("protein"),
			SaturatedFat:       r.FormValue("saturated-fat"),
			Sodium:             r.FormValue("sodium"),
			Sugars:             r.FormValue("sugars"),
			TotalCarbohydrates: r.FormValue("total-carbohydrates"),
			TotalFat:           r.FormValue("total-fat"),
			UnsaturatedFat:     "",
		},
		Times:     times,
		Tools:     nil,
		UpdatedAt: time.Time{},
		URL:       r.FormValue("source"),
		Yield:     int16(yield),
	}

	recipeID, err := s.Repository.AddRecipe(recipe, r.Context().Value("userID").(int64))
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not add recipe.", errorToast))
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.Header().Set("HX-Redirect", "/recipes/"+strconv.FormatInt(recipeID, 10))
	w.WriteHeader(http.StatusCreated)
}

func recipeAddManualIngredientHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not parse the form.", errorToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	i := 1
	num := strconv.Itoa(i)
	for {
		if !r.Form.Has("ingredient-" + num) {
			break
		}

		i++
		num = strconv.Itoa(i)
	}

	if r.Form.Get(fmt.Sprintf("ingredient-%d", i-1)) == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var sb strings.Builder
	xs := []string{
		`<li class="pb-2 pl-2">`,
		`<label><input autofocus type="text" name="ingredient-` + num + `" placeholder="Ingredient #` + num + `" required class="w-8/12 py-1 pl-1 text-gray-600 placeholder-gray-400 bg-white border border-gray-400 dark:bg-gray-900 dark:border-none dark:text-gray-200" _="on keydown if event.key is 'Enter' then halt the event then get next <button/> from the parentElement of me then call htmx.trigger(it, 'click')"></label>`,
		`&nbsp;<button type="button" class="w-10 h-10 text-center bg-green-300 border border-gray-800 rounded-lg md:w-7 md:h-7 hover:bg-green-600 hover:text-white center dark:bg-green-500" title="Shortcut: Enter" hx-post="/recipes/add/manual/ingredient" hx-target="#ingredients-list" hx-swap="beforeend" hx-include="[name^='ingredient']">+</button>`,
		`&nbsp;<button type="button" class="delete-button w-10 h-10 bg-red-300 border border-gray-800 rounded-lg md:w-7 md:h-7 hover:bg-red-600 hover:text-white center dark:bg-red-500" hx-target="#ingredients-list" hx-post="/recipes/add/manual/ingredient/` + num + `" hx-include="[name^='ingredient']">-</button>`,
		`&nbsp;<div class="inline-block h-4 cursor-move handle"><svg xmlns="http://www.w3.org/2000/svg" class="md:w-4 md:h-4 w-7 h-7" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"/></svg></div>`,
	}
	for _, x := range xs {
		sb.WriteString(x)
	}
	fmt.Fprintf(w, sb.String())
}

func recipeAddManualIngredientDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not parse the form.", errorToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	count := 0
	i := 1
	for {
		if !r.Form.Has("ingredient-" + strconv.Itoa(i)) {
			break
		}

		i++
		count++
	}

	if count == 1 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var sb strings.Builder
	entry := chi.URLParam(r, "entry")

	curr := 1
	i = 1
	for {
		key := "ingredient-" + strconv.Itoa(i)
		if !r.Form.Has(key) {
			break
		}

		n, _ := strconv.Atoi(entry)
		if n == i {
			i++
			continue
		}

		currStr := strconv.Itoa(curr)
		xs := []string{
			`<li class="pb-2 pl-2">`,
			`<label><input type="text" name="ingredient-` + currStr + `" placeholder="Ingredient #` + currStr + `" required class="w-8/12 py-1 pl-1 text-gray-600 placeholder-gray-400 bg-white border border-gray-400 dark:bg-gray-900 dark:border-none dark:text-gray-200" ` + `value="` + r.Form.Get(key) + `" _="on keydown if event.key is 'Enter' then halt the event then get next <button/> from the parentElement of me then call htmx.trigger(it, 'click')"></label>`,
			`&nbsp;<button type="button" class="w-10 h-10 text-center bg-green-300 border border-gray-800 rounded-lg md:w-7 md:h-7 hover:bg-green-600 hover:text-white center dark:bg-green-500" title="Shortcut: Enter" hx-post="/recipes/add/manual/ingredient" hx-target="#ingredients-list" hx-swap="beforeend" hx-include="[name^='ingredient']">+</button>`,
			`&nbsp;<button type="button" class="delete-button w-10 h-10 bg-red-300 border border-gray-800 rounded-lg md:w-7 md:h-7 hover:bg-red-600 hover:text-white center dark:bg-red-500" hx-target="#ingredients-list" hx-post="/recipes/add/manual/ingredient/` + currStr + `" hx-include="[name^='ingredient']">-</button>`,
			`&nbsp;<div class="inline-block h-4 cursor-move handle"><svg xmlns="http://www.w3.org/2000/svg" class="md:w-4 md:h-4 w-7 h-7" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"/></svg></div>`,
		}
		for _, x := range xs {
			sb.WriteString(x)
		}

		i++
		curr++
	}

	fmt.Fprintf(w, sb.String())
}

func recipeAddManualInstructionHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not parse the form.", errorToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	i := 1
	num := fmt.Sprintf("%d", i)
	for {
		if !r.Form.Has("instruction-" + num) {
			break
		}

		i++
		num = fmt.Sprintf("%d", i)
	}

	if r.Form.Get(fmt.Sprintf("instruction-%d", i-1)) == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var sb strings.Builder
	xs := []string{
		`<li class="pb-2 pl-2 md:pl-0">`,
		`<label><textarea autofocus required name="instruction-` + num + `" rows="3" class="w-9/12 border border-gray-300 pl-1 md:w-5/6 xl:w-11/12 dark:bg-gray-900 dark:border-none" placeholder="Instruction #` + num + `" _="on keydown if event.ctrlKey and event.key === 'Enter' then halt the event then get next <div/> from the parentElement of me then get last <button/> in it then call htmx.trigger(it, 'click')"></textarea>&nbsp;</label>`,
		`<div class="inline-flex flex-col-reverse">`,
		`<button type="button" class="delete-button mt-4 md:flex-initial w-10 h-10 right-0.5 md:w-7 md:h-7 md:right-auto bg-red-300 border border-gray-800 rounded-lg top-3 hover:bg-red-600 hover:text-white center dark:bg-red-500" hx-target="#instructions-list" hx-post="/recipes/add/manual/instruction/` + num + `" hx-include="[name^='instruction']">-</button>`,
		`<button type="button" class="bottom-0 right-0.5 md:flex-initial md:w-7 md:h-7 md:right-auto w-10 h-10 text-center bg-green-300 border border-gray-800 rounded-lg hover:bg-green-600 hover:text-white center dark:bg-green-500" title="Shortcut: CTRL + Enter" hx-post="/recipes/add/manual/instruction" hx-target="#instructions-list" hx-swap="beforeend" hx-include="[name^='instruction']">+</button>`,
		`</div>&nbsp;<div class="inline-block h-4 cursor-move handle"><svg xmlns="http://www.w3.org/2000/svg" class="md:w-4 md:h-4 w-7 h-7" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"/></svg></div>`,
	}
	for _, x := range xs {
		sb.WriteString(x)
	}
	fmt.Fprintf(w, sb.String())
}

func recipeAddManualInstructionDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Could not parse the form.", errorToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	count := 0
	i := 1
	for {
		if !r.Form.Has("instruction-" + strconv.Itoa(i)) {
			break
		}

		i++
		count++
	}

	if count == 1 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var sb strings.Builder
	entry := chi.URLParam(r, "entry")

	curr := 1
	i = 1
	for {
		key := "instruction-" + strconv.Itoa(i)
		if !r.Form.Has(key) {
			break
		}

		n, _ := strconv.Atoi(entry)
		if n == i {
			i++
			continue
		}

		currStr := strconv.Itoa(curr)
		value := r.Form.Get(key)

		xs := []string{
			`<li class="pb-2 pl-2 md:pl-0">`,
			`<label><textarea required name="instruction-` + currStr + `" rows="3" class="w-9/12 border border-gray-300 pl-1 md:w-5/6 xl:w-11/12 dark:bg-gray-900 dark:border-none" placeholder="Instruction #` + currStr + `" _="on keydown if event.ctrlKey and event.key === 'Enter' then halt the event then get next <div/> from the parentElement of me then get last <button/> in it then call htmx.trigger(it, 'click')">` + value + `</textarea>&nbsp;</label>`,
			`<div class="inline-flex flex-col-reverse">`,
			`<button type="button" class="delete-button mt-4 md:flex-initial w-10 h-10 right-0.5 md:w-7 md:h-7 md:right-auto bg-red-300 border border-gray-800 rounded-lg top-3 hover:bg-red-600 hover:text-white center dark:bg-red-500" hx-target="#instructions-list" hx-post="/recipes/add/manual/instruction/` + currStr + `" hx-include="[name^='instruction']">-</button>`,
			`<button type="button" class="bottom-0 right-0.5 md:flex-initial md:w-7 md:h-7 md:right-auto w-10 h-10 text-center bg-green-300 border border-gray-800 rounded-lg hover:bg-green-600 hover:text-white center dark:bg-green-500" title="Shortcut: CTRL + Enter" hx-post="/recipes/add/manual/instruction" hx-target="#instructions-list" hx-swap="beforeend" hx-include="[name^='instruction']">+</button>`,
			`</div>&nbsp;<div class="inline-block h-4 cursor-move handle"><svg xmlns="http://www.w3.org/2000/svg" class="md:w-4 md:h-4 w-7 h-7" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"/></svg></div>`,
		}
		for _, x := range xs {
			sb.WriteString(x)
		}

		i++
		curr++
	}

	fmt.Fprintf(w, sb.String())
}

func (s *Server) recipesAddRequestWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	s.Email.Send(app.Config.Email.From, templates.EmailRequestWebsite, templates.EmailData{
		Text: r.FormValue("website"),
	})

	w.Header().Set("HX-Redirect", "/recipes/add")
	w.Header().Set("HX-Trigger", makeToast("I love chicken", infoToast))
	http.Redirect(w, r, "/recipes/add", http.StatusSeeOther)
}

func (s *Server) recipesAddWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	rawURL := r.Header.Get("HX-Prompt")
	if rawURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(rawURL); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Invalid URI.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rs, err := scraper.Scrape(rawURL, s.Files)
	if err != nil {
		templates.RenderComponent(w, "recipes", "unsupported-website", templates.Data{
			IsAuthenticated: true,
			Scraper: templates.ScraperData{
				UnsupportedWebsite: rawURL,
			},
		})
		return
	}

	recipe, err := rs.Recipe()
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Recipe schema is invalid.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := s.Repository.AddRecipe(recipe, r.Context().Value("userID").(int64)); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Recipe could not be added.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func (s *Server) recipeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("userID").(int64)

	rowsAffected, err := s.Repository.DeleteRecipe(id, userID)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Recipe could not be deleted.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		w.Header().Set("HX-Trigger", makeToast("Recipe not found.", errorToast))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) recipesExportHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	recipes := s.Repository.Recipes(userID)
	if len(recipes) == 0 {
		w.Header().Set("HX-Trigger", makeToast("No recipes in database.", warningToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	fileName, err := s.Files.ExportRecipes(recipes)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Failed to export recipes.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/download/"+fileName)
	w.WriteHeader(http.StatusSeeOther)
}

func (s *Server) recipeShareHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var userID int64
	isLoggedIn := true
	if userID = getUserIDFromSessionCookie(r); userID == -1 {
		if userID = getUserIDFromRememberMeCookie(r, s.Repository.GetAuthToken); userID == -1 {
			isLoggedIn = false
		}
	}

	if !s.Repository.IsRecipeShared(id) {
		notFoundHandler(w, r)
		return
	}

	userRecipeID := s.Repository.RecipeUser(id)
	recipe, err := s.Repository.Recipe(id, userRecipeID)
	if err != nil {
		notFoundHandler(w, r)
		return
	}

	data := templates.Data{
		IsAuthenticated: isLoggedIn,
		Title:           recipe.Name,
		View:            templates.NewViewRecipeData(id, recipe, userID == userRecipeID, true),
	}

	if r.Header.Get("Hx-Request") == "true" {
		templates.RenderComponent(w, "recipes", "view-recipe", data)
	} else {
		templates.Render(w, templates.ViewRecipePage, data)
	}
}

func (s *Server) recipeSharePostHandler(w http.ResponseWriter, r *http.Request) {
	recipeID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	link := app.Config.Address() + r.URL.String()
	if err := s.Repository.AddShareLink(link, recipeID); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Failed to create share link.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templates.RenderComponent(w, "recipes", "share-recipe", templates.Data{Content: link})
}

func (s *Server) recipesSupportedWebsitesHandler(w http.ResponseWriter, _ *http.Request) {
	websites := s.Repository.Websites()
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprintf(w, websites.TableHTML())
}

func (s *Server) recipesViewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID := r.Context().Value("userID").(int64)
	recipe, err := s.Repository.Recipe(id, userID)
	if err != nil {
		notFoundHandler(w, r)
		return
	}

	data := templates.Data{
		IsAuthenticated: true,
		Title:           recipe.Name,
		View:            templates.NewViewRecipeData(id, recipe, true, false),
	}

	if r.Header.Get("Hx-Request") == "true" {
		templates.RenderComponent(w, "recipes", "view-recipe", data)
	} else {
		templates.Render(w, templates.ViewRecipePage, data)
	}
}
