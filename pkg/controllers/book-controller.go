package controllers

import (
	"enconding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/sancheschris/go-bookstore/pkg/utils"
	"github.com/sancheschris/go-bookstore/pkg/models"
)

var NewBook