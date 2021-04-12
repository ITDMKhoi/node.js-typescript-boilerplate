
// The opensourcelive command starts a web server to render markdown recipe
// files as HTML web pages.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"html/template"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

// Page represents an HTML page.
type Page struct {
	Body template.HTML