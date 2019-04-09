package usermodel

import (
    "bytes"
    "encoding/gob"
    "crypto/md5"
    "encoding/hex"
    "avnos/api/db"
)

type User struct {
    Email, Username, Fullname, Address string
    Password string `json: "-"`
}

type UserLogin struct {
    Username, Password string
}

func RevertUser(data []byte) (User, error) {
    buffer := bytes.NewBuffer(data)
    var res User
    enc := gob.NewDecoder(buffer)
    err := enc.Decode(&res)
    if err != nil {
        return User{}, err
    }
    return res, nil
}

func ConvertUser(s User) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(s)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func (u *User) EncryptPassword(){
    u.Password = Encrypt(u.Password)
}

func (u *UserLogin) EncryptPassword(){
    u.Password = Encrypt(u.Password)
}

func Encrypt(s string) string{
    hasher := md5.New()
    hasher.Write([]byte(s))
    return hex.EncodeToString(hasher.Sum(nil))
}

func GetUser(id string) (User, error) {
    d := db.GetSession()
    defer d.Close()
    data, err := d.Get(id)
    if err!= nil {
        return User{}, err
    }
    res, errc := RevertUser(data)
    if errc != nil {
        return User{}, errc
    }
    return res, nil
}