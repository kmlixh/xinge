package xinge

import "time"

// PlatformIOS 修改push平台为iOS
func PlatformiOSOpt() ReqOption {
	return func(r *Request) error {
		r.Platform = PlatformiOS
		return nil
	}
}

// PlatformAndroid 修改push平台为PlatformAndroid
func PlatformAndroidOpt() ReqOption {
	return func(r *Request) error {
		r.Platform = PlatformAndroid
		return nil
	}
}

// EnvProd 修改请求环境为product，只对iOS有效
func EnvProductionOpt() ReqOption {
	return func(r *Request) error {
		r.Environment = EnvProd
		return nil
	}
}

// EnvDev 修改请求环境为dev，只对iOS有效
func EnvDevelopOpt() ReqOption {
	return func(r *Request) error {
		r.Environment = EnvDev
		return nil
	}
}

// Title 修改push title
func TitleOpt(t string) ReqOption {
	return func(r *Request) error {
		r.Message.Title = t
		if r.Message.IOS != nil {
			if r.Message.IOS.Aps != nil {
				r.Message.IOS.Aps.Alert["title"] = t
			} else {
				r.Message.IOS.Aps = &Aps{
					Alert: map[string]string{"title": t},
				}
			}
		} else {
			r.Message.IOS = &IOSParams{
				Aps: &Aps{
					Alert: map[string]string{"title": t},
				},
			}
		}
		return nil
	}
}

// Content 修改push content
func ContentOpt(c string) ReqOption {
	return func(r *Request) error {
		r.Message.Content = c
		if r.Message.IOS != nil {
			if r.Message.IOS.Aps != nil {
				r.Message.IOS.Aps.Alert["body"] = c
			} else {
				r.Message.IOS.Aps = &Aps{
					Alert: map[string]string{"body": c},
				}
			}
		} else {
			r.Message.IOS = &IOSParams{
				Aps: &Aps{
					Alert: map[string]string{"body": c},
				},
			}
		}
		return nil
	}
}

// TODO: accept_time modify

// NID 修改nid
func NIDOpt(id int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.NID = id
		return nil
	}
}

// BuilderID 修改builder_id
func BuilderIDOpt(id int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.BuilderID = id
		return nil
	}
}

// Ring 修改ring
func RingOpt(ring int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.Ring = ring
		return nil
	}
}

// RingRaw 修改ring_raw
func RingRawOpt(rr string) ReqOption {
	return func(r *Request) error {
		r.Message.Android.RingRaw = rr
		return nil
	}
}

// Vibrate 修改vibrate
func VibrateOpt(v int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.Vibrate = v
		return nil
	}
}

// Lights 修改lights
func LightsOpt(l int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.Lights = l
		return nil
	}
}

// Clearable 修改clearable
func ClearableOpt(c int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.Clearable = c
		return nil
	}
}

// IconType 修改icon_type
func IconTypeOpt(it int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.IconType = it
		return nil
	}
}

// IconRes 修改icon_res
func IconResOpt(ir string) ReqOption {
	return func(r *Request) error {
		r.Message.Android.IconRes = ir
		return nil
	}
}

// StyleID 修改style_id
func StyleIDOpt(s int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.StyleID = s
	}
}

// SmallIcon 修改small_icon
func SmallIconOpt(si int) ReqOption {
	return func(r *Request) error {
		r.Message.Android.SmallIcon = si
		return nil
	}
}

// Action 修改action
func ActionOpt(a map[string]interface{}) ReqOption {
	return func(r *Request) error {
		r.Message.Android.Action = a
		return nil
	}
}

// AddAction 添加action
func AddActionOpt(k string, v interface{}) ReqOption {
	return func(r *Request) error {
		if r.Message.Android.Action == nil {
			r.Message.Android.Action = map[string]interface{}{k: v}
		} else {
			r.Message.Android.Action[k] = v
		}
		return nil
	}
}

// CustomContent 修改custom_content 和 custom
func CustomContentOpt(ct map[string]string) ReqOption {
	return func(r *Request) error {
		r.Message.Android.CustomContent = ct
		r.Message.IOS.Custom = ct
		return nil
	}
}

