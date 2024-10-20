package student

import (
	"crudoperation/internal/types"
	"crudoperation/internal/utils/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {

			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadGateway, response.GeneralError(err))
			return
		}

		// request validator

		if err := validator.New().Struct(student); err != nil {
			validatErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadGateway, response.ValidationError(validatErrs))
			return
		}

		slog.Info("creating a students")

		// w.Write([]byte("Welcome to students api"))

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})

	}
}
