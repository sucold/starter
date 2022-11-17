package response

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

func Access(r *ghttp.Request) {
	r.Middleware.Next()
	r.LeaveTime = gtime.TimestampMilli()
	if r.GetError() != nil {
		handleError(r, r.GetError())
	} else {
		if exception := recover(); exception != nil {
			r.Response.WriteStatus(http.StatusInternalServerError)
			if v, ok := exception.(error); ok {
				if code := gerror.Code(v); code != gcode.CodeNil {
					handleError(r, v)
				} else {
					handleError(r, gerror.WrapCodeSkip(gcode.CodeInternalError, 1, v, ""))
				}
			} else {
				handleError(r, gerror.NewCodeSkipf(gcode.CodeInternalError, 1, "%+v", exception))
			}
		} else {
			handleAccess(r)
		}
	}
}
func handleAccess(r *ghttp.Request) {
	r.Server.Logger().File("{Y-m-d}.access.log").
		Stdout(r.Server.Logger().GetConfig().StdoutPrint).
		Printf(
			r.Context(),
			`%d "%s %s%s %s" %.3f, %s, "%s", "%s"`,
			r.Response.Status, r.Method, r.Host, r.URL.String(), r.Proto,
			float64(r.LeaveTime-r.EnterTime)/1000,
			r.GetClientIp(), r.Referer(),
			"",
			//r.UserAgent(),
		)
}
func handleError(r *ghttp.Request, err error) {
	var (
		code          = gerror.Code(err)
		codeDetail    = code.Detail()
		proto         = r.Header.Get("X-Forwarded-Proto")
		codeDetailStr string
	)
	if r.TLS != nil || gstr.Equal(proto, "https") {
	}
	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content := fmt.Sprintf(
		`%d "%s %s%s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		r.Response.Status, r.Method, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime-r.EnterTime)/1000,
		r.GetClientIp(), r.Referer(),
		"",
		//r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr,
	)
	if stack := gerror.Stack(err); stack != "" {
		content += "\nStack:\n" + stack
	} else {
		content += ", " + err.Error()
	}
	r.Server.Logger().File("{Y-m-d}.error.log").
		Stdout(r.Server.Logger().GetConfig().StdoutPrint).
		Print(r.Context(), content)
}
