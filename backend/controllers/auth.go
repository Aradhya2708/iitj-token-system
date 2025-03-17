// ref https://github.com/jtblin/go-ldap-client

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	ldap "github.com/jtblin/go-ldap-client"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var client = &ldap.LDAPClient{
	Base:               "dc=example,dc=com",
	Host:               "ldap.forumsys.com", // test server ref https://www.forumsys.com/2022/05/10/online-ldap-test-server/
	Port:               389,
	UseSSL:             false,
	InsecureSkipVerify: false,
	SkipTLS:            true, // fix for tls handshake err [maybe Host doesnt support newer TLS]
	BindDN:             "cn=read-only-admin,dc=example,dc=com",
	BindPassword:       "password",
	UserFilter:         "(uid=%s)",
	GroupFilter:        "(memberUid=%s)",
	Attributes:         []string{"givenName", "sn", "mail", "uid"},
}

// handles LDAP auth req
func AuthenticateLDAP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	defer client.Close()

	authenticated, user, err := client.Authenticate(creds.Username, creds.Password)
	if err != nil {
		log.Printf("LDAP Authentication Error: %v", err)
		http.Error(w, "Authentication failed due to server error", http.StatusInternalServerError)
		return
	}

	if !authenticated {
		log.Printf("Invalid credentials for user %s", creds.Username)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	groups, err := client.GetGroupsOfUser(creds.Username)
	if err != nil {
		log.Printf("Error fetching groups for user %s: %v", creds.Username, err)
		http.Error(w, "Authenticated but failed to retrieve groups", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Authenticated successfully",
		"user":    user,
		"groups":  groups,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
