package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net.http"
	"strconv"
	"github.com/DurianDan/go-bookstore/pkg/utils"
	"github.com/DurianDan/go-bookstore/pkg/models"
)

var NewBook models.Book