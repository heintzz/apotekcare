package middleware

import (
	"context"
	"heintzz/ecommerce/internal/constants"
	"heintzz/ecommerce/internal/helper"
	"heintzz/ecommerce/internal/utils"
	"log"
	"time"

	"net/http"
	"strings"
)

func Tracer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
    
		log.Printf("method=%v path=%v type=[INFO] message='incoming request'", r.Method, r.URL.Path)
    
		h.ServeHTTP(w, r)
    
		end := time.Since(now).Milliseconds()
    
		log.Printf("method=%v path=%v type=[INFO] message='finish request' response_time=%vms", r.Method, r.URL.Path, end)
	})
}

func CheckToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")

		if bearer == "" {
			resp := helper.APIResponse{
				HttpCode:  http.StatusUnauthorized,				
				Message:   "no token provided",
			}
			resp.WriteJsonResponse(w)
			return
		}

		tokenSlice := strings.Split(bearer, "Bearer ")

		if len(tokenSlice) < 2 {
			resp := helper.APIResponse{
				HttpCode:  http.StatusUnauthorized,
				Error: 		 "unauthorized",
				Message:   "invalid token",
			}
			resp.WriteJsonResponse(w)
			return
		}

		tokenString := tokenSlice[1]
		token, err := utils.VerifyToken(tokenString)
		
		if err != nil {			
			resp := helper.APIResponse{
				HttpCode:  http.StatusUnauthorized,
				Error: 		 "unauthorized",				
			}
			resp.WriteJsonResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), constants.AUTH_EMAIL, token.Email)		
		ctx = context.WithValue(ctx, constants.AUTH_ROLE, token.Role)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}


func VerifyRole(allowedRoles ...string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {			
			role, ok := r.Context().Value(constants.AUTH_ROLE).(string)
						
			if !ok {
				resp := helper.APIResponse{
					HttpCode: http.StatusForbidden,
					Success:  false,
					Message:  "status forbidden",
					Error:    "invalid role",
				}
      	resp.WriteJsonResponse(w)			
				return
			}

			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					h.ServeHTTP(w, r)
					return
				}
			}
	
			resp := helper.APIResponse{
				HttpCode: http.StatusForbidden,
				Success:  false,
				Message:  "status forbidden",
				Error:    "you don't have access to this resource",
			}
      resp.WriteJsonResponse(w)			
		})
	}
}
