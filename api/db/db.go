package db

import (
    "github.com/prologic/bitcask"
    )

func GetSession() *bitcask.Bitcask {
    db, _ := bitcask.Open("/tmp/db")
    return db
}