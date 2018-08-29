package xinge

import "time"

// PlatformIOS 修改push平台为iOS
func OptionPlatIos() ReqOption {
	return OptionPlatform(PlatformiOS)
}

// PlatformAndroid 修改push平台为PlatformAndroid
func OptionPlatAndroid() ReqOption {
	return OptionPlatform(PlatformAndroid)
}

// Platform 修改platform
func OptionPlatform(p Platform) ReqOption {
	return func(r *PushRequest) error {
		r.Platform = p
		if p == PlatformiOS {
			r.Android = nil
			r.IOS = &IOSParams{}
		} else if p == PlatformAndroid {
			r.IOS = nil
			r.Android = &AndroidParams{}
		}
		return nil
	}
}

// EnvProd 修改请求环境为product，只对iOS有效
func OptionEnvProduction() ReqOption {
	return func(r *PushRequest) error {
		r.Environment = EnvProd
		return nil
	}
}

// EnvDev 修改请求环境为dev，只对iOS有效
func OptionEnvDevelop() ReqOption {
	return func(r *PushRequest) error {
		r.Environment = EnvDev
		return nil
	}
}

// Title 修改push title
func OptionTitle(t string) ReqOption {
	return func(r *PushRequest) error {
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
func OptionContent(c string) ReqOption {
	return func(r *PushRequest) error {
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
func OptionNID(id int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.NID = id
		return nil
	}
}

// BuilderID 修改builder_id
func OptionBuilderID(id int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.BuilderID = id
		return nil
	}
}

// Ring 修改ring
func OptionRing(ring int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.Ring = ring
		return nil
	}
}

// RingRaw 修改ring_raw
func OptionRingRaw(rr string) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.RingRaw = rr
		return nil
	}
}

// Vibrate 修改vibrate
func OptionVibrate(v int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.Vibrate = v
		return nil
	}
}

// Lights 修改lights
func OptionLights(l int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.Lights = l
		return nil
	}
}

// Clearable 修改clearable
func OptionClearable(c int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.Clearable = c
		return nil
	}
}

// IconType 修改icon_type
func OptionIconType(it int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.IconType = it
		return nil
	}
}

// IconRes 修改icon_res
func OptionIconRes(ir string) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.IconRes = ir
		return nil
	}
}

// StyleID 修改style_id
func OptionStyleID(s int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.StyleID = s
		return nil
	}
}

// SmallIcon 修改small_icon
func OptionSmallIcon(si int) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.SmallIcon = si
		return nil
	}
}

// Action 修改action
func OptionAction(a map[string]interface{}) ReqOption {
	return func(r *PushRequest) error {
		r.Message.Android.Action = a
		return nil
	}
}

// AddAction 添加action
func OptionAddAction(k string, v interface{}) ReqOption {
	return func(r *PushRequest) error {
		if r.Message.Android.Action == nil {
			r.Message.Android.Action = map[string]interface{}{k: v}
		} else {
			r.Message.Android.Action[k] = v
		}
		return nil
	}
}

// CustomContent 修改custom_content 和 custom
func OptionCustomContent(ct map[string]string) ReqOption {
	return func(r *PushRequest) error {
		if r.Platform == PlatformAndroid {
			r.Message.Android.CustomContent = ct
			r.Message.IOS = nil
		} else {
			r.Message.Android = nil
			r.Message.IOS.Custom = ct

		}
		return nil
	}
}

// CustomContentSet 设置custom_content和custom的某个字段
func OptionCustomContentSet(k, v string) ReqOption {
	return func(r *PushRequest) error {
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
func OptionAps(aps *Aps) ReqOption {
	return func(r *PushRequest) error {
		r.Message.IOS.Aps = aps
		return nil
	}
}

// RequestAudienceType 修改audience_type
func OptionAudienceType(at AudienceType) ReqOption {
	return func(r *PushRequest) error {
		r.AudienceType = at
		return nil
	}
}

// Message 修改message
func OptionMessage(m Message) ReqOption {
	return func(r *PushRequest) error {
		r.Message = m
		return nil
	}
}

//TagList 俭省写法

func OptionTagList(op TagOpration, tags ...string) ReqOption {
	return func(r *PushRequest) error {
		r.TagList = &TagList{Tags: tags, TagOpration: op}
		return nil
	}
}

// TagList 修改tag_list
func OptionTagListOpt2(tl TagList) ReqOption {
	return func(r *PushRequest) error {
		r.TagList = &tl
		return nil
	}
}

// TokenList 修改token_list
func OptionTokenList(tl []string) ReqOption {
	return func(r *PushRequest) error {
		r.TokenList = tl
		return nil
	}
}

// TokenListAdd 给token_list添加一个token
func OptionTokenListAdd(t string) ReqOption {
	return func(r *PushRequest) error {
		if r.TokenList != nil {
			r.TokenList = append(r.TokenList, t)
		} else {
			r.TokenList = []string{t}
		}
		return nil
	}
}

// AccountList 修改account_list
func OptionAccountList(al []string) ReqOption {
	return func(r *PushRequest) error {
		r.AccountList = al
		return nil
	}
}

// AccountListAdd 给account_list添加一个account
func OptionAccountListAdd(a string) ReqOption {
	return func(r *PushRequest) error {
		if r.AccountList != nil {
			r.AccountList = append(r.AccountList, a)
		} else {
			r.AccountList = []string{a}
		}
		return nil
	}
}

// ExpireTime 修改expire_time
func OptionExpireTime(et time.Time) ReqOption {
	return func(r *PushRequest) error {
		r.ExpireTime = int(et.Unix())
		return nil
	}
}

// SendTime 修改send_time
func OptionSendTime(st time.Time) ReqOption {
	return func(r *PushRequest) error {
		r.SendTime = st.Format("2006-01-02 15:04:05:07")
		return nil
	}
}

// MultiPkg 修改multi_pkg
func OptionMultiPkg(mp bool) ReqOption {
	return func(r *PushRequest) error {
		r.MultiPkg = mp
		return nil
	}
}

// LoopTimes 修改loop_times
func OptionLoopTimes(lt int) ReqOption {
	return func(r *PushRequest) error {
		r.LoopTimes = lt
		return nil
	}
}

// StatTag 修改stat_tag
func OptionStatTag(st string) ReqOption {
	return func(r *PushRequest) error {
		r.StatTag = st
		return nil
	}
}

// Seq 修改seq
func OptionSeq(s int64) ReqOption {
	return func(r *PushRequest) error {
		r.Seq = s
		return nil
	}
}

// RequestAudienceType 修改account_type
func OptionAccountType(at int) ReqOption {
	return func(r *PushRequest) error {
		r.AccountType = at
		return nil
	}
}

// PushID 修改push_id
func OptionPushID(pid string) ReqOption {
	return func(r *PushRequest) error {
		r.PushID = pid
		return nil
	}
}

// MessageType 修改消息类型
func OptionMessageType(t MessageType) ReqOption {
	return func(r *PushRequest) error {
		r.MessageType = t
		return nil
	}
}
