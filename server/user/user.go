
package user

import "sync"

type User struct {
    Username string
    Password string
}

var users = struct {
    sync.RWMutex
    m map[string]User
}{m: make(map[string]User)}

func RegisterUser(username, password string) bool {
    users.Lock()
    defer users.Unlock()

    if _, exists := users.m[username]; exists {
        return false
    }

    users.m[username] = User{Username: username, Password: password}
    return true
}

func AuthenticateUser(username, password string) bool {
    users.RLock()
    defer users.RUnlock()

    user, exists := users.m[username]
    if !exists || user.Password != password {
        return false
    }

    return true
}

