package middleware

import (
    "errors"
    "net/http"

    "github.com/z4rathustr4/goapi/api"
    "github.com/z4rathustr4/goapi/internal/tools"
    log "github.com/sirupsen/logrus"

)

var UnauthorizedError = errors.New("ERR: Unauthorized")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        var username string = r.URL.Query().Get("username")
        var token = r.Header.Get("Authorization")
        var err error

        if username == "" || token == "" {
            log.Error(UnauthorizedError)
            api.RequestErrorHandler(w, UnauthorizedError)
            return
        }
        
        var db *tools.DBInterface
        db, err = tools.NewDB()
        if err != nil {
            api.InternalErrorHandler(w)
            return
        }
        
        var loginDetails *tools.LoginDetails
        loginDetails = (*db).GetUserLoginDetails(username)
        if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
            log.Error(UnauthorizedError)
            api.RequestErrorHandler(w, UnauthorizedError)
            return
        }
        next.ServeHTTP(w, r)
    })
}
