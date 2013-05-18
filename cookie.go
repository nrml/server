package server

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, name string, value string, expire time.Time) {

	cookie := http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
		//Domain:"localhost:8080",
		Expires:    expire,
		RawExpires: expire.Format(time.UnixDate),
		MaxAge:     86400,
		//Secure: true,
		//HttpOnly: true,
		Raw:      "user=" + value,
		Unparsed: []string{"user=" + value},
	}

	http.SetCookie(w, &cookie)
}

func ClearCookie(w http.ResponseWriter, name string) {

	cookie := http.Cookie{
		Name:  name,
		Value: "rubbish",
		Path:  "/",
		//Domain:"localhost:8080",
		//Expires: expire,
		//RawExpires: expire.Format(time.UnixDate),
		MaxAge: -1,
		//Secure: true,
		//HttpOnly: true,
		Raw:      "user=" + "rubbish",
		Unparsed: []string{"user=" + "rubbish"},
	}

	http.SetCookie(w, &cookie)
}
