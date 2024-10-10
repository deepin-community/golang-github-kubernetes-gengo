//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package maps

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	if in.Byte != nil {
		in, out := &in.Byte, &out.Byte
		*out = make(map[string]byte, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Int16 != nil {
		in, out := &in.Int16, &out.Int16
		*out = make(map[string]int16, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Int32 != nil {
		in, out := &in.Int32, &out.Int32
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Int64 != nil {
		in, out := &in.Int64, &out.Int64
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Uint8 != nil {
		in, out := &in.Uint8, &out.Uint8
		*out = make(map[string]byte, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Uint16 != nil {
		in, out := &in.Uint16, &out.Uint16
		*out = make(map[string]uint16, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Uint32 != nil {
		in, out := &in.Uint32, &out.Uint32
		*out = make(map[string]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Uint64 != nil {
		in, out := &in.Uint64, &out.Uint64
		*out = make(map[string]uint64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Float32 != nil {
		in, out := &in.Float32, &out.Float32
		*out = make(map[string]float32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Float64 != nil {
		in, out := &in.Float64, &out.Float64
		*out = make(map[string]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StringPtr != nil {
		in, out := &in.StringPtr, &out.StringPtr
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.StringPtrPtr != nil {
		in, out := &in.StringPtrPtr, &out.StringPtrPtr
		*out = make(map[string]**string, len(*in))
		for key, val := range *in {
			var outVal **string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(*string)
				if **in != nil {
					in, out := *in, *out
					*out = new(string)
					**out = **in
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.MapPtr != nil {
		in, out := &in.MapPtr, &out.MapPtr
		*out = make(map[string]*map[string]string, len(*in))
		for key, val := range *in {
			var outVal *map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(map[string]string)
				if **in != nil {
					in, out := *in, *out
					*out = make(map[string]string, len(*in))
					for key, val := range *in {
						(*out)[key] = val
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.SlicePtr != nil {
		in, out := &in.SlicePtr, &out.SlicePtr
		*out = make(map[string]*[]string, len(*in))
		for key, val := range *in {
			var outVal *[]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new([]string)
				if **in != nil {
					in, out := *in, *out
					*out = make([]string, len(*in))
					copy(*out, *in)
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Struct != nil {
		in, out := &in.Struct, &out.Struct
		*out = make(map[string]Ttest, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.StructPtr != nil {
		in, out := &in.StructPtr, &out.StructPtr
		*out = make(map[string]*Ttest, len(*in))
		for key, val := range *in {
			var outVal *Ttest
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Ttest)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ttest.
func (in *Ttest) DeepCopy() *Ttest {
	if in == nil {
		return nil
	}
	out := new(Ttest)
	in.DeepCopyInto(out)
	return out
}
