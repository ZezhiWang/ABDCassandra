package main 

type TagVal struct {
	Id 		string
	Ver 	int
	Val		string
}

func (t *TagVal) smaller(x TagVal) bool {
	var res bool
	if t.Ver < x.Ver {
		res = true
	} else if t.Ver > x.Ver {
		res = false
	} else {
		res = t.Id < x.Id
	}
	return res
}

func (t *Tagval) update(id string, val string) {
	tv.Id = id
	tv.Ver += 1
	tv.Val = val
}
