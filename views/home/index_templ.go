// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/sonastea/ticketopia/internal/models"
	"github.com/sonastea/ticketopia/views/layouts"
	"math"
	"time"
)

type Events map[string][]models.Event

func Index(events Events, page string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n      @keyframes spin {\n        from {\n          transform: rotate(0deg);\n        }\n        to {\n          transform: rotate(360deg);\n        }\n      }\n\n      .event-details {\n        display: flex;\n        align-items: center;\n      }\n\n      .event-location:empty::before {\n\t\t\t\tdisplay: none;\n\t\t\t}\n\n      .event-location::before {\n        display: inline-block;\n        width: 0.25em;\n        height: 0.25em;\n        background-color: currentcolor;\n        border-radius: 50%;\n        content: \"\";\n        margin: 0 0.3em 0 0.1em;\n        transform: translateY(-50%);\n      }\n\n      .htmx-request.htmx-indicator {\n        animation: spin 1s linear infinite;\n      }\n    </style> <div class=\"flex flex-col items-center justify-center overflow-auto p-4\"><ul id=\"event-list\" class=\"max-w-screen-lg w-full grid grid-cols-1 gap-4 divide-y-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for name, event := range events {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"flex items-center justify-between overflow-hidden group hover:cursor-pointer\" data-name=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 52, Col: 110}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-on:click=\"alert(this.getAttribute(&#39;data-name&#39;))\"><div class=\"flex-shrink-0 w-24 h-24\"><img src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(getImageWithSize(event[len(event)-1].Images, 480, 360))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 54, Col: 72}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full h-full object-cover rounded-lg\" alt=\"Event thumbnail\"></div><div class=\"flex-grow p-4\"><div class=\"flex items-center justify-between\"><div><h2 class=\"font-semibold text-gray-800 group-hover:underline underline-offset-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 59, Col: 96}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><div class=\"text-sm text-gray-800\"><span class=\"sale-start-date\">Start ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(convertToLocalDate(event[len(event)-1].Sales.Public.StartDateTime))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 62, Col: 85}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"sale-end-date\">- ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(convertToLocalDate(event[len(event)-1].Sales.Public.EndDateTime))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 65, Col: 79}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><div class=\"text-sm text-gray-600\"><span class=\"event-city\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if len(event[len(event)-1].Embedded.Venues) > 0 {
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(getEventCityName(event[len(event)-1].Embedded.Venues[0].City.Name, event[len(event)-1].Embedded.Venues[0].State.Name))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 71, Col: 131}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(event[len(event)-1].Place.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 73, Col: 44}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"event-location\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(getEventLocationName(event[len(event)-1]))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 77, Col: 54}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div><div class=\"text-xs text-blue-600\"><a href=\"{eventLink}\" target=\"_blank\" class=\"hover:underline hidden sm:flex whitespace-nowrap\">Find Tickets</a></div></div></div></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex justify-center items-center\" id=\"pagination-button-container\"><button id=\"pagination-button\" type=\"button\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/?page=%v", page))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 89, Col: 89}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-target=\"#pagination-button-container\" hx-swap=\"delete\" hx-indicator=\"#spinner\">Load more</button><div id=\"spinner\" class=\"opacity-0 htmx-indicator ml-2 rounded-full h-5 w-5 border-2 border-t-blue-700\"></div></div></ul></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func convertToLocalDate(utcTime string) string {
	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return ""
	}

	localTime := t.Local()
	return localTime.Format("Jan 02")
}

func getEventCityName(city, state string) string {
	if state != "" || state == "" {
		return city
	}

	return fmt.Sprintf("%s, %s", city, state)
}

func getEventLocationName(event models.Event) string {
	if len(event.Embedded.Venues) > 0 {
		if event.Embedded.Venues[0].Name == "" {
			return event.Embedded.Venues[0].Country.Name
		}
		return event.Embedded.Venues[0].Name
	}

	return event.Place.City.Name
}

func getImageWithSize(images []models.EventImage, minWidth, minHeight int) string {
	var image string
	minDiff := math.MaxInt32

	for _, img := range images {
		diffWidth := abs(img.Width - minWidth)
		diffHeight := abs(img.Height - minHeight)
		totalDiff := diffWidth + diffHeight

		if totalDiff < minDiff {
			minDiff = totalDiff
			image = img.Url
		}
	}

	return image
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
