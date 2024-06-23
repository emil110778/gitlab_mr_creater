package response

import (
	"io"

	"github.com/xanzy/go-gitlab"
)

type Response struct {
	resp gitlab.Response
}

func GetResponse(resp *gitlab.Response) *Response {
	if resp == nil {
		return nil
	}
	return &Response{*resp}
}

func (r *Response) IsError() bool {
	return r.resp.StatusCode >= 400
}

func (r *Response) StatusCode() int {
	return r.resp.StatusCode
}
func (r *Response) Body() []byte {
	resp, err := io.ReadAll(r.resp.Body)
	if err != nil {
		return nil
	}
	return resp
}
