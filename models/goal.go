package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Goal struct {
	ID          int64
	Name        string
	Description string
	DateTime    time.Time
}
