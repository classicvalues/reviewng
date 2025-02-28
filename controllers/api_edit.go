package controllers

// MIT License
//
// Copyright (c) 2021 Damian Zaremba
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"encoding/json"
	"net/http"
)

func (app *App) ApiEditListHandler(w http.ResponseWriter, r *http.Request) {
	// Not logged in, return an error
	user := app.getAuthenticatedUser(r)
	if user == nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	// Not an admin, return an error
	if !user.Admin {
		http.Error(w, "Forbidden", 403)
		return
	}
}

func (app *App) ApiEditCreateHandler(w http.ResponseWriter, r *http.Request) {
	// Not logged in, return an error
	user := app.getAuthenticatedUser(r)
	if user == nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	// Not an admin, return an error
	if !user.Admin {
		http.Error(w, "Forbidden", 403)
		return
	}
}

func (app *App) ApiEditGetHandler(w http.ResponseWriter, r *http.Request) {
	// Not logged in, return an error
	user := app.getAuthenticatedUser(r)
	if user == nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	// Not an admin, return an error
	if !user.Admin {
		http.Error(w, "Forbidden", 403)
		return
	}
}

func (app *App) ApiEditUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Not logged in, return an error
	user := app.getAuthenticatedUser(r)
	if user == nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	// Not an admin, return an error
	if !user.Admin {
		http.Error(w, "Forbidden", 403)
		return
	}
}

func (app *App) ApiEditNextHandler(w http.ResponseWriter, r *http.Request) {
	// Not logged in, return an error
	user := app.getAuthenticatedUser(r)
	if user == nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	// Get an edit
	edit, err := app.dbh.CalculateRandomPendingEditForUser(user)
	if err != nil {
		panic(err)
	}

	if edit == nil {
		http.Error(w, "Not Found", 404)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"edit_id": edit.Id,
	})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		panic(err)
	}
}
