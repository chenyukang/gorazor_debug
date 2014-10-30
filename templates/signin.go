package templates

import (
	"bytes"
	"github.com/troyk/gorazor_debug/layout"
)

func Signin() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n    <style>\n      #signinCard {\n        width:275px;\n        margin:0 auto 25px;\n        margin-top:25px;\n      }\n\n      .card {\n        background-color: #f7f7f7;\n        border-radius: 2px;\n        box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);\n        padding:30px 30px;\n      }\n    </style>\n\n    <div id=\"signinCard\">\n      <div class=\"card\">\n        <form role=\"form\" action=\"#signin\">\n          <fieldset class=\"form-group\">\n              <input type=\"text\" class=\"form-control\" name=\"username\" placeholder=\"Username\" style=\"margin-bottom:.5em\">\n              <input type=\"password\" class=\"form-control\" name=\"password\" placeholder=\"Password\" style=\"margin-bottom:.5em\">\n          </fieldset>\n          <div id=\"signin_err\" style=\"text-align:center;color:red;font-weight:bold;\">\n            invalid username or password\n          </div>\n          <button type=\"submit\" class=\"btn btn-primary form-control\">Sign in</button>\n        </form>\n\n      </div>\n    </div>\n\n    <script language=\"javascript\">\n      console.log('here');\n    </script>")

	return layout.App(_buffer.String())
}
