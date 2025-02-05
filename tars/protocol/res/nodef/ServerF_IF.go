//Package nodef comment
// This file war generated by tars2go 1.1
// Generated from NodeF.tars
package nodef

import (
	"context"
	"fmt"
	m "github.com/HyperLiar/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//ServerF struct
type ServerF struct {
	s m.Servant
}

//KeepAlive is the proxy function for the method defined in the tars file, with the context
func (_obj *ServerF) KeepAlive(ServerInfo *ServerInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = ServerInfo.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "keepAlive", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//KeepAliveWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *ServerF) KeepAliveWithContext(ctx context.Context, ServerInfo *ServerInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = ServerInfo.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "keepAlive", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//ReportVersion is the proxy function for the method defined in the tars file, with the context
func (_obj *ServerF) ReportVersion(App string, ServerName string, Version string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(App, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(ServerName, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(Version, 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "reportVersion", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//ReportVersionWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *ServerF) ReportVersionWithContext(ctx context.Context, App string, ServerName string, Version string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(App, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(ServerName, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(Version, 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "reportVersion", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *ServerF) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *ServerF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

type _impServerF interface {
	KeepAlive(ServerInfo *ServerInfo) (ret int32, err error)
	ReportVersion(App string, ServerName string, Version string) (ret int32, err error)
}
type _impServerFWithContext interface {
	KeepAlive(ctx context.Context, ServerInfo *ServerInfo) (ret int32, err error)
	ReportVersion(ctx context.Context, App string, ServerName string, Version string) (ret int32, err error)
}

func keepAlive(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var ServerInfo ServerInfo
	err = ServerInfo.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}
	if withContext == false {
		_imp := _val.(_impServerF)
		ret, err := _imp.KeepAlive(&ServerInfo)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impServerFWithContext)
		ret, err := _imp.KeepAlive(ctx, &ServerInfo)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func reportVersion(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var App string
	err = _is.Read_string(&App, 1, true)
	if err != nil {
		return err
	}
	var ServerName string
	err = _is.Read_string(&ServerName, 2, true)
	if err != nil {
		return err
	}
	var Version string
	err = _is.Read_string(&Version, 3, true)
	if err != nil {
		return err
	}
	if withContext == false {
		_imp := _val.(_impServerF)
		ret, err := _imp.ReportVersion(App, ServerName, Version)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impServerFWithContext)
		ret, err := _imp.ReportVersion(ctx, App, ServerName, Version)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *ServerF) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "keepAlive":
		err := keepAlive(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "reportVersion":
		err := reportVersion(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
