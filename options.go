package xinge

import "time"

//OptionPlatIos
func OptionPlatIos() PushMsgOption {
	return OptionPlatform(PlatformiOS)
}

//OptionPlatAndroid
func OptionPlatAndroid() PushMsgOption {
	return OptionPlatform(PlatformAndroid)
}

//OptionPlatform
func OptionPlatform(p Platform) PushMsgOption {
	return func(r *PushMsg) error {
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

//OptionEnvProduction
func OptionEnvProduction() PushMsgOption {
	return func(r *PushMsg) error {
		r.Environment = EnvProd
		return nil
	}
}

//OptionEnvDevelop
func OptionEnvDevelop() PushMsgOption {
	return func(r *PushMsg) error {
		r.Environment = EnvDev
		return nil
	}
}

//OptionTitle
func OptionTitle(t string) PushMsgOption {
	return func(r *PushMsg) error {
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

//OptionContent
func OptionContent(c string) PushMsgOption {
	return func(r *PushMsg) error {
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

//OptionNID
func OptionNID(id int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.NID = id
		return nil
	}
}

//OptionBuilderID
func OptionBuilderID(id int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.BuilderID = id
		return nil
	}
}

//OptionRing
func OptionRing(ring int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.Ring = ring
		return nil
	}
}

//OptionRingRaw
func OptionRingRaw(rr string) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.RingRaw = rr
		return nil
	}
}

//OptionVibrate
func OptionVibrate(v int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.Vibrate = v
		return nil
	}
}

//OptionLights
func OptionLights(l int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.Lights = l
		return nil
	}
}

//OptionCleanable
func OptionCleanable(c int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.Cleanable = c
		return nil
	}
}

//OptionIconType
func OptionIconType(it int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.IconType = it
		return nil
	}
}

//OptionIconRes
func OptionIconRes(ir string) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.IconRes = ir
		return nil
	}
}

//OptionStyleID
func OptionStyleID(s int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.StyleID = s
		return nil
	}
}

//OptionSmallIcon
func OptionSmallIcon(si int) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android.SmallIcon = si
		return nil
	}
}

//OptionAddAction
func OptionAddAction(k string, v interface{}) PushMsgOption {
	return func(r *PushMsg) error {
		if r.Message.Android.Action == nil {
			r.Message.Android.Action = map[string]interface{}{k: v}
		} else {
			r.Message.Android.Action[k] = v
		}
		return nil
	}
}

//OptionCustomContent
func OptionCustomContent(ct map[string]string) PushMsgOption {
	return func(r *PushMsg) error {
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

//OptionCustomContentSet
func OptionCustomContentSet(k, v string) PushMsgOption {
	return func(r *PushMsg) error {
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

//OptionAps
func OptionAps(aps *Aps) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.IOS.Aps = aps
		return nil
	}
}

//OptionAudienceType
func OptionAudienceType(at AudienceType) PushMsgOption {
	return func(r *PushMsg) error {
		r.AudienceType = at
		return nil
	}
}

//OptionMessage
func OptionMessage(m Message) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message = m
		return nil
	}
}

//OptionTagList
func OptionTagList(op TagOperation, tags ...string) PushMsgOption {
	return func(r *PushMsg) error {
		r.TagList = &TagList{Tags: tags, TagOperation: op}
		return nil
	}
}

//OptionTagListOpt2
func OptionTagListOpt2(tl TagList) PushMsgOption {
	return func(r *PushMsg) error {
		r.TagList = &tl
		return nil
	}
}

//OptionTokenList
func OptionTokenList(tl []string) PushMsgOption {
	return func(r *PushMsg) error {
		r.TokenList = tl
		return nil
	}
}

//OptionTokenListAdd
func OptionTokenListAdd(t string) PushMsgOption {
	return func(r *PushMsg) error {
		if r.TokenList != nil {
			r.TokenList = append(r.TokenList, t)
		} else {
			r.TokenList = []string{t}
		}
		return nil
	}
}

//OptionAccountList
func OptionAccountList(al []string) PushMsgOption {
	return func(r *PushMsg) error {
		r.AccountList = al
		return nil
	}
}

//OptionAccountListAdd
func OptionAccountListAdd(a string) PushMsgOption {
	return func(r *PushMsg) error {
		if r.AccountList != nil {
			r.AccountList = append(r.AccountList, a)
		} else {
			r.AccountList = []string{a}
		}
		return nil
	}
}

//OptionExpireTime
func OptionExpireTime(et time.Time) PushMsgOption {
	return func(r *PushMsg) error {
		r.ExpireTime = int(et.Unix())
		return nil
	}
}

//OptionSendTime 修改发送时间
func OptionSendTime(st time.Time) PushMsgOption {
	return func(r *PushMsg) error {
		r.SendTime = st.Format("2006-01-02 15:04:05:07")
		return nil
	}
}

//OptionMultiPkg
func OptionMultiPkg(mp bool) PushMsgOption {
	return func(r *PushMsg) error {
		r.MultiPkg = mp
		return nil
	}
}

//OptionLoopTimes
func OptionLoopTimes(lt int) PushMsgOption {
	return func(r *PushMsg) error {
		r.LoopTimes = lt
		return nil
	}
}

//OptionStatTag
func OptionStatTag(st string) PushMsgOption {
	return func(r *PushMsg) error {
		r.StatTag = st
		return nil
	}
}

//OptionSeq
func OptionSeq(s int64) PushMsgOption {
	return func(r *PushMsg) error {
		r.Seq = s
		return nil
	}
}

//OptionAccountType
func OptionAccountType(at int) PushMsgOption {
	return func(r *PushMsg) error {
		r.AccountType = at
		return nil
	}
}

//OptionPushID
func OptionPushID(pid string) PushMsgOption {
	return func(r *PushMsg) error {
		r.PushID = pid
		return nil
	}
}

//OptionMessageType
func OptionMessageType(t MsgType) PushMsgOption {
	return func(r *PushMsg) error {
		r.MsgType = t
		return nil
	}
}