// CustomContentSet 设置custom_content和custom的某个字段
func CustomContentSetOpt(k, v string) ReqOption {
	return func(r *Request) error {
		if r.Message.Android.CustomContent == nil {
			r.Message.Android.CustomContent = map[string]string{k: v}
		} else {
			r.Message.Android.CustomContent[k] = v
		}
		if r.Message.IOS.Custom == nil {
			r.Message.IOS.Custom = map[string]string{k: v}
		} else {
			r.Message.IOS.Custom[k] = v
		}
		return nil
	}
}

// Aps 修改aps
func ApsOpt(aps *Aps) ReqOption {
	return func(r *Request) error {
		r.Message.IOS.Aps = aps
		return nil
	}
}

// RequestAudienceType 修改audience_type
func AudienceTypeOpt(at AudienceType) ReqOption {
	return func(r *Request) error {
		r.AudienceType = at
		return nil
	}
}

// Platform 修改platform
func PlatformOpt(p Platform) ReqOption {
	return func(r *Request) error {
		r.Platform = p
		return nil
	}
}

// Message 修改message
func MessageOpt(m Message) ReqOption {
	return func(r *Request) error {
		r.Message = m
		return nil
	}
}

//TagList 俭省写法

func TagListOpt(op TagOpration, tags ...string) ReqOption {
	return func(r *Request) error {
		r.TagList = TagList{Tags: tags, TagOpration: op}
		return nil
	}
}

// TagList 修改tag_list
func TagListOpt2(tl TagList) ReqOption {
	return func(r *Request) error {
		r.TagList = tl
		return nil
	}
}

// TokenList 修改token_list
func TokenListOpt(tl []string) ReqOption {
	return func(r *Request) error {
		r.TokenList = tl
		return nil
	}
}

// TokenListAdd 给token_list添加一个token
func TokenListAddOpt(t string) ReqOption {
	return func(r *Request) error {
		if r.TokenList != nil {
			r.TokenList = append(r.TokenList, t)
		} else {
			r.TokenList = []string{t}
		}
		return nil
	}
}

// AccountList 修改account_list
func AccountListOpt(al []string) ReqOption {
	return func(r *Request) error {
		r.AccountList = al
		return nil
	}
}

// AccountListAdd 给account_list添加一个account
func AccountListAddOpt(a string) ReqOption {
	return func(r *Request) error {
		if r.AccountList != nil {
			r.AccountList = append(r.AccountList, a)
		} else {
			r.AccountList = []string{a}
		}
		return nil
	}
}

// ExpireTime 修改expire_time
func ExpireTimeOpt(et time.Time) ReqOption {
	return func(r *Request) error {
		r.ExpireTime = int(et.Unix())
		return nil
	}
}

// SendTime 修改send_time
func SendTimeOpt(st time.Time) ReqOption {
	return func(r *Request) error {
		r.SendTime = st.Format("2006-01-02 15:04:05:07")
		return nil
	}
}

// MultiPkg 修改multi_pkg
func MultiPkgOpt(mp bool) ReqOption {
	return func(r *Request) error {
		r.MultiPkg = mp
		return nil
	}
}

// LoopTimes 修改loop_times
func LoopTimesOpt(lt int) ReqOption {
	return func(r *Request) error {
		r.LoopTimes = lt
		return nil
	}
}

// StatTag 修改stat_tag
func StatTagOpt(st string) ReqOption {
	return func(r *Request) error {
		r.StatTag = st
		return nil
	}
}

// Seq 修改seq
func SeqOpt(s int64) ReqOption {
	return func(r *Request) error {
		r.Seq = s
		return nil
	}
}

// RequestAudienceType 修改account_type
func AccountTypeOpt(at int) ReqOption {
	return func(r *Request) error {
		r.AccountType = at
		return nil
	}
}

// PushID 修改push_id
func PushIDOpt(pid string) ReqOption {
	return func(r *Request) error {
		r.PushID = pid
		return nil
	}
}

// MessageType 修改消息类型
func MessageTypeOpt(t MessageType) ReqOption {
	return func(r *Request) error {
		r.MessageType = t
		return nil
	}
}
