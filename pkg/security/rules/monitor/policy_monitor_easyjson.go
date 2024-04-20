// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package monitor

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(in *jlexer.Lexer, out *RulesetLoadedEvent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "policies":
			if in.IsNull() {
				in.Skip()
				out.Policies = nil
			} else {
				in.Delim('[')
				if out.Policies == nil {
					if !in.IsDelim(']') {
						out.Policies = make([]*PolicyState, 0, 8)
					} else {
						out.Policies = []*PolicyState{}
					}
				} else {
					out.Policies = (out.Policies)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PolicyState
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PolicyState)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Policies = append(out.Policies, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "service":
			out.Service = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(out *jwriter.Writer, in RulesetLoadedEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"policies\":"
		out.RawString(prefix[1:])
		if in.Policies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Policies {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		out.String(string(in.Service))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RulesetLoadedEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RulesetLoadedEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(in *jlexer.Lexer, out *RuleState) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "expression":
			out.Expression = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Tags = make(map[string]string)
				} else {
					out.Tags = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 string
					v4 = string(in.String())
					(out.Tags)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "actions":
			if in.IsNull() {
				in.Skip()
				out.Actions = nil
			} else {
				in.Delim('[')
				if out.Actions == nil {
					if !in.IsDelim(']') {
						out.Actions = make([]RuleAction, 0, 2)
					} else {
						out.Actions = []RuleAction{}
					}
				} else {
					out.Actions = (out.Actions)[:0]
				}
				for !in.IsDelim(']') {
					var v5 RuleAction
					(v5).UnmarshalEasyJSON(in)
					out.Actions = append(out.Actions, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(out *jwriter.Writer, in RuleState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"expression\":"
		out.RawString(prefix)
		out.String(string(in.Expression))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if len(in.Tags) != 0 {
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v6First := true
			for v6Name, v6Value := range in.Tags {
				if v6First {
					v6First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v6Name))
				out.RawByte(':')
				out.String(string(v6Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Actions) != 0 {
		const prefix string = ",\"actions\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v7, v8 := range in.Actions {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(in *jlexer.Lexer, out *RuleSetAction) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "value":
			if m, ok := out.Value.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Value.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Value = in.Interface()
			}
		case "field":
			out.Field = string(in.String())
		case "append":
			out.Append = bool(in.Bool())
		case "scope":
			out.Scope = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(out *jwriter.Writer, in RuleSetAction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Value != nil {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Value.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Value.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Value))
		}
	}
	if in.Field != "" {
		const prefix string = ",\"field\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Field))
	}
	if in.Append {
		const prefix string = ",\"append\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Append))
	}
	if in.Scope != "" {
		const prefix string = ",\"scope\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Scope))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleSetAction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleSetAction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(in *jlexer.Lexer, out *RuleKillAction) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "signal":
			out.Signal = string(in.String())
		case "scope":
			out.Scope = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(out *jwriter.Writer, in RuleKillAction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Signal != "" {
		const prefix string = ",\"signal\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Signal))
	}
	if in.Scope != "" {
		const prefix string = ",\"scope\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Scope))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleKillAction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleKillAction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor4(in *jlexer.Lexer, out *RuleAction) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "filter":
			if in.IsNull() {
				in.Skip()
				out.Filter = nil
			} else {
				if out.Filter == nil {
					out.Filter = new(string)
				}
				*out.Filter = string(in.String())
			}
		case "set":
			if in.IsNull() {
				in.Skip()
				out.Set = nil
			} else {
				if out.Set == nil {
					out.Set = new(RuleSetAction)
				}
				(*out.Set).UnmarshalEasyJSON(in)
			}
		case "kill":
			if in.IsNull() {
				in.Skip()
				out.Kill = nil
			} else {
				if out.Kill == nil {
					out.Kill = new(RuleKillAction)
				}
				(*out.Kill).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor4(out *jwriter.Writer, in RuleAction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Filter != nil {
		const prefix string = ",\"filter\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Filter))
	}
	if in.Set != nil {
		const prefix string = ",\"set\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Set).MarshalEasyJSON(out)
	}
	if in.Kill != nil {
		const prefix string = ",\"kill\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Kill).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleAction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleAction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor4(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor5(in *jlexer.Lexer, out *PolicyState) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "source":
			out.Source = string(in.String())
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*RuleState, 0, 8)
					} else {
						out.Rules = []*RuleState{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *RuleState
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(RuleState)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.Rules = append(out.Rules, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor5(out *jwriter.Writer, in PolicyState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		out.String(string(in.Source))
	}
	{
		const prefix string = ",\"rules\":"
		out.RawString(prefix)
		if in.Rules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Rules {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PolicyState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PolicyState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor5(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor6(in *jlexer.Lexer, out *HeartbeatEvent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "policy":
			if in.IsNull() {
				in.Skip()
				out.Policy = nil
			} else {
				if out.Policy == nil {
					out.Policy = new(PolicyState)
				}
				(*out.Policy).UnmarshalEasyJSON(in)
			}
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "service":
			out.Service = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor6(out *jwriter.Writer, in HeartbeatEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"policy\":"
		out.RawString(prefix[1:])
		if in.Policy == nil {
			out.RawString("null")
		} else {
			(*in.Policy).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		out.String(string(in.Service))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HeartbeatEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor6(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HeartbeatEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor6(l, v)
}
