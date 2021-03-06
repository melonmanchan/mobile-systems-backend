package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/melonmanchan/mobile-systems-backend/app"
	"github.com/melonmanchan/mobile-systems-backend/models"
	"github.com/melonmanchan/mobile-systems-backend/types"
	"github.com/melonmanchan/mobile-systems-backend/utils"

	"github.com/gorilla/mux"
	"github.com/guregu/null"
)

// UserHandler ...
func UserHandler(app app.App, r *mux.Router) {
	client := app.Client
	uploader := app.Uploader

	r.HandleFunc("/remove_device", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.DeviceRegisterRequest
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		valid, errs := req.IsValid()

		if !valid {
			utils.FailResponse(w, errs, http.StatusBadRequest)
			return
		}

		err = client.RemoveTokenFromUser(user, req.Token)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("PUT")

	r.HandleFunc("/register_device", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.DeviceRegisterRequest
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		valid, errs := req.IsValid()

		if !valid {
			utils.FailResponse(w, errs, http.StatusBadRequest)
			return
		}

		deviceAlreadyRegistered := false

		for _, token := range user.DeviceTokens {
			if token == req.Token {
				deviceAlreadyRegistered = true
				break
			}
		}

		if deviceAlreadyRegistered == false {
			err = client.AddTokenToUser(user, req.Token)
		}

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, _ := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")

	r.HandleFunc("/change_avatar", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)

		r.ParseMultipartForm(32 << 20)

		file, _, err := r.FormFile("avatar")

		defer file.Close()

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorUpdateProfileFailed}, http.StatusBadRequest)
			return
		}

		url, err := uploader.UploadAvatar(file)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorUpdateProfileFailed}, http.StatusBadRequest)
			return
		}

		user.Avatar = null.StringFrom(url)

		err = client.ChangeUserAvatar(user)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorUpdateProfileFailed}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: user, Status: 200}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")

	r.HandleFunc("/register_tutor_extra", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.RegisterTutorExtraRequest
		defer r.Body.Close()

		if user.UserType != models.TutorType {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericUserNotTutor}, http.StatusForbidden)
			return
		}

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		valid, errs := req.IsValid()

		if !valid {
			utils.FailResponse(w, errs, http.StatusBadRequest)
			return
		}

		user.Description = req.Description
		user.Price = null.IntFrom(req.Price)

		err = client.UpdateTutorProfile(user, req.Subjects)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorRegisterTutorFailed}, http.StatusInternalServerError)
			return
		}

		APIResp := types.APIResponse{Status: 200}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("POST")

	r.HandleFunc("/update_profile", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(types.UserKey).(*models.User)
		decoder := json.NewDecoder(r.Body)

		var req types.UpdateUserRequest
		defer r.Body.Close()

		err := decoder.Decode(&req)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorGenericRead}, http.StatusBadRequest)
			return
		}

		newUser := req.User
		newUser.ID = user.ID

		err = client.UpdateUserProfile(&newUser)

		if err != nil {
			utils.FailResponse(w, []types.APIError{types.ErrorUpdateProfileFailed}, http.StatusBadRequest)
			return
		}

		APIResp := types.APIResponse{Result: newUser, Status: 200}
		encoded, err := json.Marshal(APIResp)
		w.Write(encoded)
	}).Methods("PUT")
}
